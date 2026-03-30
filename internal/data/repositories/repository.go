package repositories

import (
	"context"
	"fmt"
	"os"
	"strings"

	"connectrpc.com/connect"
	"github.com/Nerzal/gocloak/v13"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
	"github.com/charmmtech/identity/internal/types"
	"github.com/charmmtech/identity/utils"
)

type keycloakRepository struct {
	client       *gocloak.GoCloak
	baseURL      string
	clientID     string
	clientSecret string
	realm        string
}

const (
	roleManagedByAttribute = "managed_by"
	roleManagedByValue     = "identity-service"
)

func (repository *keycloakRepository) GetUser(ctx context.Context, userID string) (*identityv1.User, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	kcUser, err := repository.client.GetUserByID(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, repository.internal(err)
	}

	return repository.hydrateUser(ctx, token, kcUser, true, true)
}

func (repository *keycloakRepository) ListUsers(ctx context.Context, req *identityv1.ListUsersRequest) ([]*identityv1.User, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	params := gocloak.GetUsersParams{
		BriefRepresentation: gocloak.BoolP(false),
		Email:               utils.StringPtrOrNil(req.GetEmail()),
		Username:            utils.StringPtrOrNil(req.GetUsername()),
		Search:              utils.StringPtrOrNil(req.GetSearch()),
	}

	if page := req.GetPage(); page != nil {
		first := int(page.GetPage() * page.GetLimit())
		pageSize := int(page.GetLimit())
		if pageSize > 0 {
			params.First = &first
			params.Max = &pageSize
		}
	}

	var users []*gocloak.User
	if req.GetRole() != "" {
		users, err = repository.client.GetUsersByRoleName(ctx, token, repository.realm, req.GetRole(), gocloak.GetUsersByRoleParams{})
	} else {
		users, err = repository.client.GetUsers(ctx, token, repository.realm, params)
	}
	if err != nil {
		return nil, repository.internal(err)
	}

	response := make([]*identityv1.User, 0, len(users))
	for _, user := range users {
		hydrated, err := repository.hydrateUser(ctx, token, user, req.GetIncludeRoles(), req.GetIncludeGroups())
		if err != nil {
			return nil, err
		}
		if req.GetGroupId() != "" && !contains(hydrated.Groups, req.GetGroupId()) {
			continue
		}
		response = append(response, hydrated)
	}

	return response, nil
}

func (repository *keycloakRepository) CreateUser(ctx context.Context, req *identityv1.CreateUserRequest) (*identityv1.User, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	user := utils.KCUserFromInput(req.GetUser(), "")
	userID, err := repository.client.CreateUser(ctx, token, repository.realm, user)
	if err != nil {
		return nil, repository.internal(err)
	}

	if password := req.GetUser().GetPassword(); password != "" {
		if err := repository.client.SetPassword(ctx, token, userID, repository.realm, password, req.GetUser().GetTemporaryPassword()); err != nil {
			return nil, repository.internal(err)
		}
	}

	return repository.GetUser(ctx, userID)
}

func (repository *keycloakRepository) UpdateUser(ctx context.Context, req *identityv1.UpdateUserRequest) (*identityv1.User, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	user := utils.KCUserFromInput(req.GetUser(), req.GetId())
	if err := repository.client.UpdateUser(ctx, token, repository.realm, user); err != nil {
		return nil, repository.internal(err)
	}

	if password := req.GetUser().GetPassword(); password != "" {
		if err := repository.client.SetPassword(ctx, token, req.GetId(), repository.realm, password, req.GetUser().GetTemporaryPassword()); err != nil {
			return nil, repository.internal(err)
		}
	}

	return repository.GetUser(ctx, req.GetId())
}

func (repository *keycloakRepository) DeleteUser(ctx context.Context, userID string) error {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return repository.internal(err)
	}
	return repository.client.DeleteUser(ctx, token, repository.realm, userID)
}

