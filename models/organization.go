package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type OrganizationJob struct {
	monger.Document `json:",inline" bson:",inline"`

	Name           string
	Slug           string
	OrganizationID bson.ObjectId
}

type PathIndex struct {
	ID   bson.ObjectId
	Name string
	Slug string
	Sort int
}

type Paperwork struct {
	FileID   bson.ObjectId
	UploadAt time.Time
}

type ClubProfile struct {
	Logo       string
	Scale      uint
	Paperworks []Paperwork
}

type Organization struct {
	monger.Document `json:",inline" bson:",inline"`
	Kind            string            // 组织类型 'club': 社团, 'department': 部门
	Name            string            // 组织架构名称
	Slug            string            // slug
	SchoolID        bson.ObjectId     // 归属学校ID
	Description     string            // 组织描述
	ParentID        bson.ObjectId     // 上级ID
	PathIndexs      []PathIndex       // 路径索引
	ClubProfile     ClubProfile       // 社团信息
	Jobs            []OrganizationJob // 所有的职位
	Users           []User
}
