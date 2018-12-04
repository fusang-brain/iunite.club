package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
)

type Permission struct {
	monger.Schema `json:",inline" bson:",inline"`

	PermissionName string `json:"permission_name" bson:"permission_name"`
	PermissionKey  string `json:"permission_key" bson:"permission_key"`
	PathRule       string `json:"path_rule" bson:"path_rule"`
	Method         string `json:"method" bson:"method"`
}

type CasbinRule struct {
	monger.Schema `json:",inline" bson:",inline"`

	PermissionID bson.ObjectId `json:"permission_id" bson:"permission_id"`
	Ptype        string        `json:"ptype" bson:"ptype"`
	V0           string        `json:"v0" bson:"v0"`
	V1           string        `json:"v1" bson:"v1"`
	V2           string        `json:"v2" bson:"v2"`
	V3           string        `json:"v3" bson:"v3"`
	V4           string        `json:"v4" bson:"v4"`
	V5           string        `json:"v5" bson:"v5"`
}
