package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
)

type SocialApplication struct {
	monger.Schema `json:",inline" bson:",inline"`

	SubjectID  bson.ObjectId `json:"subject_id,omitempty" bson:"subject_id,omitempty"`   // 归属对象ID
	SenderID   bson.ObjectId `json:"sender_id,omitempty" bson:"sender_id,omitempty"`     // 发送者ID
	ReceiverID bson.ObjectId `json:"receiver_id,omitempty" bson:"receiver_id,omitempty"` // 接收者ID
	Kind       string        `json:"kind,omitempty" bson:"kind,omitempty"`               // 请求类型
	Body       string        `json:"body,omitempty" bson:"body,omitempty"`               // 消息体
	State      int32         `json:"state,omitempty" bson:"state,omitempty"`             // 状态
}