func (repository *keycloakRepository) SetUserEnabled(ctx context.Context, userID string, enabled bool) (*identityv1.User, error) {
	user, err := repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	req := &identityv1.UpdateUserRequest{
		Id: userID,
		User: &identityv1.UserRequest{
			Username:        user.GetUsername(),
			Email:           user.GetEmail(),
			EmailVerified:   user.GetEmailVerified(),
			FirstName:       user.GetFirstName(),
			LastName:        user.GetLastName(),
			Enabled:         enabled,
			MfaEnabled:      user.GetMfaEnabled(),
			RequiredActions: user.GetRequiredActions(),
			Attributes:      user.GetAttributes(),
		},
	}

	return repository.UpdateUser(ctx, req)
}

func (repository *keycloakRepository) SetUserPassword(ctx context.Context, req *identityv1.SetUserPasswordRequest) error {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return repository.internal(err)
	}
	return repository.client.SetPassword(ctx, token, req.GetId(), repository.realm, req.GetPassword(), req.GetTemporary())
}

func (repository *keycloakRepository) ListUserSessions(ctx context.Context, userID string) ([]*identityv1.Session, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	sessions, err := repository.client.GetUserSessions(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, repository.internal(err)
	}

	resp := make([]*identityv1.Session, 0, len(sessions))
	for _, session := range sessions {
		resp = append(resp, utils.SessionFromKeycloak(session))
	}
	return resp, nil
}

func (repository *keycloakRepository) RevokeUserSession(ctx context.Context, req *identityv1.RevokeUserSessionRequest) error {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return repository.internal(err)
	}
	return repository.client.LogoutUserSession(ctx, token, repository.realm, req.GetSessionId())
}

func (repository *keycloakRepository) GetUserMfa(ctx context.Context, userID string) (*identityv1.MfaStatus, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	user, err := repository.client.GetUserByID(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, repository.internal(err)
	}

	configured, err := repository.client.GetConfiguredUserStorageCredentialTypes(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, repository.internal(err)
	}

	credentials, err := repository.client.GetCredentials(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, repository.internal(err)
	}

	disableable := make([]string, 0, len(credentials))
	for _, credential := range credentials {
		if credential != nil && credential.Type != nil {
			disableable = append(disableable, *credential.Type)
		}
	}

	return &identityv1.MfaStatus{
		Enabled:                gocloak.PBool(user.Totp) || contains(configured, "otp"),
		ConfiguredCredentials:  configured,
		DisableableCredentials: dedupe(disableable),
		RequiredActions:        append([]string(nil), gocloak.PStringSlice(user.RequiredActions)...),
	}, nil
}

func (repository *keycloakRepository) UpdateUserMfa(ctx context.Context, req *identityv1.UpdateUserMfaRequest) (*identityv1.MfaStatus, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	user, err := repository.client.GetUserByID(ctx, token, repository.realm, req.GetUserId())
	if err != nil {
		return nil, repository.internal(err)
	}

	user.Totp = gocloak.BoolP(req.GetEnabled())
	actions := utils.NormalizeStrings(req.GetRequiredActions())
	if req.GetEnabled() && !contains(actions, "CONFIGURE_TOTP") {
		actions = append(actions, "CONFIGURE_TOTP")
	}
	if !req.GetEnabled() {
		actions = removeString(actions, "CONFIGURE_TOTP")
		if err := repository.client.DisableAllCredentialsByType(ctx, token, repository.realm, req.GetUserId(), []string{"otp"}); err != nil {
			return nil, repository.internal(err)
		}
	}
	user.RequiredActions = &actions

	if err := repository.client.UpdateUser(ctx, token, repository.realm, *user); err != nil {
		return nil, repository.internal(err)
	}

	return repository.GetUserMfa(ctx, req.GetUserId())
}

func (repository *keycloakRepository) AssignRolesToUser(ctx context.Context, req *identityv1.AssignRolesToUserRequest) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	roles, err := repository.resolveRoles(ctx, token, req.GetRoleNames())
	if err != nil {
		return nil, err
	}
	if err := repository.client.AddRealmRoleToUser(ctx, token, repository.realm, req.GetUserId(), roles); err != nil {
		return nil, repository.internal(err)
	}
	return repository.ListUserRoles(ctx, req.GetUserId())
}

