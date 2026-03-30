package utils

import (
	"strings"
	"time"

	"github.com/Nerzal/gocloak/v13"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProviderRef(realm string) *identityv1.ProviderRef {
	return &identityv1.ProviderRef{
		Type:  identityv1.IdentityProviderType_IDENTITY_PROVIDER_TYPE_KEYCLOAK,
		Realm: realm,
	}
}

func ToAttributeMap(src *map[string][]string) map[string]*identityv1.AttributeValues {
	if src == nil {
		return nil
	}

	dst := make(map[string]*identityv1.AttributeValues, len(*src))
	for key, values := range *src {
		copied := append([]string(nil), values...)
		dst[key] = &identityv1.AttributeValues{Values: copied}
	}
	return dst
}

func FromAttributeMap(src map[string]*identityv1.AttributeValues) *map[string][]string {
	if len(src) == 0 {
		return nil
	}

	dst := make(map[string][]string, len(src))
	for key, values := range src {
		if values == nil {
			dst[key] = nil
			continue
		}
		dst[key] = append([]string(nil), values.Values...)
	}
	return &dst
}

func KCUserFromInput(input *identityv1.UserRequest, userID string) gocloak.User {
	user := gocloak.User{
		ID:            StringPtrOrNil(userID),
		Username:      StringPtrOrNil(input.GetUsername()),
		Email:         StringPtrOrNil(input.GetEmail()),
		EmailVerified: gocloak.BoolP(input.GetEmailVerified()),
		FirstName:     StringPtrOrNil(input.GetFirstName()),
		LastName:      StringPtrOrNil(input.GetLastName()),
		Enabled:       gocloak.BoolP(input.GetEnabled()),
		Attributes:    FromAttributeMap(input.GetAttributes()),
	}

	if actions := NormalizeStrings(input.GetRequiredActions()); len(actions) > 0 {
		user.RequiredActions = &actions
	}

	if input.GetMfaEnabled() {
		user.Totp = gocloak.BoolP(true)
	}

	return user
}

func UserFromKeycloak(user *gocloak.User, realm string, roles []string, groups []string, groupDetails []*identityv1.Group) *identityv1.User {
	if user == nil {
		return nil
	}

	resp := &identityv1.User{
		Id:              gocloak.PString(user.ID),
		Provider:        ProviderRef(realm),
		Username:        gocloak.PString(user.Username),
		Email:           gocloak.PString(user.Email),
		EmailVerified:   gocloak.PBool(user.EmailVerified),
		FirstName:       gocloak.PString(user.FirstName),
		LastName:        gocloak.PString(user.LastName),
		Enabled:         gocloak.PBool(user.Enabled),
		MfaEnabled:      gocloak.PBool(user.Totp),
		RequiredActions: append([]string(nil), gocloak.PStringSlice(user.RequiredActions)...),
		Roles:           append([]string(nil), roles...),
		Groups:          append([]string(nil), groups...),
		GroupDetails:    append([]*identityv1.Group(nil), groupDetails...),
		Attributes:      ToAttributeMap(user.Attributes),
	}

	if user.CreatedTimestamp != nil && *user.CreatedTimestamp > 0 {
		resp.CreatedAt = timestamppb.New(time.UnixMilli(*user.CreatedTimestamp))
	}

	return resp
}

func RoleFromKeycloak(role *gocloak.Role, realm string) *identityv1.Role {
	if role == nil {
		return nil
	}

	return &identityv1.Role{
		Id:          gocloak.PString(role.ID),
		Provider:    ProviderRef(realm),
		Name:        gocloak.PString(role.Name),
		Description: gocloak.PString(role.Description),
		Composite:   gocloak.PBool(role.Composite),
		ClientRole:  gocloak.PBool(role.ClientRole),
		Attributes:  ToAttributeMap(role.Attributes),
	}
}

func GroupFromKeycloak(group *gocloak.Group, realm string) *identityv1.Group {
	if group == nil {
		return nil
	}

	return &identityv1.Group{
		Id:         gocloak.PString(group.ID),
		Provider:   ProviderRef(realm),
		Name:       gocloak.PString(group.Name),
		Path:       gocloak.PString(group.Path),
		Roles:      append([]string(nil), gocloak.PStringSlice(group.RealmRoles)...),
		Attributes: ToAttributeMap(group.Attributes),
	}
}

func SessionFromKeycloak(session *gocloak.UserSessionRepresentation) *identityv1.Session {
	if session == nil {
		return nil
	}

	resp := &identityv1.Session{
		Id:        gocloak.PString(session.ID),
		UserId:    gocloak.PString(session.UserID),
		Username:  gocloak.PString(session.Username),
		IpAddress: gocloak.PString(session.IPAddress),
	}

	if session.Clients != nil {
		resp.Clients = *session.Clients
	}
	if session.Start != nil && *session.Start > 0 {
		resp.StartedAt = timestamppb.New(time.Unix(*session.Start, 0))
	}
	if session.LastAccess != nil && *session.LastAccess > 0 {
		resp.LastAccessedAt = timestamppb.New(time.UnixMilli(*session.LastAccess))
	}

	return resp
}

func StringPtrOrNil(value string) *string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	return gocloak.StringP(value)
}

func NormalizeStrings(values []string) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}
