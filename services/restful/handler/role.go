package handler

import (
	"context"
	"github.com/micro/go-micro/client"
	pb "iunite.club/services/core/proto/role"
)

type RoleHandler struct {
	BaseHandler

	roleService pb.RoleService
}

func NewRoleHandler(c client.Client) *RoleHandler {
	return &RoleHandler{
		roleService: pb.NewRoleService(CoreService, c),
	}
}

func (hdl *RoleHandler) UpdateRoleOrGroup() error {
	ctx := context.Background()

	//params := struct {
	//
	//}{}
}