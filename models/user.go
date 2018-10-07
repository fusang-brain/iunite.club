package models

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	assistant "github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/utils"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	orgPB "iunite.club/services/organization/proto"
	kit_iron_srv_user "iunite.club/services/user/proto"
)

type SecruityInfo struct {
	AuthType      string `json:"auth_type,omitempty" bson:"auth_type,omitempty"`
	Key           string `json:"key,omitempty" bson:"key,omitempty"`
	Secret        string `json:"secret,omitempty" bson:"secret,omitempty"`
	PlainPassword string `json:"plain_password,omitempty" bson:"plain_password,omitempty"`
}

type User struct {
	monger.Schema `json:",inline" bson:",inline"`

	Username         string            `json:"username,omitempty" bson:"username,omitempty"`
	Enabled          bool              `json:"enabled,omitempty" bson:"enabled"`
	SchoolID         bson.ObjectId     `json:"schoolID,omitempty" bson:"school_id,omitempty"`
	Profile          *Profile          `json:"profile,omitempty" bson:"profile" monger:"hasOne,foreignKey=user_id"`
	SecruityInfos    []SecruityInfo    `json:"-,omitempty" bson:"secruity_infos,omitempty"`
	DefaultClubID    bson.ObjectId     `json:"defaultClubID,omitempty" bson:"defaultClubID"`
	UserClubProfiles []UserClubProfile `json:"user_club_profiles,omitempty" bson:"user_club_profiles" monger:"hasMany,foreignKey=user_id"`
}

type Profile struct {
	monger.Schema `json:",inline" bson:",inline"`

	Avatar           string        `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Mobile           string        `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Email            string        `json:"email,omitempty" bson:"email,omitempty"`
	Firstname        string        `json:"firstname,omitempty" bson:"first_name,omitempty"`
	Lastname         string        `json:"lastname,omitempty" bson:"last_name,omitempty"`
	Gender           string        `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthday         time.Time     `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Nickname         string        `json:"nickname,omitempty" bson:"nickname,omitempty"`
	UserID           bson.ObjectId `json:"userID,omitempty" bson:"user_id,omitempty"`
	User             *User         `json:"user,omitempty" bson:"user" monger:"belongTo,foreignKey=user_id"`
	SchoolDepartment string        `json:"school_department,omitempty" bson:"school_department,omitempty"`
	SchoolClass      string        `json:"school_class,omitempty" bson:"school_class,omitempty"`
	Major            string        `json:"major,omitempty" bson:"major,omitempty"`
	AdvisorMobile    string        `json:"advisor_mobile,omitempty" bson:"advisor_mobile,omitempty"` // 辅导员手机
	AdvisorName      string        `json:"advisor_name,omitempty" bson:"advisor_name,omitempty"`     // 辅导员姓名
	StudentID        string        `json:"student_id,omitempty" bson:"student_id,omitempty"`         // 学号
	RoomNumber       string        `json:"room_number,omitempty" bson:"room_number,omitempty"`       // 寝室号
}

type UserClubProfile struct {
	monger.Schema `json:",inline" bson:",inline"`

	State          int              `json:"state,omitempty" bson:"state,omitempty"` // 0: 申请中 1: 在职 2: 离职 3: 重新申请 4: 拒绝加入
	UserID         bson.ObjectId    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	User           *User            `json:"user,omitempty" bson:"user,omitempty" monger:"belongTo,foreignKey=user_id"`
	OrganizationID bson.ObjectId    `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	Organization   *Organization    `json:"organization,omitempty" bson:"organization,omitempty" monger:"belongTo,foreignKey=organization_id"`
	IsCreator      bool             `json:"is_creator,omitempty" bson:"is_creator"`
	IsMaster       bool             `json:"is_master,omitempty" bson:"is_master"`
	JoinTime       time.Time        `json:"join_time,omitempty" bson:"join_time,omitempty"`
	LeaveTime      time.Time        `json:"leave_time,omitempty" bson:"leave_time,omitempty"`
	JobID          bson.ObjectId    `json:"job_id,omitempty" bson:"job_id,omitempty"`
	Job            *OrganizationJob `json:"job,omitempty" bson:"job,omitempty" monger:"belongTo,foreignKey=job_id"`
	DepartmentID   bson.ObjectId    `json:"department_id,omitempty" bson:"department_id,omitempty"`
	Department     *Organization    `json:"department,omitempty" bson:"department,omitempty" monger:"belongTo,foreignKey=department_id"`
}

func (p *Profile) ToPB() *kit_iron_srv_user.Profile {
	pb := kit_iron_srv_user.Profile{
		ID:        p.ID.Hex(),
		Avatar:    p.Avatar,
		Mobile:    p.Mobile,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Gender:    p.Gender,
		Nickname:  p.Nickname,
		UserID:    p.UserID.Hex(),
	}
	if birthdayPB, err := ptypes.TimestampProto(p.Birthday); err == nil {
		pb.Birthday = birthdayPB
	}
	if t, err := ptypes.TimestampProto(p.CreatedAt); err == nil {
		pb.CreatedAt = t
	}
	if t, err := ptypes.TimestampProto(p.UpdatedAt); err == nil {
		pb.UpdatedAt = t
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
		DefaultClubID: u.DefaultClubID.Hex(),
	}
	if !assistant.IsZero(u.Profile) {
		upb.Profile = u.Profile.ToPB()
	}
	return &upb
}

func (ucp *UserClubProfile) ToPB() *orgPB.UserClubProfile {
	pb := orgPB.UserClubProfile{
		ID:             ucp.ID.Hex(),
		UserID:         ucp.UserID.Hex(),
		OrganizationID: ucp.OrganizationID.Hex(),
		State:          int64(ucp.State),
		IsCreator:      ucp.IsCreator,
		IsMaster:       ucp.IsMaster,
		JoinTime:       ucp.JoinTime.String(),
		LeaveTime:      ucp.LeaveTime.String(),
		JobID:          ucp.JobID.Hex(),
		DepartmentID:   ucp.DepartmentID.Hex(),
	}
	if !utils.IsZero(ucp.User) {
		pb.User = ucp.User.ToPB()
	}
	if !utils.IsZero(ucp.Department) {
		pb.Department = ucp.Department.ToPB()
	}

	if !utils.IsZero(ucp.Organization) {
		pb.Organization = ucp.Organization.ToPB()
	}
	if !utils.IsZero(ucp.Job) {
		pb.Job = ucp.Job.ToPB()
	}

	return &pb
}
