package handler

import (
	"context"

	go_api "github.com/micro/go-api/proto"
)

type UserHandler struct {
	BaseHandler
}

func (u *UserHandler) Info(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UpdateCurrentOrg(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) ForgetPassword(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) AllUser(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetCurrentOrganization(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetAllMembers(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) CreateMember(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) RemvoeMemberFromOrg(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UpdateMember(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetMemberDetails(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) RemoveOrg(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UpdateUserInfo(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) FlagMemberState(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) GetHotUsers(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UploadAvatar(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) ExportList(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) DownloadExportTemplate(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}

func (u *UserHandler) UploadUserList(context.Context, *go_api.Request, *go_api.Response) error {
	panic("not implemented")
}
