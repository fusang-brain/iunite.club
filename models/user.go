package models

import (
	assistant "github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	kit_iron_srv_user "iunite.club/services/user/proto"
	"time"
)

type SecruityInfo struct {
	AuthType      string `json:"auth_type,omitempty" bson:"auth_type,omitempty"`
	Key           string `json:"key,omitempty" bson:"key,omitempty"`
	Secret        string `json:"secret,omitempty" bson:"secret,omitempty"`
	PlainPassword string `json:"plain_password,omitempty" bson:"plain_password,omitempty"`
}

type User struct {
	monger.Document `json:",inline" bson:",inline"`

	Username      string         `json:"username,omitempty" bson:"username,omitempty"`
	Enabled       bool           `json:"enabled,omitempty" bson:"enabled,omitempty"`
	SchoolID      bson.ObjectId  `json:"schoolID,omitempty" bson:"school_id,omitempty"`
	Profile       *Profile       `json:"profile,omitempty" bson:"profile" monger:"hasOne,foreignKey=user_id"`
	SecruityInfos []SecruityInfo `json:"-" bson:"secruity_infos,omitempty"`
}

type Profile struct {
	monger.Document `json:",inline" bson:",inline"`

	Avatar    string        `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Mobile    string        `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Firstname string        `json:"firstname,omitempty" bson:"first_name,omitempty"`
	Lastname  string        `json:"lastname,omitempty" bson:"last_name,omitempty"`
	Gender    string        `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthday  time.Time     `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Nickname  string        `json:"nickname,omitempty" bson:"nickname,omitempty"`
	UserID    bson.ObjectId `json:"userID,omitempty" bson:"user_id,omitempty"`
	User      *User         `json:"user,omitempty" bson:"user" monger:"belongTo,foreignKey=user_id"`
}

type UserClubProfile struct {
	monger.Document `json:",inline" bson:",inline"`

	State          int              `json:"state,omitempty" bson:"state,omitempty"` // 0: 申请中 1: 在职 2: 离职 3: 重新申请 4: 拒绝加入
	UserID         bson.ObjectId    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OrganizationID bson.ObjectId    `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	Organization   *Organization    `json:"organization,omitempty" bson:"organization,omitempty" monger:"belongTo,foreignKey=organization_id"`
	IsCreator      bool             `json:"is_creator,omitempty" bson:"is_creator,omitempty"`
	IsMaster       bool             `json:"is_master,omitempty" bson:"is_master,omitempty"`
	JoinTime       time.Time        `json:"join_time,omitempty" bson:"join_time,omitempty"`
	LeaveTime      time.Time        `json:"leave_time,omitempty" bson:"leave_time,omitempty"`
	JobID          bson.ObjectId    `json:"job_id,omitempty" bson:"job_id,omitempty"`
	Job            *OrganizationJob `json:"job,omitempty" bson:"job,omitempty" monger:"belongTo,foreignKey=job_id"`
	DepartmentID   bson.ObjectId    `json:"department_id,omitempty" bson:"department_id,omitempty"`
	Department     *Organization    `json:"department,omitempty" bson:"department,omitempty" monger:"belongTo,foreignKey=department_id"`
}

// func (p *Profile) FromPB()

func (p *Profile) ToPB() *kit_iron_srv_user.Profile {
	pb := kit_iron_srv_user.Profile{
		ID:        p.ID.Hex(),
		Avatar:    p.Avatar,
		Mobile:    p.Mobile,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Gender:    p.Gender,
		Birthday:  p.Birthday.String(),
		Nickname:  p.Nickname,
		UserID:    p.UserID.Hex(),
	}
	return &pb
}

func (u *User) ToPB() *kit_iron_srv_user.User {
	// kit_iron_srv_user.User
	upb := kit_iron_srv_user.User{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		Enabled:   u.Enabled,
		SchoolID:  u.SchoolID.Hex(),
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
		// Profile:  u.Profile.ToPB(),
	}
	if !assistant.IsZero(u.Profile) {
		upb.Profile = u.Profile.ToPB()
	}
	return &upb
}

// type ClubProfile struct {
// 	monger.Document `json:",inline" bson:",inline"`

// 	JobID bson.ObjectId `json:"-" bson:"job_id"`
// 	OrganizationID bson.ObjectId `json:"-" bson:"organization_id"`

// 	// 社团职务
// 	Job string `json:"job"`

// }
