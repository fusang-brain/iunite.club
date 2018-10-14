package handler

import (
	"context"

	go_api "github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"

	"iunite.club/services/restful/dto"
	"iunite.club/services/restful/file"
	storagePB "iunite.club/services/storage/proto"
)

type FileHandler struct {
	BaseHandler
	fileService storagePB.StorageService
}

func NewFileHandler(c client.Client) *FileHandler {
	return &FileHandler{
		fileService: storagePB.NewStorageService(StorageService, c),
	}
}

func (f *FileHandler) UploadSingleFile(req *go_api.Request, rsp *go_api.Response) {
	ctx := context.Background()
	request := req.Request
	_, fh, err := request.FormFile("file")
	if err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
		return
	}

	file, err := file.SaveFile(fh)

	if err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
		return
	}

	type savedResp struct {
		fileResp *storagePB.FileResponse
		err      error
	}

	ch := make(chan savedResp, 1)

	go func() {
		res, err := f.fileService.SaveFileInfo(ctx, &storagePB.FileInfoRequest{File: file})

		ch <- savedResp{fileResp: res, err: err}
	}()

	reply := <-ch

	if reply.err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(reply.err.Error()))
		return
	}

	SuccessResponse(rsp, D{
		"File": dto.PBToFile(reply.fileResp.File),
	})

}

func (f *FileHandler) UploadMutipartFile(req *go_api.Request, rsp *go_api.Response) {
	request := req.Request
	ctx := context.Background()
	if err := request.ParseMultipartForm(32 << 20); err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
		return
	}

	files := request.MultipartForm.File["files"]
	// filePBs := make([]*storagePB.FilePB, 0, len(files))
	fileRes := make([]*dto.File, 0, len(files))
	for _, fh := range files {

		file, err := file.SaveFile(fh)

		if err != nil {
			// ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
			// return
			continue
		}

		type savedResp struct {
			fileResp *storagePB.FileResponse
			err      error
		}

		ch := make(chan savedResp, 1)

		go func() {
			res, err := f.fileService.SaveFileInfo(ctx, &storagePB.FileInfoRequest{File: file})

			ch <- savedResp{fileResp: res, err: err}
		}()

		reply := <-ch

		if reply.err != nil {
			// ErrorResponse(rsp, f.Error().BadRequest(reply.err.Error()))
			continue
		}

		fileRes = append(fileRes, dto.PBToFile(reply.fileResp.File))
	}

	SuccessResponse(rsp, D{
		"Files": fileRes,
	})
}