func (repository *keycloakRepository) RemoveRolesFromUser(ctx context.Context, req *identityv1.RemoveRolesFromUserRequest) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	roles, err := repository.resolveRoles(ctx, token, req.GetRoleNames())
	if err != nil {
		return nil, err
	}
	if err := repository.client.DeleteRealmRoleFromUser(ctx, token, repository.realm, req.GetUserId(), roles); err != nil {
		return nil, repository.internal(err)
	}
	return repository.ListUserRoles(ctx, req.GetUserId())
}

func (repository *keycloakRepository) AddUserToGroups(ctx context.Context, req *identityv1.AddUserToGroupsRequest) ([]*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	for _, groupID := range req.GetGroupIds() {
		if err := repository.client.AddUserToGroup(ctx, token, repository.realm, req.GetUserId(), groupID); err != nil {
			return nil, repository.internal(err)
		}
	}
	return repository.ListUserGroups(ctx, req.GetUserId())
}

func (repository *keycloakRepository) RemoveUserFromGroups(ctx context.Context, req *identityv1.RemoveUserFromGroupsRequest) ([]*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	for _, groupID := range req.GetGroupIds() {
		if err := repository.client.DeleteUserFromGroup(ctx, token, repository.realm, req.GetUserId(), groupID); err != nil {
			return nil, repository.internal(err)
		}
	}
	return repository.ListUserGroups(ctx, req.GetUserId())
}

func (repository *keycloakRepository) GetRole(ctx context.Context, roleName string) (*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}
	role, err := repository.findRoleByName(ctx, token, roleName)
	if err != nil {
		return nil, repository.internal(err)
	}
	return utils.RoleFromKeycloak(role, repository.realm), nil
}

func (repository *keycloakRepository) ListRoles(ctx context.Context, req *identityv1.ListRolesRequest) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	params := gocloak.GetRoleParams{Search: utils.StringPtrOrNil(req.GetSearch())}
	if page := req.GetPage(); page != nil {
		first := int(page.GetPage() * page.GetLimit())
		pageSize := int(page.GetLimit())
		if pageSize > 0 {
			params.First = &first
			params.Max = &pageSize
		}
	}

	roles, err := repository.client.GetRealmRoles(ctx, token, repository.realm, params)
	if err != nil {
		return nil, repository.internal(err)
	}

	resp := make([]*identityv1.Role, 0, len(roles))
	for _, role := range roles {
		if req.GetManagedOnly() && !isManagedRole(role, repository.realm) {
			continue
		}
		resp = append(resp, utils.RoleFromKeycloak(role, repository.realm))
	}
	return resp, nil
}

func (repository *keycloakRepository) CreateRole(ctx context.Context, req *identityv1.CreateRoleRequest) (*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	role := gocloak.Role{
		Name:        utils.StringPtrOrNil(req.GetRole().GetName()),
		Description: utils.StringPtrOrNil(req.GetRole().GetDescription()),
		Attributes:  managedRoleAttributes(req.GetRole().GetAttributes()),
	}
	if _, err := repository.client.CreateRealmRole(ctx, token, repository.realm, role); err != nil {
		return nil, repository.internal(err)
	}
	return repository.GetRole(ctx, req.GetRole().GetName())
}

func (repository *keycloakRepository) UpdateRole(ctx context.Context, req *identityv1.UpdateRoleRequest) (*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	role := gocloak.Role{
		Name:        utils.StringPtrOrNil(req.GetRole().GetName()),
		Description: utils.StringPtrOrNil(req.GetRole().GetDescription()),
		Attributes:  managedRoleAttributes(req.GetRole().GetAttributes()),
	}
	if err := repository.client.UpdateRealmRole(ctx, token, repository.realm, req.GetRoleName(), role); err != nil {
		return nil, repository.internal(err)
	}

	name := req.GetRoleName()
	if req.GetRole().GetName() != "" {
		name = req.GetRole().GetName()
	}
	return repository.GetRole(ctx, name)
}

func (repository *keycloakRepository) DeleteRole(ctx context.Context, roleName string) error {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return repository.internal(err)
	}
	return repository.client.DeleteRealmRole(ctx, token, repository.realm, roleName)
}

