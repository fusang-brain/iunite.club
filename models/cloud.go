package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	cloudPB "iunite.club/services/storage/proto/cloud"
)

type Cloud struct {
	monger.Schema `json:",inline" bson:",inline"`

	ClubID        bson.ObjectId   `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Name          string          `json:"name,omitempty" bson:"name,omitempty"`
	OriginalName  string          `json:"original_name,omitempty" bson:"original_name,omitempty"`
	ParentID      bson.ObjectId   `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	Kind          int             `json:"kind,omitempty" bson:"kind,omitempty"`         // 0: 目录 1: 文件
	OwnerID       bson.ObjectId   `json:"owner_id,omitempty" bson:"owner_id,omitempty"` // 创建人(拥有者)
	Owner         *User           `json:"owner,omitempty" bson:"owner,omitempty" monger:"belongTo,foreignKey=owner_id"`
	FileID        bson.ObjectId   `json:"file_id,omitempty" bson:"file_id,omitempty"`
	File          *File           `json:"file,omitempty" bson:"file,omitempty" monger:"belongTo,foreignKey=file_id"`
	EnabledToAll  bool            `json:"enabled_to_all,omitempty" bson:"enabled_to_all,omitempty"`
	UserIDS       []bson.ObjectId `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	DepartmentIDS []bson.ObjectId `json:"department_ids,omitempty" bson:"department_ids,omitempty"`
}

func (self *Cloud) ToPB() *cloudPB.CloudPB {
	res := &cloudPB.CloudPB{
		ID:           self.ID.Hex(),
		CreatedAt:    hptypes.TimestampProto(self.CreatedAt),
		UpdatedAt:    hptypes.TimestampProto(self.UpdatedAt),
		Name:         self.Name,
		OriginalName: self.OriginalName,
		Kind:         int32(self.Kind),
		ParentID:     self.ParentID.Hex(),
		ClubID:       self.ClubID.Hex(),
		OwnerID:      self.OwnerID.Hex(),
		FileID:       self.FileID.Hex(),
		EnabledToAll: self.EnabledToAll,
	}

	if self.Owner != nil {
		res.Owner = self.Owner.ToPB()
	}

	if self.File != nil {
		res.File = self.File.ToPB()
	}

	if len(self.DepartmentIDS) > 0 {
		depts := make([]string, 0)

		for _, v := range self.DepartmentIDS {
			depts = append(depts, v.Hex())
		}

		res.DepartmentIDS = depts
	}

	if len(self.UserIDS) > 0 {
		users := make([]string, 0)
		for _, v := range self.UserIDS {
			users = append(users, v.Hex())
		}

		res.UserIDS = users
	}

	return res
}
