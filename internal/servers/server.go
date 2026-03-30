package servers

import (
	"context"

	"connectrpc.com/connect"
	identityv1 "github.com/charmmtech/identity/gen/charmmtech/identity/v1"
	"github.com/charmmtech/identity/gen/charmmtech/identity/v1/v1connect"
	"github.com/charmmtech/identity/internal/types"
)

type identityServer struct {
	service types.IdentityService
}

func (service *identityServer) GetUser(ctx context.Context, req *connect.Request[identityv1.GetUserRequest]) (*connect.Response[identityv1.GetUserResponse], error) {
	resp, err := service.service.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) ListUsers(ctx context.Context, req *connect.Request[identityv1.ListUsersRequest]) (*connect.Response[identityv1.ListUsersResponse], error) {
	resp, err := service.service.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) CreateUser(ctx context.Context, req *connect.Request[identityv1.CreateUserRequest]) (*connect.Response[identityv1.CreateUserResponse], error) {
	resp, err := service.service.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) UpdateUser(ctx context.Context, req *connect.Request[identityv1.UpdateUserRequest]) (*connect.Response[identityv1.UpdateUserResponse], error) {
	resp, err := service.service.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) DeleteUser(ctx context.Context, req *connect.Request[identityv1.DeleteUserRequest]) (*connect.Response[identityv1.DeleteUserResponse], error) {
	resp, err := service.service.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) SetUserEnabled(ctx context.Context, req *connect.Request[identityv1.SetUserEnabledRequest]) (*connect.Response[identityv1.SetUserEnabledResponse], error) {
	resp, err := service.service.SetUserEnabled(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) SetUserPassword(ctx context.Context, req *connect.Request[identityv1.SetUserPasswordRequest]) (*connect.Response[identityv1.SetUserPasswordResponse], error) {
	resp, err := service.service.SetUserPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) ListUserSessions(ctx context.Context, req *connect.Request[identityv1.ListUserSessionsRequest]) (*connect.Response[identityv1.ListUserSessionsResponse], error) {
	resp, err := service.service.ListUserSessions(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) RevokeUserSession(ctx context.Context, req *connect.Request[identityv1.RevokeUserSessionRequest]) (*connect.Response[identityv1.RevokeUserSessionResponse], error) {
	resp, err := service.service.RevokeUserSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) GetUserMfa(ctx context.Context, req *connect.Request[identityv1.GetUserMfaRequest]) (*connect.Response[identityv1.GetUserMfaResponse], error) {
	resp, err := service.service.GetUserMfa(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) UpdateUserMfa(ctx context.Context, req *connect.Request[identityv1.UpdateUserMfaRequest]) (*connect.Response[identityv1.UpdateUserMfaResponse], error) {
	resp, err := service.service.UpdateUserMfa(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) AssignRolesToUser(ctx context.Context, req *connect.Request[identityv1.AssignRolesToUserRequest]) (*connect.Response[identityv1.AssignRolesToUserResponse], error) {
	resp, err := service.service.AssignRolesToUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) RemoveRolesFromUser(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromUserRequest]) (*connect.Response[identityv1.RemoveRolesFromUserResponse], error) {
	resp, err := service.service.RemoveRolesFromUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) AddUserToGroups(ctx context.Context, req *connect.Request[identityv1.AddUserToGroupsRequest]) (*connect.Response[identityv1.AddUserToGroupsResponse], error) {
	resp, err := service.service.AddUserToGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) RemoveUserFromGroups(ctx context.Context, req *connect.Request[identityv1.RemoveUserFromGroupsRequest]) (*connect.Response[identityv1.RemoveUserFromGroupsResponse], error) {
	resp, err := service.service.RemoveUserFromGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) GetRole(ctx context.Context, req *connect.Request[identityv1.GetRoleRequest]) (*connect.Response[identityv1.GetRoleResponse], error) {
	resp, err := service.service.GetRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) ListRoles(ctx context.Context, req *connect.Request[identityv1.ListRolesRequest]) (*connect.Response[identityv1.ListRolesResponse], error) {
	resp, err := service.service.ListRoles(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) CreateRole(ctx context.Context, req *connect.Request[identityv1.CreateRoleRequest]) (*connect.Response[identityv1.CreateRoleResponse], error) {
	resp, err := service.service.CreateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) UpdateRole(ctx context.Context, req *connect.Request[identityv1.UpdateRoleRequest]) (*connect.Response[identityv1.UpdateRoleResponse], error) {
	resp, err := service.service.UpdateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) DeleteRole(ctx context.Context, req *connect.Request[identityv1.DeleteRoleRequest]) (*connect.Response[identityv1.DeleteRoleResponse], error) {
	resp, err := service.service.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) GetGroup(ctx context.Context, req *connect.Request[identityv1.GetGroupRequest]) (*connect.Response[identityv1.GetGroupResponse], error) {
	resp, err := service.service.GetGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) ListGroups(ctx context.Context, req *connect.Request[identityv1.ListGroupsRequest]) (*connect.Response[identityv1.ListGroupsResponse], error) {
	resp, err := service.service.ListGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) CreateGroup(ctx context.Context, req *connect.Request[identityv1.CreateGroupRequest]) (*connect.Response[identityv1.CreateGroupResponse], error) {
	resp, err := service.service.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) UpdateGroup(ctx context.Context, req *connect.Request[identityv1.UpdateGroupRequest]) (*connect.Response[identityv1.UpdateGroupResponse], error) {
	resp, err := service.service.UpdateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) DeleteGroup(ctx context.Context, req *connect.Request[identityv1.DeleteGroupRequest]) (*connect.Response[identityv1.DeleteGroupResponse], error) {
	resp, err := service.service.DeleteGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) AssignRolesToGroup(ctx context.Context, req *connect.Request[identityv1.AssignRolesToGroupRequest]) (*connect.Response[identityv1.AssignRolesToGroupResponse], error) {
	resp, err := service.service.AssignRolesToGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (service *identityServer) RemoveRolesFromGroup(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromGroupRequest]) (*connect.Response[identityv1.RemoveRolesFromGroupResponse], error) {
	resp, err := service.service.RemoveRolesFromGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func NewIdentityServer(service types.IdentityService) v1connect.IdentityServiceHandler {
	return &identityServer{service: service}
}
