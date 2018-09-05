package models

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/srv/user-srv/proto/user"
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
