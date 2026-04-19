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

func (server *identityServer) GetUser(ctx context.Context, req *connect.Request[identityv1.GetUserRequest]) (*connect.Response[identityv1.GetUserResponse], error) {
	resp, err := server.service.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) ListUsers(ctx context.Context, req *connect.Request[identityv1.ListUsersRequest]) (*connect.Response[identityv1.ListUsersResponse], error) {
	resp, err := server.service.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) CreateUser(ctx context.Context, req *connect.Request[identityv1.CreateUserRequest]) (*connect.Response[identityv1.CreateUserResponse], error) {
	resp, err := server.service.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) UpdateUser(ctx context.Context, req *connect.Request[identityv1.UpdateUserRequest]) (*connect.Response[identityv1.UpdateUserResponse], error) {
	resp, err := server.service.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) DeleteUser(ctx context.Context, req *connect.Request[identityv1.DeleteUserRequest]) (*connect.Response[identityv1.DeleteUserResponse], error) {
	resp, err := server.service.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) SetUserEnabled(ctx context.Context, req *connect.Request[identityv1.SetUserEnabledRequest]) (*connect.Response[identityv1.SetUserEnabledResponse], error) {
	resp, err := server.service.SetUserEnabled(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) SetUserPassword(ctx context.Context, req *connect.Request[identityv1.SetUserPasswordRequest]) (*connect.Response[identityv1.SetUserPasswordResponse], error) {
	resp, err := server.service.SetUserPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) ListUserSessions(ctx context.Context, req *connect.Request[identityv1.ListUserSessionsRequest]) (*connect.Response[identityv1.ListUserSessionsResponse], error) {
	resp, err := server.service.ListUserSessions(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) RevokeUserSession(ctx context.Context, req *connect.Request[identityv1.RevokeUserSessionRequest]) (*connect.Response[identityv1.RevokeUserSessionResponse], error) {
	resp, err := server.service.RevokeUserSession(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) GetUserMfa(ctx context.Context, req *connect.Request[identityv1.GetUserMfaRequest]) (*connect.Response[identityv1.GetUserMfaResponse], error) {
	resp, err := server.service.GetUserMfa(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) UpdateUserMfa(ctx context.Context, req *connect.Request[identityv1.UpdateUserMfaRequest]) (*connect.Response[identityv1.UpdateUserMfaResponse], error) {
	resp, err := server.service.UpdateUserMfa(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) AssignRolesToUser(ctx context.Context, req *connect.Request[identityv1.AssignRolesToUserRequest]) (*connect.Response[identityv1.AssignRolesToUserResponse], error) {
	resp, err := server.service.AssignRolesToUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) RemoveRolesFromUser(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromUserRequest]) (*connect.Response[identityv1.RemoveRolesFromUserResponse], error) {
	resp, err := server.service.RemoveRolesFromUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) AddUserToGroups(ctx context.Context, req *connect.Request[identityv1.AddUserToGroupsRequest]) (*connect.Response[identityv1.AddUserToGroupsResponse], error) {
	resp, err := server.service.AddUserToGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) RemoveUserFromGroups(ctx context.Context, req *connect.Request[identityv1.RemoveUserFromGroupsRequest]) (*connect.Response[identityv1.RemoveUserFromGroupsResponse], error) {
	resp, err := server.service.RemoveUserFromGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) GetRole(ctx context.Context, req *connect.Request[identityv1.GetRoleRequest]) (*connect.Response[identityv1.GetRoleResponse], error) {
	resp, err := server.service.GetRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) ListRoles(ctx context.Context, req *connect.Request[identityv1.ListRolesRequest]) (*connect.Response[identityv1.ListRolesResponse], error) {
	resp, err := server.service.ListRoles(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) CreateRole(ctx context.Context, req *connect.Request[identityv1.CreateRoleRequest]) (*connect.Response[identityv1.CreateRoleResponse], error) {
	resp, err := server.service.CreateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) UpdateRole(ctx context.Context, req *connect.Request[identityv1.UpdateRoleRequest]) (*connect.Response[identityv1.UpdateRoleResponse], error) {
	resp, err := server.service.UpdateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) DeleteRole(ctx context.Context, req *connect.Request[identityv1.DeleteRoleRequest]) (*connect.Response[identityv1.DeleteRoleResponse], error) {
	resp, err := server.service.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) GetGroup(ctx context.Context, req *connect.Request[identityv1.GetGroupRequest]) (*connect.Response[identityv1.GetGroupResponse], error) {
	resp, err := server.service.GetGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) ListGroups(ctx context.Context, req *connect.Request[identityv1.ListGroupsRequest]) (*connect.Response[identityv1.ListGroupsResponse], error) {
	resp, err := server.service.ListGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) CreateGroup(ctx context.Context, req *connect.Request[identityv1.CreateGroupRequest]) (*connect.Response[identityv1.CreateGroupResponse], error) {
	resp, err := server.service.CreateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) UpdateGroup(ctx context.Context, req *connect.Request[identityv1.UpdateGroupRequest]) (*connect.Response[identityv1.UpdateGroupResponse], error) {
	resp, err := server.service.UpdateGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) DeleteGroup(ctx context.Context, req *connect.Request[identityv1.DeleteGroupRequest]) (*connect.Response[identityv1.DeleteGroupResponse], error) {
	resp, err := server.service.DeleteGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) AssignRolesToGroup(ctx context.Context, req *connect.Request[identityv1.AssignRolesToGroupRequest]) (*connect.Response[identityv1.AssignRolesToGroupResponse], error) {
	resp, err := server.service.AssignRolesToGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (server *identityServer) RemoveRolesFromGroup(ctx context.Context, req *connect.Request[identityv1.RemoveRolesFromGroupRequest]) (*connect.Response[identityv1.RemoveRolesFromGroupResponse], error) {
	resp, err := server.service.RemoveRolesFromGroup(ctx, req)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func NewIdentityServer(service types.IdentityService) v1connect.IdentityServiceHandler {
	return &identityServer{service: service}
}
