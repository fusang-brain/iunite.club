package models

import (
	"time"

	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	orgPB "iunite.club/services/organization/proto"
)

// OrganizationJob 组织内职位
type OrganizationJob struct {
	monger.Schema `json:",inline" bson:",inline"`

	Name           string        `json:"name,omitempty" bson:"name,omitempty"`
	Slug           string        `json:"slug,omitempty" bson:"slug,omitempty"`
	OrganizationID bson.ObjectId `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
}

// PathIndex 路径索引
type PathIndex struct {
	ID   bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	Name string        `json:"name,omitempty" bson:"name,omitempty"`
	Slug string        `json:"slug,omitempty" bson:"slug,omitempty"`
	Sort int           `json:"sort,omitempty" bson:"sort,omitempty"`
}

// Paperwork 社团资质
type Paperwork struct {
	FileID   bson.ObjectId `json:"file_id,omitempty" bson:"file_id,omitempty"`
	UploadAt time.Time     `json:"upload_at,omitempty" bson:"upload_at,omitempty"`
}

// ClubProfile 社团信息
type ClubProfile struct {
	Logo       string      `json:"logo,omitempty" bson:"logo,omitempty"`
	Scale      int32       `json:"scale,omitempty" bson:"scale,omitempty"`
	Paperworks []Paperwork `json:"paperworks,omitempty" bson:"paperworks,omitempty"`
}

// Organization 社团
type Organization struct {
	monger.Schema `json:",inline" bson:",inline"`
	Enabled       bool              `json:"enabled" bson:"enabled"`
	Kind          string            `json:"kind,omitempty" bson:"kind,omitempty"`                                             // 组织类型 'club': 社团, 'department': 部门
	Name          string            `json:"name,omitempty" bson:"name,omitempty"`                                             // 组织架构名称
	Slug          string            `json:"slug,omitempty" bson:"slug,omitempty"`                                             // slug
	SchoolID      bson.ObjectId     `json:"school_id,omitempty" bson:"school_id,omitempty"`                                   // 归属学校ID
	Description   string            `json:"description,omitempty" bson:"description,omitempty"`                               // 组织描述
	ParentID      bson.ObjectId     `json:"parent_id,omitempty" bson:"parent_id,omitempty"`                                   // 上级ID
	PathIndexs    []PathIndex       `json:"path_indexs,omitempty" bson:"pathindexs,omitempty"`                                // 路径索引
	ClubProfile   ClubProfile       `json:"club_profile,omitempty" bson:"club_profile,omitempty"`                             // 社团信息
	Jobs          []OrganizationJob `json:"jobs,omitempty" bson:"jobs,omitempty" monger:"hasMany,foreignKey=organization_id"` // 所有的职位
}

// JoinOrganizationAccept 申请加入社团
type OrganizationAccept struct {
	monger.Schema  `json:",inline" bson:",inline"`
	UserID         bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`                                                         // 用户ID
	User           *User         `json:"user,omitempty" bson:"user,omitempty" monger:"belongTo,foreignKey=user_id"`                          // 用户
	OrganizationID bson.ObjectId `json:"organization_id,omitempty" bson:"organization_id,omitempty"`                                         // 组织ID
	Organization   *Organization `json:"organization,omitempty" bson:"organization,omitempty" monger:"belongTo,foreignKey=organization_id" ` // 组织
	State          int           `json:"state,omitempty" bson:"state"`                                                                       // 0: 待审批 1: 已拒绝 2: 已通过
	Kind           int           `json:"kind,omitempty" bson:"kind"`                                                                         // 1: 创建社团申请 2: 加入社团申请
}

func (paperWork *Paperwork) ToPB() *orgPB.Paperwork {
	return &orgPB.Paperwork{
		UploadAt: paperWork.UploadAt.String(),
		FileID:   paperWork.FileID.Hex(),
	}
}

func (clubProfile *ClubProfile) ToPB() *orgPB.ClubProfile {
	res := orgPB.ClubProfile{
		Logo:  clubProfile.Logo,
		Scale: clubProfile.Scale,
	}

	if len(clubProfile.Paperworks) > 0 {
		res.Paperworks = []*orgPB.Paperwork{}

		for _, p := range clubProfile.Paperworks {
			pb := p.ToPB()
			res.Paperworks = append(res.Paperworks, pb)
		}
	}

	return &res
}

func (o *Organization) ToPB() *orgPB.Organization {
	pb := o.ClubProfile.ToPB()
	orgRaw := orgPB.Organization{
		ID:          o.ID.Hex(),
		Kind:        o.Kind,
		Name:        o.Name,
		Slug:        o.Slug,
		SchoolID:    o.SchoolID.Hex(),
		Description: o.Description,
		ParentID:    o.ParentID.Hex(),
		ClubProfile: pb,
		// 	OrganizationID: o.PathIndexs[0].ID.Hex(),
	}

	if len(o.PathIndexs) > 0 {
		// o.PathIndexs[0].
		orgRaw.ClubID = o.PathIndexs[0].ID.Hex()
	}

	return &orgRaw
}

func (j *OrganizationJob) ToPB() *orgPB.Job {
	return &orgPB.Job{
		Name: j.Name,
		ID:   j.ID.Hex(),
		Slug: j.Slug,
	}
}
