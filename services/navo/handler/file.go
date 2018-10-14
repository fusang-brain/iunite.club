package handler

import (
	"context"

	"iunite.club/services/navo/dto"

	"iunite.club/services/navo/client"
	storagePB "iunite.club/services/storage/proto"

	go_api "github.com/micro/go-api/proto"
)

type FileHandler struct {
	BaseHandler
}

func (f *FileHandler) UploadSingleFile(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	file, err := f.FormFile(req, "file")
	f.MultipartForm
	if err != nil {
		return ErrorResponse(rsp, f.Error(ctx).BadRequest(err.Error()))
	}

	filep, err := file.Open()

	if err != nil {
		return ErrorResponse(rsp, f.Error(ctx).BadRequest(err.Error()))
	}

	defer func() {
		filep.Close()
	}()

	fileContent := make([]byte, file.Size)

	count, err := filep.Read(fileContent)
	if err != nil {
		return ErrorResponse(rsp, f.Error(ctx).InternalServerError(err.Error()))
	}

	storageSrv, ok := client.StroageServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, f.Error(ctx).InternalServerError("not found storage service"))
	}

	savedResp, err := storageSrv.SaveFile(ctx, &storagePB.FileRequest{
		File: &storagePB.File{
			Filename: file.Filename,
			Size:     int64(count),
			Content:  fileContent[:count],
		},
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	return SuccessResponse(rsp, D{
		"File": dto.PBToFile(savedResp.File),
	})
}

func (f *FileHandler) UploadMutipartFile(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	// panic("not implemented")
	files, err := f.FormFiles(req, "files")
	if err != nil {
		return ErrorResponse(rsp, f.Error(ctx).BadRequest(err.Error()))
	}

	// fileContents := make([][]byte, 0, len(files))
	filePBs := make([]*storagePB.File, 0, len(files))

	for _, file := range files {
		filep, err := file.Open()

		if err != nil {
			return ErrorResponse(rsp, f.Error(ctx).BadRequest(err.Error()))
		}

		fileContent := make([]byte, file.Size)

		count, err := filep.Read(fileContent)
		if err != nil {
			return ErrorResponse(rsp, f.Error(ctx).BadRequest(err.Error()))
		}
		// fileContents = append(fileContents, fileContent[:count])
		filePBs = append(filePBs, &storagePB.File{
			Filename: file.Filename,
			Size:     int64(count),
			Content:  fileContent[:count],
		})
		filep.Close()
	}

	storageSrv, ok := client.StroageServiceFromContext(ctx)

	if !ok {
		return ErrorResponse(rsp, f.Error(ctx).InternalServerError("not found storage service"))
	}

	savedResp, err := storageSrv.SaveFiles(ctx, &storagePB.FilesRequest{
		Files: filePBs,
	})

	if err != nil {
		return ErrorResponse(rsp, err)
	}

	fileRes := make([]*dto.File, 0, len(savedResp.Files))

	for _, f := range savedResp.Files {
		fileRes = append(fileRes, dto.PBToFile(f))
	}
	return SuccessResponse(rsp, D{
		"Files": fileRes,
	})
}
