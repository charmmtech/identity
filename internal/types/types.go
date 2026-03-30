package types

import (
	"context"

	"connectrpc.com/connect"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
)

type IdentityRepository interface {
	GetUser(context.Context, string) (*identityv1.User, error)
	ListUsers(context.Context, *identityv1.ListUsersRequest) ([]*identityv1.User, error)
	CreateUser(context.Context, *identityv1.CreateUserRequest) (*identityv1.User, error)
	UpdateUser(context.Context, *identityv1.UpdateUserRequest) (*identityv1.User, error)
	DeleteUser(context.Context, string) error

	SetUserEnabled(context.Context, string, bool) (*identityv1.User, error)
	SetUserPassword(context.Context, *identityv1.SetUserPasswordRequest) error
	ListUserSessions(context.Context, string) ([]*identityv1.Session, error)
	RevokeUserSession(context.Context, *identityv1.RevokeUserSessionRequest) error
	GetUserMfa(context.Context, string) (*identityv1.MfaStatus, error)
	UpdateUserMfa(context.Context, *identityv1.UpdateUserMfaRequest) (*identityv1.MfaStatus, error)

	AssignRolesToUser(context.Context, *identityv1.AssignRolesToUserRequest) ([]*identityv1.Role, error)
	RemoveRolesFromUser(context.Context, *identityv1.RemoveRolesFromUserRequest) ([]*identityv1.Role, error)
	AddUserToGroups(context.Context, *identityv1.AddUserToGroupsRequest) ([]*identityv1.Group, error)
	RemoveUserFromGroups(context.Context, *identityv1.RemoveUserFromGroupsRequest) ([]*identityv1.Group, error)

	GetRole(context.Context, string) (*identityv1.Role, error)
	ListRoles(context.Context, *identityv1.ListRolesRequest) ([]*identityv1.Role, error)
	CreateRole(context.Context, *identityv1.CreateRoleRequest) (*identityv1.Role, error)
	UpdateRole(context.Context, *identityv1.UpdateRoleRequest) (*identityv1.Role, error)
	DeleteRole(context.Context, string) error

	GetGroup(context.Context, string) (*identityv1.Group, error)
	ListGroups(context.Context, *identityv1.ListGroupsRequest) ([]*identityv1.Group, error)
	CreateGroup(context.Context, *identityv1.CreateGroupRequest) (*identityv1.Group, error)
	UpdateGroup(context.Context, *identityv1.UpdateGroupRequest) (*identityv1.Group, error)
	DeleteGroup(context.Context, string) error

	AssignRolesToGroup(context.Context, *identityv1.AssignRolesToGroupRequest) ([]*identityv1.Role, error)
	RemoveRolesFromGroup(context.Context, *identityv1.RemoveRolesFromGroupRequest) ([]*identityv1.Role, error)
}

type IdentityService interface {
	GetUser(context.Context, *connect.Request[identityv1.GetUserRequest]) (*identityv1.GetUserResponse, error)
	ListUsers(context.Context, *connect.Request[identityv1.ListUsersRequest]) (*identityv1.ListUsersResponse, error)
	CreateUser(context.Context, *connect.Request[identityv1.CreateUserRequest]) (*identityv1.CreateUserResponse, error)
	UpdateUser(context.Context, *connect.Request[identityv1.UpdateUserRequest]) (*identityv1.UpdateUserResponse, error)
	DeleteUser(context.Context, *connect.Request[identityv1.DeleteUserRequest]) (*identityv1.DeleteUserResponse, error)
	SetUserEnabled(context.Context, *connect.Request[identityv1.SetUserEnabledRequest]) (*identityv1.SetUserEnabledResponse, error)
	SetUserPassword(context.Context, *connect.Request[identityv1.SetUserPasswordRequest]) (*identityv1.SetUserPasswordResponse, error)
	ListUserSessions(context.Context, *connect.Request[identityv1.ListUserSessionsRequest]) (*identityv1.ListUserSessionsResponse, error)
	RevokeUserSession(context.Context, *connect.Request[identityv1.RevokeUserSessionRequest]) (*identityv1.RevokeUserSessionResponse, error)
	GetUserMfa(context.Context, *connect.Request[identityv1.GetUserMfaRequest]) (*identityv1.GetUserMfaResponse, error)
	UpdateUserMfa(context.Context, *connect.Request[identityv1.UpdateUserMfaRequest]) (*identityv1.UpdateUserMfaResponse, error)
	AssignRolesToUser(context.Context, *connect.Request[identityv1.AssignRolesToUserRequest]) (*identityv1.AssignRolesToUserResponse, error)
	RemoveRolesFromUser(context.Context, *connect.Request[identityv1.RemoveRolesFromUserRequest]) (*identityv1.RemoveRolesFromUserResponse, error)
	AddUserToGroups(context.Context, *connect.Request[identityv1.AddUserToGroupsRequest]) (*identityv1.AddUserToGroupsResponse, error)
	RemoveUserFromGroups(context.Context, *connect.Request[identityv1.RemoveUserFromGroupsRequest]) (*identityv1.RemoveUserFromGroupsResponse, error)
	GetRole(context.Context, *connect.Request[identityv1.GetRoleRequest]) (*identityv1.GetRoleResponse, error)
	ListRoles(context.Context, *connect.Request[identityv1.ListRolesRequest]) (*identityv1.ListRolesResponse, error)
	CreateRole(context.Context, *connect.Request[identityv1.CreateRoleRequest]) (*identityv1.CreateRoleResponse, error)
	UpdateRole(context.Context, *connect.Request[identityv1.UpdateRoleRequest]) (*identityv1.UpdateRoleResponse, error)
	DeleteRole(context.Context, *connect.Request[identityv1.DeleteRoleRequest]) (*identityv1.DeleteRoleResponse, error)
	GetGroup(context.Context, *connect.Request[identityv1.GetGroupRequest]) (*identityv1.GetGroupResponse, error)
	ListGroups(context.Context, *connect.Request[identityv1.ListGroupsRequest]) (*identityv1.ListGroupsResponse, error)
	CreateGroup(context.Context, *connect.Request[identityv1.CreateGroupRequest]) (*identityv1.CreateGroupResponse, error)
	UpdateGroup(context.Context, *connect.Request[identityv1.UpdateGroupRequest]) (*identityv1.UpdateGroupResponse, error)
	DeleteGroup(context.Context, *connect.Request[identityv1.DeleteGroupRequest]) (*identityv1.DeleteGroupResponse, error)
	AssignRolesToGroup(context.Context, *connect.Request[identityv1.AssignRolesToGroupRequest]) (*identityv1.AssignRolesToGroupResponse, error)
	RemoveRolesFromGroup(context.Context, *connect.Request[identityv1.RemoveRolesFromGroupRequest]) (*identityv1.RemoveRolesFromGroupResponse, error)
}
