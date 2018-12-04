package handler

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/cas"
)

type Role struct {
	BaseHandler
}

func NewRoleHandler(c client.Client) *Role {
	return &Role{}
}

func (role *Role) CreateRole(req *restful.Request, rsp *restful.Response) {
	// cas.Auth.

	params := struct {
		Role string `json:"role"`
		Path string `json:"path"`
		Method string `json:"method"`
	}{}

	if err := role.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	fmt.Println(params, "角色参数")

	cas.Auth.CasEnforcer().AddPolicy(params.Role, params.Path, params.Method)
	cas.Auth.SavePolicy()

	WriteJsonResponse(rsp, nil)
}

// CreateAccess 创建权限
func (role *Role) CreatePermission(req *restful.Request, rsp *restful.Response) {
	params := struct {
		Role string `json:"role"`
		Path string `json:"path"`
		Method string `json:"method"`
	}{}

	if err := role.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	cas.Auth.CasEnforcer().AddPolicy(params.Role, params.Path, params.Method)
	cas.Auth.SavePolicy()

	WriteJsonResponse(rsp, nil)
}

func (role *Role) CheckTest(req *restful.Request, rsp *restful.Response) {
	params := struct {
		Role string `json:"role" query:"role"`
		Path string `json:"path" query:"path"`
		Method string `json:"method" query:"method"`
	}{}

	if err := role.BindAndValidate(req, &params); err != nil {
		WriteError(rsp, err)
		return
	}

	if cas.Auth.Enforce(params.Role, params.Path, params.Method) {
		WriteJsonResponse(rsp, D{
			"msg": "授权通过",
		})
		return
	}

	WriteJsonResponse(rsp, D{
		"msg": "授权失败",
	})
	return
}