package services

import (
	"context"

	"connectrpc.com/connect"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
	"github.com/charmmtech/identity/internal/types"
)

type identityService struct {
	repo types.IdentityRepository
}

func (s *identityService) GetUser(ctx context.Context, req *connect.Request[identityv1.GetUserRequest]) (*identityv1.GetUserResponse, error) {
	user, err := s.repo.GetUser(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetUserResponse{User: user}, nil
}

func (s *identityService) ListUsers(ctx context.Context, req *connect.Request[identityv1.ListUsersRequest]) (*identityv1.ListUsersResponse, error) {
	users, err := s.repo.ListUsers(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListUsersResponse{Users: users}, nil
}

func (s *identityService) CreateUser(ctx context.Context, req *connect.Request[identityv1.CreateUserRequest]) (*identityv1.CreateUserResponse, error) {
	user, err := s.repo.CreateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateUserResponse{User: user}, nil
}

func (s *identityService) UpdateUser(ctx context.Context, req *connect.Request[identityv1.UpdateUserRequest]) (*identityv1.UpdateUserResponse, error) {
	user, err := s.repo.UpdateUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateUserResponse{User: user}, nil
}

func (s *identityService) DeleteUser(ctx context.Context, req *connect.Request[identityv1.DeleteUserRequest]) (*identityv1.DeleteUserResponse, error) {
	if err := s.repo.DeleteUser(ctx, req.Msg.GetId()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteUserResponse{}, nil
}

func (s *identityService) SetUserEnabled(ctx context.Context, req *connect.Request[identityv1.SetUserEnabledRequest]) (*identityv1.SetUserEnabledResponse, error) {
	user, err := s.repo.SetUserEnabled(ctx, req.Msg.GetId(), req.Msg.GetEnabled())
	if err != nil {
		return nil, err
	}
	return &identityv1.SetUserEnabledResponse{User: user}, nil
}

func (s *identityService) SetUserPassword(ctx context.Context, req *connect.Request[identityv1.SetUserPasswordRequest]) (*identityv1.SetUserPasswordResponse, error) {
	if err := s.repo.SetUserPassword(ctx, req.Msg); err != nil {
		return nil, err
	}
	return &identityv1.SetUserPasswordResponse{}, nil
}

func (s *identityService) ListUserSessions(ctx context.Context, req *connect.Request[identityv1.ListUserSessionsRequest]) (*identityv1.ListUserSessionsResponse, error) {
	sessions, err := s.repo.ListUserSessions(ctx, req.Msg.GetUserId())
	if err != nil {
		return nil, err
	}
	return &identityv1.ListUserSessionsResponse{Sessions: sessions}, nil
}

func (s *identityService) RevokeUserSession(ctx context.Context, req *connect.Request[identityv1.RevokeUserSessionRequest]) (*identityv1.RevokeUserSessionResponse, error) {
	if err := s.repo.RevokeUserSession(ctx, req.Msg); err != nil {
		return nil, err
	}
	return &identityv1.RevokeUserSessionResponse{}, nil
}

func (s *identityService) GetUserMfa(ctx context.Context, req *connect.Request[identityv1.GetUserMfaRequest]) (*identityv1.GetUserMfaResponse, error) {
	status, err := s.repo.GetUserMfa(ctx, req.Msg.GetUserId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetUserMfaResponse{Status: status}, nil
}

func (s *identityService) UpdateUserMfa(ctx context.Context, req *connect.Request[identityv1.UpdateUserMfaRequest]) (*identityv1.UpdateUserMfaResponse, error) {
	status, err := s.repo.UpdateUserMfa(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateUserMfaResponse{Status: status}, nil
}

func (s *identityService) AssignRolesToUser(ctx context.Context, req *connect.Request[identityv1.AssignRolesToUserRequest]) (*identityv1.AssignRolesToUserResponse, error) {
	roles, err := s.repo.AssignRolesToUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AssignRolesToUserResponse{Roles: roles}, nil
}

func (s *identityService) RemoveRolesFromUser(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromUserRequest]) (*identityv1.RemoveRolesFromUserResponse, error) {
	roles, err := s.repo.RemoveRolesFromUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveRolesFromUserResponse{Roles: roles}, nil
}

func (s *identityService) AddUserToGroups(ctx context.Context, req *connect.Request[identityv1.AddUserToGroupsRequest]) (*identityv1.AddUserToGroupsResponse, error) {
	groups, err := s.repo.AddUserToGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AddUserToGroupsResponse{Groups: groups}, nil
}

func (s *identityService) RemoveUserFromGroups(ctx context.Context, req *connect.Request[identityv1.RemoveUserFromGroupsRequest]) (*identityv1.RemoveUserFromGroupsResponse, error) {
	groups, err := s.repo.RemoveUserFromGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveUserFromGroupsResponse{Groups: groups}, nil
}

func (s *identityService) GetRole(ctx context.Context, req *connect.Request[identityv1.GetRoleRequest]) (*identityv1.GetRoleResponse, error) {
	role, err := s.repo.GetRole(ctx, req.Msg.GetRoleName())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetRoleResponse{Role: role}, nil
}

func (s *identityService) ListRoles(ctx context.Context, req *connect.Request[identityv1.ListRolesRequest]) (*identityv1.ListRolesResponse, error) {
	roles, err := s.repo.ListRoles(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListRolesResponse{Roles: roles}, nil
}

func (s *identityService) CreateRole(ctx context.Context, req *connect.Request[identityv1.CreateRoleRequest]) (*identityv1.CreateRoleResponse, error) {
	role, err := s.repo.CreateRole(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateRoleResponse{Role: role}, nil
}

func (s *identityService) UpdateRole(ctx context.Context, req *connect.Request[identityv1.UpdateRoleRequest]) (*identityv1.UpdateRoleResponse, error) {
	role, err := s.repo.UpdateRole(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateRoleResponse{Role: role}, nil
}

func (s *identityService) DeleteRole(ctx context.Context, req *connect.Request[identityv1.DeleteRoleRequest]) (*identityv1.DeleteRoleResponse, error) {
	if err := s.repo.DeleteRole(ctx, req.Msg.GetRoleName()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteRoleResponse{}, nil
}

func (s *identityService) GetGroup(ctx context.Context, req *connect.Request[identityv1.GetGroupRequest]) (*identityv1.GetGroupResponse, error) {
	group, err := s.repo.GetGroup(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}
	return &identityv1.GetGroupResponse{Group: group}, nil
}

func (s *identityService) ListGroups(ctx context.Context, req *connect.Request[identityv1.ListGroupsRequest]) (*identityv1.ListGroupsResponse, error) {
	groups, err := s.repo.ListGroups(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.ListGroupsResponse{Groups: groups}, nil
}

func (s *identityService) CreateGroup(ctx context.Context, req *connect.Request[identityv1.CreateGroupRequest]) (*identityv1.CreateGroupResponse, error) {
	group, err := s.repo.CreateGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.CreateGroupResponse{Group: group}, nil
}

func (s *identityService) UpdateGroup(ctx context.Context, req *connect.Request[identityv1.UpdateGroupRequest]) (*identityv1.UpdateGroupResponse, error) {
	group, err := s.repo.UpdateGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.UpdateGroupResponse{Group: group}, nil
}

func (s *identityService) DeleteGroup(ctx context.Context, req *connect.Request[identityv1.DeleteGroupRequest]) (*identityv1.DeleteGroupResponse, error) {
	if err := s.repo.DeleteGroup(ctx, req.Msg.GetId()); err != nil {
		return nil, err
	}
	return &identityv1.DeleteGroupResponse{}, nil
}

func (s *identityService) AssignRolesToGroup(ctx context.Context, req *connect.Request[identityv1.AssignRolesToGroupRequest]) (*identityv1.AssignRolesToGroupResponse, error) {
	roles, err := s.repo.AssignRolesToGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.AssignRolesToGroupResponse{Roles: roles}, nil
}

func (s *identityService) RemoveRolesFromGroup(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromGroupRequest]) (*identityv1.RemoveRolesFromGroupResponse, error) {
	roles, err := s.repo.RemoveRolesFromGroup(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return &identityv1.RemoveRolesFromGroupResponse{Roles: roles}, nil
}

func NewIdentityService(repo types.IdentityRepository) types.IdentityService {
	return &identityService{repo: repo}
}
