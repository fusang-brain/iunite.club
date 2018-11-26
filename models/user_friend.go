package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
)

type UserFriend struct {
	monger.Schema `json:",inline" bson:",inline"`

	UserID   bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	FriendID bson.ObjectId `json:"friend_id,omitempty" bson:"friend_id,omitempty"`
	User     User          `json:"user,omitempty" bson:"user,omitempty" monger:"belongTo,foreignKey=user_id"`
	Friend   User          `json:"friend,omitempty" bson:"friend,omitempty" monger:"belongTo,foreignKey=friend_id"`
}
