package handler

import (
	"context"
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"iunite.club/services/restful/dto"
	"iunite.club/services/restful/file"
	"iunite.club/services/storage/proto"
	cloudPB "iunite.club/services/storage/proto/cloud"
)

type CloudHandler struct {
	BaseHandler

	cloudService cloudPB.CloudDiskService
	fileService  iunite_club_srv_storage.StorageService
}

type (
	savedFileInfoReply struct {
		FileResponse *iunite_club_srv_storage.FileResponse
		Err          error
	}
)

func NewCloudService(c client.Client) *CloudHandler {
	return &CloudHandler{
		cloudService: cloudPB.NewCloudDiskService(StorageService, c),
		fileService:  iunite_club_srv_storage.NewStorageService(StorageService, c),
	}
}

func (self *CloudHandler) DownloadFile(req *restful.Request, rsp *restful.Response) {
	// panic("The function has not Impl")
	ctx := context.Background()
	params := struct {
		ID string `query:"id" validate:"nonzero,objectid"`
	}{}

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	fileResp, err := self.cloudService.FindFile(ctx, &cloudPB.ByFileID{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	f := fileResp.File

	// fileResp.File.
	link := fmt.Sprintf("http://%s/%s?attname=", f.Host, f.FileKey)

	http.Redirect(rsp.ResponseWriter, req.Request, link, http.StatusSeeOther)
	return
}

func (self *CloudHandler) ShowFile(req *restful.Request, rsp *restful.Response) {
	// panic("The function has not Impl")
	ctx := context.Background()
	id := req.PathParameter("id")

	fileResp, err := self.cloudService.FindFile(ctx, &cloudPB.ByFileID{
		ID: id,
	})

	if err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	f := fileResp.File

	// fileResp.File.
	link := fmt.Sprintf("http://%s/%s", f.Host, f.FileKey)

	http.Redirect(rsp.ResponseWriter, req.Request, link, http.StatusSeeOther)
	return
}

func (self *CloudHandler) saveFileInfo(ctx context.Context, in *iunite_club_srv_storage.FileInfoRequest) chan savedFileInfoReply {
	ch := make(chan savedFileInfoReply, 1)

	go func() {
		resp, err := self.fileService.SaveFileInfo(ctx, in)

		ch <- savedFileInfoReply{
			FileResponse: resp,
			Err:          err,
		}
	}()

	return ch
}

// func (self *CloudHandler)

func (self *CloudHandler) UploadFile(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	request := req.Request
	parentID := request.FormValue("parentID")
	org := request.FormValue("org")

	_, fh, err := request.FormFile("file")
	if err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	type savedResp struct {
		filePB *iunite_club_srv_storage.FilePB
		err    error
	}
	// 并发上传文件
	saveCh := make(chan savedResp, 1)
	go func() {
		filePB, err := file.SaveFile(fh)

		saveCh <- savedResp{
			filePB: filePB,
			err:    err,
		}
	}()

	savedReply := <-saveCh

	fmt.Println(savedReply.filePB, "file pb info")

	if savedReply.err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(savedReply.err.Error()))
		return
	}

	savedFileInfoCh := self.saveFileInfo(ctx, &iunite_club_srv_storage.FileInfoRequest{
		File: savedReply.filePB,
	})
	saveFileInfoReply := <-savedFileInfoCh

	if saveFileInfoReply.Err != nil {
		ErrorResponse(rsp, saveFileInfoReply.Err)
		return
	}
	fileInfo := saveFileInfoReply.FileResponse.File
	createdResp, err := self.cloudService.CreateItem(ctx, &cloudPB.WithItemBundle{
		Name:     fileInfo.OriginalFilename,
		ParentID: parentID,
		ClubID:   org,
		FileID:   fileInfo.ID,
		Kind:     1,
		UserID:   self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	// createdResp.CloudItem =
	SuccessResponse(rsp, D{
		"DiskItem": dto.PBToCloudDisk(createdResp.CloudItem),
	})
	return
}

func (self *CloudHandler) CreateDIR(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		Name         string   `json:"name"`
		ParentID     string   `json:"parentID"`
		Organization string   `json:"org"`
		DRange       []string `json:"range"`
	}{}

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	resp, err := self.cloudService.CreateItem(ctx, &cloudPB.WithItemBundle{
		Name:        params.Name,
		ParentID:    params.ParentID,
		ClubID:      params.Organization,
		Kind:        0,
		UserID:      self.GetUserIDFromRequest(req),
		Departments: params.DRange,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"DiskItem": dto.PBToCloudDisk(resp.CloudItem),
	})
}

func (self *CloudHandler) List(req *restful.Request, rsp *restful.Response) {
	params := struct {
		Page         int64  `query:"page"`
		Limit        int64  `query:"limit"`
		Organization string `query:"org" validate:"nonzero,objectid"`
		ParentID     string `query:"pid" validate:"objectid"`
		Search       string `query:"search"`
	}{}
	ctx := context.Background()
	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	listResp, err := self.cloudService.List(ctx, &cloudPB.ByParentAndClubID{
		ClubID:   params.Organization,
		ParentID: params.ParentID,
		UserID:   self.GetUserIDFromRequest(req),
		Page:     int32(params.Page),
		Limit:    int32(params.Limit),
		Search:   params.Search,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	items := make([]*dto.CloudDisk, 0)

	for _, v := range listResp.Items {
		items = append(items, dto.PBToCloudDisk(v))
	}

	SuccessResponseWithPage(rsp, params.Page, params.Limit, listResp.Total, items)
}

func (self *CloudHandler) Details(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		ID string `json:"id" query:"id"`
	}{}

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	resp, err := self.cloudService.GetDetails(ctx, &cloudPB.ByID{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	// panic("The function has not Impl")

	SuccessResponse(rsp, D{
		"Details": dto.PBToCloudDisk(resp.Item),
	})
}

func (self *CloudHandler) UpdatePermission(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID          string   `json:"id"`
		Departments []string `json:"range"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	_, err := self.cloudService.UpdatePermission(ctx, &cloudPB.WithDepartmentsByFileID{
		ID:          params.ID,
		Departments: params.Departments,
		UserID:      self.GetUserIDFromRequest(req),
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}
