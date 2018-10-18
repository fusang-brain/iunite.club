package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/core/proto/conversation"
)

type ConversationMember struct {
	UserID   bson.ObjectId `json:"user_id,omitempty"  bson:"user_id"`
	Nickname string        `json:"nickname,omitempty"  bson:"nickname"`
	IsTop    bool          `json:"is_top,omitempty"  bson:"is_top"`
}

type Conversation struct {
	monger.Schema `json:",inline" bson:",inline"`

	Kind            string               `json:"kind,omitempty"  bson:"kind,omitempty"`                 // 会话类型 'group': 群组, 'simple': 简单（私聊）, 'temporary': 临时
	Name            string               `json:"name,omitempty"  bson:"name,omitempty"`                 // 会话名称(群组会话时有效)
	Avatar          string               `json:"avatar,omitempty"  bson:"avatar,omitempty"`             // 头像
	Members         []ConversationMember `json:"members,omitempty"  bson:"members,omitempty"`           // 会话涉及用户
	Master          bson.ObjectId        `json:"master,omitempty"  bson:"master,omitempty"`             // 会话管理者
	IsStartValidate bool                 `json:"is_start_validate,omitempty"  bson:"is_start_validate"` // 是否开启验证
	IsTop           bool                 `json:"is_top,omitempty"  bson:"is_top"`                       // 是否置顶
}

func (self *Conversation) ToPB() *pb.ConversationPB {
	res := &pb.ConversationPB{
		ID:              self.ID.Hex(),
		UpdatedAt:       hptypes.TimestampProto(self.UpdatedAt),
		CreatedAt:       hptypes.TimestampProto(self.CreatedAt),
		Kind:            self.Kind,
		Name:            self.Name,
		Avatar:          self.Avatar,
		Master:          self.Master.Hex(),
		IsStartValidate: self.IsStartValidate,
		IsTop:           self.IsTop,
	}

	if len(self.Members) > 0 {
		members := make([]*pb.ConversationMember, 0)
		for _, v := range self.Members {
			members = append(members, &pb.ConversationMember{
				Nickname: v.Nickname,
				UserID:   v.UserID.Hex(),
				IsTop:    v.IsTop,
			})
		}
	}

	return res
}
