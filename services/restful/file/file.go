package file

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	uuid "github.com/satori/go.uuid"

	storagePB "iunite.club/services/storage/proto"
)

const (
	AccessKey   = "AOFW65TP9sFumutNR94ahiOw0BpoCkeIqXKNmZ0d"
	SecretKey   = "FUj1hdiAUHarlFtPRP3PnetZJ9nIZjvnnShu8oYp"
	QiniuDomain = "storage.iunite.club"
)

func SaveFile(f *multipart.FileHeader) (*storagePB.FilePB, error) {
	fileSize := f.Size
	filename := f.Filename

	filenameArr := strings.Split(filename, ".")
	ext := filenameArr[len(filenameArr)-1]

	file := new(storagePB.FilePB)
	file.FileKey = fmt.Sprintf("%s.%s", uuid.NewV1().String(), ext)
	file.Size = fileSize
	file.Host = QiniuDomain
	file.Ext = ext
	file.OriginalFilename = filename
	file.Path = "/" + filename
	file.Bulket = "unite"
	file.Storage = "qiniu"
	fp, err := f.Open()
	if err != nil {
		return nil, err
	}

	if err := uploadToQiniu(file, fp); err != nil {
		return nil, err
	}

	defer func() {
		fp.Close()
	}()

	return file, nil
}

func uploadToQiniu(file *storagePB.FilePB, reader io.Reader) error {
	putPolicy := storage.PutPolicy{
		Scope: file.Bulket,
	}
	fileSize := file.Size
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
	// filep := bytes.NewReader(content)
	// fileSize := int64(len(content))
	// formUploader.Put()
	if err := formUploader.Put(context.Background(), &ret, upToken, file.FileKey, reader, fileSize, &putExtra); err != nil {
		return err
	}

	return nil
}
