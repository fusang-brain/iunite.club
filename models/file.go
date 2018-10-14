package models

import (
	"github.com/iron-kit/monger"
	storagePB "iunite.club/services/storage/proto"
)

type File struct {
	monger.Schema `json:",inline" bson:",inline"`

	FileKey          string `json:"file_key,omitempty"`
	Storage          string `json:"storage,omitempty"` // qiniu
	Ext              string `json:"ext,omitempty"`
	Host             string `json:"host,omitempty"`
	Bulket           string `json:"bulket,omitempty"`
	OriginalFilename string `json:"original_filename,omitempty"`
	Size             int64  `json:"size,omitempty"`
	Path             string `json:"path,omitempty"`
}

func (f *File) ToPB() *storagePB.FilePB {
	return &storagePB.FilePB{
		ID:               f.ID.Hex(),
		FileKey:          f.FileKey,
		Storage:          f.Storage,
		Ext:              f.Ext,
		Host:             f.Host,
		Bulket:           f.Bulket,
		OriginalFilename: f.OriginalFilename,
		Size:             f.Size,
		Path:             f.Path,
	}
}
