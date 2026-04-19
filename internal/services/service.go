package services

import (
	"context"

	"connectrpc.com/connect"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
	"github.com/charmmtech/identity/internal/types"
)

type identityService struct {
	repository types.IdentityRepository
}

func (service *identityService) GetUser(ctx context.Context, req *connect.Request[identityv1.GetUserRequest]) (*identityv1.GetUserResponse, error) {
	user, err := service.repository.GetUser(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetUserResponse{User: user}, nil
}

func (service *identityService) ListUsers(ctx context.Context, req *connect.Request[identityv1.ListUsersRequest]) (*identityv1.ListUsersResponse, error) {
	users, err := service.repository.ListUsers(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListUsersResponse{Users: users}, nil
}

func (service *identityService) CreateUser(ctx context.Context, req *connect.Request[identityv1.CreateUserRequest]) (*identityv1.CreateUserResponse, error) {
	user, err := service.repository.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateUserResponse{User: user}, nil
}

func (service *identityService) UpdateUser(ctx context.Context, req *connect.Request[identityv1.UpdateUserRequest]) (*identityv1.UpdateUserResponse, error) {
	user, err := service.repository.UpdateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateUserResponse{User: user}, nil
}

func (service *identityService) DeleteUser(ctx context.Context, req *connect.Request[identityv1.DeleteUserRequest]) (*identityv1.DeleteUserResponse, error) {
	if err := service.repository.DeleteUser(ctx, req.Msg.GetId()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteUserResponse{}, nil
}

func (service *identityService) SetUserEnabled(ctx context.Context, req *connect.Request[identityv1.SetUserEnabledRequest]) (*identityv1.SetUserEnabledResponse, error) {
	user, err := service.repository.SetUserEnabled(ctx, req.Msg.GetId(), req.Msg.GetEnabled())
	if err != nil {
		return nil, err
	}
	return &identityv1.SetUserEnabledResponse{User: user}, nil
}

func (service *identityService) SetUserPassword(ctx context.Context, req *connect.Request[identityv1.SetUserPasswordRequest]) (*identityv1.SetUserPasswordResponse, error) {
	if err := service.repository.SetUserPassword(ctx, req.Msg); err != nil {
		return nil, err
	}
	return &identityv1.SetUserPasswordResponse{}, nil
}

func (service *identityService) ListUserSessions(ctx context.Context, req *connect.Request[identityv1.ListUserSessionsRequest]) (*identityv1.ListUserSessionsResponse, error) {
	sessions, err := service.repository.ListUserSessions(ctx, req.Msg.GetUserId())
	if err != nil {
		return nil, err
	}
	return &identityv1.ListUserSessionsResponse{Sessions: sessions}, nil
}

func (service *identityService) RevokeUserSession(ctx context.Context, req *connect.Request[identityv1.RevokeUserSessionRequest]) (*identityv1.RevokeUserSessionResponse, error) {
	if err := service.repository.RevokeUserSession(ctx, req.Msg); err != nil {
		return nil, err
	}
	return &identityv1.RevokeUserSessionResponse{}, nil
}

func (service *identityService) GetUserMfa(ctx context.Context, req *connect.Request[identityv1.GetUserMfaRequest]) (*identityv1.GetUserMfaResponse, error) {
	status, err := service.repository.GetUserMfa(ctx, req.Msg.GetUserId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetUserMfaResponse{Status: status}, nil
}

func (service *identityService) UpdateUserMfa(ctx context.Context, req *connect.Request[identityv1.UpdateUserMfaRequest]) (*identityv1.UpdateUserMfaResponse, error) {
	status, err := service.repository.UpdateUserMfa(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateUserMfaResponse{Status: status}, nil
}

func (service *identityService) AssignRolesToUser(ctx context.Context, req *connect.Request[identityv1.AssignRolesToUserRequest]) (*identityv1.AssignRolesToUserResponse, error) {
	roles, err := service.repository.AssignRolesToUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AssignRolesToUserResponse{Roles: roles}, nil
}

func (service *identityService) RemoveRolesFromUser(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromUserRequest]) (*identityv1.RemoveRolesFromUserResponse, error) {
	roles, err := service.repository.RemoveRolesFromUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveRolesFromUserResponse{Roles: roles}, nil
}

func (service *identityService) AddUserToGroups(ctx context.Context, req *connect.Request[identityv1.AddUserToGroupsRequest]) (*identityv1.AddUserToGroupsResponse, error) {
	groups, err := service.repository.AddUserToGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AddUserToGroupsResponse{Groups: groups}, nil
}

func (service *identityService) RemoveUserFromGroups(ctx context.Context, req *connect.Request[identityv1.RemoveUserFromGroupsRequest]) (*identityv1.RemoveUserFromGroupsResponse, error) {
	groups, err := service.repository.RemoveUserFromGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveUserFromGroupsResponse{Groups: groups}, nil
}

func (service *identityService) GetRole(ctx context.Context, req *connect.Request[identityv1.GetRoleRequest]) (*identityv1.GetRoleResponse, error) {
	role, err := service.repository.GetRole(ctx, req.Msg.GetRoleName())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetRoleResponse{Role: role}, nil
}

func (service *identityService) ListRoles(ctx context.Context, req *connect.Request[identityv1.ListRolesRequest]) (*identityv1.ListRolesResponse, error) {
	roles, err := service.repository.ListRoles(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListRolesResponse{Roles: roles}, nil
}

func (service *identityService) CreateRole(ctx context.Context, req *connect.Request[identityv1.CreateRoleRequest]) (*identityv1.CreateRoleResponse, error) {
	role, err := service.repository.CreateRole(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateRoleResponse{Role: role}, nil
}

func (service *identityService) UpdateRole(ctx context.Context, req *connect.Request[identityv1.UpdateRoleRequest]) (*identityv1.UpdateRoleResponse, error) {
	role, err := service.repository.UpdateRole(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateRoleResponse{Role: role}, nil
}

func (service *identityService) DeleteRole(ctx context.Context, req *connect.Request[identityv1.DeleteRoleRequest]) (*identityv1.DeleteRoleResponse, error) {
	if err := service.repository.DeleteRole(ctx, req.Msg.GetRoleName()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteRoleResponse{}, nil
}

func (service *identityService) GetGroup(ctx context.Context, req *connect.Request[identityv1.GetGroupRequest]) (*identityv1.GetGroupResponse, error) {
	group, err := service.repository.GetGroup(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetGroupResponse{Group: group}, nil
}

func (service *identityService) ListGroups(ctx context.Context, req *connect.Request[identityv1.ListGroupsRequest]) (*identityv1.ListGroupsResponse, error) {
	groups, err := service.repository.ListGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListGroupsResponse{Groups: groups}, nil
}

func (service *identityService) CreateGroup(ctx context.Context, req *connect.Request[identityv1.CreateGroupRequest]) (*identityv1.CreateGroupResponse, error) {
	group, err := service.repository.CreateGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateGroupResponse{Group: group}, nil
}

func (service *identityService) UpdateGroup(ctx context.Context, req *connect.Request[identityv1.UpdateGroupRequest]) (*identityv1.UpdateGroupResponse, error) {
	group, err := service.repository.UpdateGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateGroupResponse{Group: group}, nil
}

func (service *identityService) DeleteGroup(ctx context.Context, req *connect.Request[identityv1.DeleteGroupRequest]) (*identityv1.DeleteGroupResponse, error) {
	if err := service.repository.DeleteGroup(ctx, req.Msg.GetId()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteGroupResponse{}, nil
}

func (service *identityService) AssignRolesToGroup(ctx context.Context, req *connect.Request[identityv1.AssignRolesToGroupRequest]) (*identityv1.AssignRolesToGroupResponse, error) {
	roles, err := service.repository.AssignRolesToGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AssignRolesToGroupResponse{Roles: roles}, nil
}

func (service *identityService) RemoveRolesFromGroup(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromGroupRequest]) (*identityv1.RemoveRolesFromGroupResponse, error) {
	roles, err := service.repository.RemoveRolesFromGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveRolesFromGroupResponse{Roles: roles}, nil
}

func NewIdentityService(repo types.IdentityRepository) types.IdentityService {
	return &identityService{repository: repo}
}