func (repository *keycloakRepository) GetGroup(ctx context.Context, groupID string) (*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}
	group, err := repository.client.GetGroup(ctx, token, repository.realm, groupID)
	if err != nil {
		return nil, repository.internal(err)
	}
	roles, err := repository.client.GetRealmRolesByGroupID(ctx, token, repository.realm, groupID)
	if err != nil {
		return nil, repository.internal(err)
	}
	resp := utils.GroupFromKeycloak(group, repository.realm)
	resp.Roles = roleNames(roles)
	return resp, nil
}

func (repository *keycloakRepository) ListGroups(ctx context.Context, req *identityv1.ListGroupsRequest) ([]*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	params := gocloak.GetGroupsParams{
		BriefRepresentation: gocloak.BoolP(false),
		Search:              utils.StringPtrOrNil(req.GetSearch()),
	}
	if page := req.GetPage(); page != nil {
		first := int(page.GetPage() * page.GetLimit())
		pageLimit := int(page.GetLimit())
		if pageLimit > 0 {
			params.First = &first
			params.Max = &pageLimit
		}
	}

	groups, err := repository.client.GetGroups(ctx, token, repository.realm, params)
	if err != nil {
		return nil, repository.internal(err)
	}

	resp := make([]*identityv1.Group, 0, len(groups))
	for _, group := range groups {
		item := utils.GroupFromKeycloak(group, repository.realm)
		if req.GetIncludeRoles() {
			roles, err := repository.client.GetRealmRolesByGroupID(ctx, token, repository.realm, item.GetId())
			if err != nil {
				return nil, repository.internal(err)
			}
			item.Roles = roleNames(roles)
		}
		resp = append(resp, item)
	}
	return resp, nil
}

func (repository *keycloakRepository) CreateGroup(ctx context.Context, req *identityv1.CreateGroupRequest) (*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	group := gocloak.Group{
		Name:       utils.StringPtrOrNil(req.GetGroup().GetName()),
		Path:       utils.StringPtrOrNil(req.GetGroup().GetPath()),
		Attributes: utils.FromAttributeMap(req.GetGroup().GetAttributes()),
	}
	groupID, err := repository.client.CreateGroup(ctx, token, repository.realm, group)
	if err != nil {
		return nil, repository.internal(err)
	}
	return repository.GetGroup(ctx, groupID)
}

func (repository *keycloakRepository) UpdateGroup(ctx context.Context, req *identityv1.UpdateGroupRequest) (*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	group := gocloak.Group{
		ID:         gocloak.StringP(req.GetId()),
		Name:       utils.StringPtrOrNil(req.GetGroup().GetName()),
		Path:       utils.StringPtrOrNil(req.GetGroup().GetPath()),
		Attributes: utils.FromAttributeMap(req.GetGroup().GetAttributes()),
	}
	if err := repository.client.UpdateGroup(ctx, token, repository.realm, group); err != nil {
		return nil, repository.internal(err)
	}
	return repository.GetGroup(ctx, req.GetId())
}

func (repository *keycloakRepository) DeleteGroup(ctx context.Context, groupID string) error {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return repository.internal(err)
	}
	return repository.client.DeleteGroup(ctx, token, repository.realm, groupID)
}

func (repository *keycloakRepository) AssignRolesToGroup(ctx context.Context, req *identityv1.AssignRolesToGroupRequest) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	roles, err := repository.resolveRoles(ctx, token, req.GetRoleNames())
	if err != nil {
		return nil, err
	}
	if err := repository.client.AddRealmRoleToGroup(ctx, token, repository.realm, req.GetId(), roles); err != nil {
		return nil, repository.internal(err)
	}
	return repository.ListGroupRoles(ctx, req.GetId())
}

func (repository *keycloakRepository) RemoveRolesFromGroup(ctx context.Context, req *identityv1.RemoveRolesFromGroupRequest) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}

	roles, err := repository.resolveRoles(ctx, token, req.GetRoleNames())
	if err != nil {
		return nil, err
	}
	if err := repository.client.DeleteRealmRoleFromGroup(ctx, token, repository.realm, req.GetId(), roles); err != nil {
		return nil, repository.internal(err)
	}
	return repository.ListGroupRoles(ctx, req.GetId())
}

