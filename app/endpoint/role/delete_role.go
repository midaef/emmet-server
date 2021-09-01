package role

import (
	"context"
	"github.com/midaef/emmet-server/extra/role"
)

func (r *RoleEndpoint) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.RoleMessageResponse, error) {
	return &role.RoleMessageResponse{}, nil
}
