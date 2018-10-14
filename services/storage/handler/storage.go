package handler

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/qiniu/api.v7/auth/qbox"

	"github.com/qiniu/api.v7/storage"

	"iunite.club/models"
	"iunite.club/services/storage/proto"

	ironic "github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	uuid "github.com/satori/go.uuid"
)

type Storage struct {
	ironic.BaseHandler
}

const (
	AccessKey   = "AOFW65TP9sFumutNR94ahiOw0BpoCkeIqXKNmZ0d"
	SecretKey   = "FUj1hdiAUHarlFtPRP3PnetZJ9nIZjvnnShu8oYp"
	QiniuDomain = "storage.iunite.club"
)

func (s *Storage) Model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (s *Storage) updateToQiniu(file *models.File, content []byte) error {
	putPolicy := storage.PutPolicy{
		Scope: file.Bulket,
	}

	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": file.OriginalFilename,
		},
	}
	filep := bytes.NewReader(content)
	fileSize := int64(len(content))

	if err := formUploader.Put(context.Background(), &ret, upToken, file.FileKey, filep, fileSize, &putExtra); err != nil {
		return err
	}

	return nil

}

func (s *Storage) SaveFileInfo(ctx context.Context, req *iunite_club_srv_storage.FileInfoRequest, rsp *iunite_club_srv_storage.FileResponse) error {
	FileModel := s.Model(ctx, "File")
	sendFile := req.File
	file := new(models.File)

	file.FileKey = sendFile.FileKey
	file.Size = sendFile.Size
	file.Host = sendFile.Host
	file.Ext = sendFile.Ext
	file.OriginalFilename = sendFile.OriginalFilename
	file.Path = sendFile.Path
	file.Bulket = sendFile.Bulket
	file.Storage = sendFile.Storage
	FileModel.Create(file)
	rsp.OK = true
	rsp.File = file.ToPB()

	return nil
}

func (s *Storage) SaveFile(ctx context.Context, req *iunite_club_srv_storage.FileRequest, rsp *iunite_club_srv_storage.FileResponse) error {
	FileModel := s.Model(ctx, "File")
	sendFile := req.File
	file := new(models.File)

	filename := sendFile.Filename
	filenameArr := strings.Split(filename, ".")
	ext := filenameArr[len(filenameArr)-1]
	file.FileKey = fmt.Sprintf("%s.%s", uuid.NewV1().String(), ext)
	file.Size = sendFile.Size
	file.Host = QiniuDomain
	file.Ext = ext
	file.OriginalFilename = filename
	file.Path = "/" + filename
	file.Bulket = "unite"
	file.Storage = "qiniu"

	if err := s.updateToQiniu(file, sendFile.Content); err != nil {
		return s.Error(ctx).BadRequest(err.Error())
	}

	if err := FileModel.Create(file); err != nil {
		return s.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	rsp.File = file.ToPB()
	return nil
}

func (s *Storage) SaveFiles(ctx context.Context, req *iunite_club_srv_storage.FilesRequest, rsp *iunite_club_srv_storage.RepeatedFileResponse) error {
	FileModel := s.Model(ctx, "File")
	lf := len(req.Files)
	if lf < 0 {
		return s.Error(ctx).BadRequest("files can't be none")
	}

	// files := make([]models.File, 0, lf)
	resFiles := make([]*iunite_club_srv_storage.FilePB, 0, lf)
	for _, sendFile := range req.Files {
		file := new(models.File)

		filename := sendFile.Filename
		filenameArr := strings.Split(filename, ".")
		ext := filenameArr[len(filenameArr)-1]
		file.FileKey = fmt.Sprintf("%s.%s", uuid.NewV1().String(), ext)
		file.Size = sendFile.Size
		file.Host = QiniuDomain
		file.Ext = ext
		file.OriginalFilename = filename
		file.Path = "/" + filename
		file.Bulket = "unite"
		file.Storage = "qiniu"

		if err := s.updateToQiniu(file, sendFile.Content); err != nil {
			return s.Error(ctx).BadRequest(err.Error())
		}

		if err := FileModel.Create(file); err != nil {
			return s.Error(ctx).BadRequest(err.Error())
		}
		resFiles = append(resFiles, file.ToPB())
	}
	rsp.OK = true
	rsp.Files = resFiles
	return nil
}