func (repository *keycloakRepository) ListUserRoles(ctx context.Context, userID string) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}
	roles, err := repository.effectiveUserRoles(ctx, token, userID)
	if err != nil {
		return nil, repository.internal(err)
	}
	return mapRoles(repository.realm, roles), nil
}

func (repository *keycloakRepository) ListGroupRoles(ctx context.Context, groupID string) ([]*identityv1.Role, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}
	roles, err := repository.client.GetRealmRolesByGroupID(ctx, token, repository.realm, groupID)
	if err != nil {
		return nil, repository.internal(err)
	}
	return mapRoles(repository.realm, roles), nil
}

func (repository *keycloakRepository) ListUserGroups(ctx context.Context, userID string) ([]*identityv1.Group, error) {
	token, err := repository.loginAdmin(ctx)
	if err != nil {
		return nil, repository.internal(err)
	}
	groups, err := repository.client.GetUserGroups(ctx, token, repository.realm, userID, gocloak.GetGroupsParams{BriefRepresentation: gocloak.BoolP(false)})
	if err != nil {
		return nil, repository.internal(err)
	}
	resp := make([]*identityv1.Group, 0, len(groups))
	for _, group := range groups {
		resp = append(resp, utils.GroupFromKeycloak(group, repository.realm))
	}
	return resp, nil
}

func (repository *keycloakRepository) hydrateUser(ctx context.Context, token string, user *gocloak.User, includeRoles bool, includeGroups bool) (*identityv1.User, error) {
	var roles []string
	var groups []string
	var kcGroups []*gocloak.Group
	var groupDetails []*identityv1.Group
	var err error

	if includeRoles {
		var kcRoles []*gocloak.Role
		kcRoles, err = repository.effectiveUserRoles(ctx, token, gocloak.PString(user.ID))
		if err != nil {
			return nil, repository.internal(err)
		}
		roles = roleNames(kcRoles)
	}

	if includeGroups || includeRoles {
		kcGroups, err = repository.client.GetUserGroups(ctx, token, repository.realm, gocloak.PString(user.ID), gocloak.GetGroupsParams{BriefRepresentation: gocloak.BoolP(false)})
		if err != nil {
			return nil, repository.internal(err)
		}
	}

	if includeGroups {
		groups = groupIDs(kcGroups)
		groupDetails = mapGroups(repository.realm, kcGroups)
	}

	return utils.UserFromKeycloak(user, repository.realm, roles, groups, groupDetails), nil
}

func (repository *keycloakRepository) effectiveUserRoles(ctx context.Context, token string, userID string) ([]*gocloak.Role, error) {
	rolesByName := make(map[string]*gocloak.Role)

	directAndComposite, err := repository.client.GetCompositeRealmRolesByUserID(ctx, token, repository.realm, userID)
	if err != nil {
		return nil, err
	}
	for _, role := range directAndComposite {
		if role != nil && role.Name != nil {
			rolesByName[*role.Name] = role
		}
	}

	groups, err := repository.client.GetUserGroups(ctx, token, repository.realm, userID, gocloak.GetGroupsParams{BriefRepresentation: gocloak.BoolP(false)})
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group == nil || group.ID == nil {
			continue
		}
		groupRoles, err := repository.client.GetCompositeRealmRolesByGroupID(ctx, token, repository.realm, *group.ID)
		if err != nil {
			return nil, err
		}
		for _, role := range groupRoles {
			if role != nil && role.Name != nil {
				rolesByName[*role.Name] = role
			}
		}
	}

	roles := make([]*gocloak.Role, 0, len(rolesByName))
	for _, role := range rolesByName {
		roles = append(roles, role)
	}
	return roles, nil
}

