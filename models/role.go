package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/core/proto/role"
)

type RoleGroup struct {
	monger.Schema `json:",inline" bson:",inline"`

	GroupName        string        `json:"group_name,omitempty" bson:"group_description,omitempty"`
	GroupDescription string        `json:"group_description,omitempty" bson:"group_description,omitempty"`
	Organization     bson.ObjectId `json:"organization,omitempty" bson:"organization,omitempty"`
	Roles            []Role        `json:"roles,omitempty" bson:"roles,omitempty" monger:"hasMany,foreignKey=group_id"`
}

type Role struct {
	monger.Schema `json:",inline" bson:",inline"`

	Name         string        `json:"name,omitempty" bson:"name,omitempty"`
	Level        string        `json:"level,omitempty" bson:"level,omitempty"`
	GroupID      bson.ObjectId `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Organization bson.ObjectId `json:"organization,omitempty" bson:"organization,omitempty"`
	Group        *RoleGroup    `json:"group,omitempty" bson:"group,omitempty" monger:"belongTo,foreignKey=group_id"`
}

type RoleUser struct {
	monger.Schema `json:",inline" bson:",inline"`

	RoleID bson.ObjectId `json:"role_id,omitempty" bson:"role_id,omitempty"`
	UserID bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	User   *User         `json:"user,omitempty" bson:"user,omitempty" monger:"belongTo,foreignKey=user_id"`
}

func (rg *RoleGroup) ToPB() *pb.RoleGroupPB {
	res := &pb.RoleGroupPB{
		GroupName:        rg.GroupName,
		GroupDescription: rg.GroupDescription,
		Organization:     rg.Organization.Hex(),
		ID:               rg.ID.Hex(),
	}

	return res
}

func (r *Role) ToPB() *pb.RolePB {
	res := &pb.RolePB{
		ID:           r.ID.Hex(),
		Name:         r.Name,
		Level:        r.Level,
		GroupID:      r.GroupID.Hex(),
		Organization: r.Organization.Hex(),
	}

	if r.Group != nil {
		res.Group = r.Group.ToPB()
	}

	return res
}