func (repository *keycloakRepository) resolveRoles(ctx context.Context, token string, roleNames []string) ([]gocloak.Role, error) {
	roles := make([]gocloak.Role, 0, len(roleNames))
	for _, roleName := range roleNames {
		role, err := repository.findRoleByName(ctx, token, roleName)
		if err != nil {
			return nil, repository.internal(err)
		}
		roles = append(roles, *role)
	}
	return roles, nil
}

func (repository *keycloakRepository) findRoleByName(ctx context.Context, token string, roleName string) (*gocloak.Role, error) {
	trimmedName := strings.TrimSpace(roleName)
	if trimmedName == "" {
		return nil, fmt.Errorf("role name is required")
	}

	roles, err := repository.client.GetRealmRoles(ctx, token, repository.realm, gocloak.GetRoleParams{
		Search: gocloak.StringP(trimmedName),
	})
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		if role != nil && role.Name != nil && *role.Name == trimmedName {
			return role, nil
		}
	}

	return nil, fmt.Errorf("role %q not found", trimmedName)
}

func (repository *keycloakRepository) loginAdmin(ctx context.Context) (string, error) {
	jwtToken, err := repository.client.LoginClient(ctx, repository.clientID, repository.clientSecret, repository.realm)
	if err != nil {
		return "", err
	}
	return jwtToken.AccessToken, nil
}

func (repository *keycloakRepository) internal(err error) error {
	return connect.NewError(connect.CodeInternal, fmt.Errorf("keycloak provider error: %w", err))
}

func mapRoles(realm string, roles []*gocloak.Role) []*identityv1.Role {
	resp := make([]*identityv1.Role, 0, len(roles))
	for _, role := range roles {
		resp = append(resp, utils.RoleFromKeycloak(role, realm))
	}
	return resp
}

func mapGroups(realm string, groups []*gocloak.Group) []*identityv1.Group {
	resp := make([]*identityv1.Group, 0, len(groups))
	for _, group := range groups {
		resp = append(resp, utils.GroupFromKeycloak(group, realm))
	}
	return resp
}

func roleNames(roles []*gocloak.Role) []string {
	resp := make([]string, 0, len(roles))
	for _, role := range roles {
		if role != nil && role.Name != nil {
			resp = append(resp, *role.Name)
		}
	}
	return resp
}

func groupIDs(groups []*gocloak.Group) []string {
	resp := make([]string, 0, len(groups))
	for _, group := range groups {
		if group != nil && group.ID != nil {
			resp = append(resp, *group.ID)
		}
	}
	return resp
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func removeString(values []string, target string) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		if value != target {
			out = append(out, value)
		}
	}
	return out
}

func dedupe(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	out := make([]string, 0, len(values))
	for _, value := range values {
		if _, ok := seen[value]; ok || value == "" {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}

func NewKeycloakIdentityRepository(keycloak *gocloak.GoCloak) types.IdentityRepository {
	return &keycloakRepository{
		client:       keycloak,
		baseURL:      os.Getenv("KEYCLOAK.URL"),
		clientID:     os.Getenv("KEYCLOAK.CLIENT_ID"),
		clientSecret: os.Getenv("KEYCLOAK.CLIENT_SECRET"),
		realm:        os.Getenv("KEYCLOAK.REALM"),
	}
}

func managedRoleAttributes(src map[string]*identityv1.AttributeValues) *map[string][]string {
	attrs := utils.FromAttributeMap(src)
	if attrs == nil {
		attrs = &map[string][]string{}
	}
	(*attrs)[roleManagedByAttribute] = []string{roleManagedByValue}
	return attrs
}

func isManagedRole(role *gocloak.Role, realm string) bool {
	if role == nil || role.Name == nil {
		return false
	}
	if isDefaultKeycloakRole(*role.Name, realm) {
		return false
	}
	if role.Attributes == nil {
		return false
	}
	values, ok := (*role.Attributes)[roleManagedByAttribute]
	if !ok {
		return false
	}
	for _, value := range values {
		if value == roleManagedByValue {
			return true
		}
	}
	return false
}

func isDefaultKeycloakRole(roleName string, realm string) bool {
	if roleName == "offline_access" || roleName == "uma_authorization" {
		return true
	}
	return strings.EqualFold(roleName, "default-roles-"+realm)
}
