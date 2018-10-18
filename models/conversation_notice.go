package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/services/core/proto/conversation"
)

type ConversationNotice struct {
	monger.Schema `json:",inline" bson:",inline"`

	ConversationID bson.ObjectId   `json:"conversation_id,omitempty" bson:"conversation_id,omitempty"`
	Title          string          `json:"title,omitempty" bson:"title,omitempty"`
	Body           string          `json:"body,omitempty" bson:"body,omitempty"`
	Readers        []bson.ObjectId `json:"readers,omitempty" bson:"readers,omitempty"`
}

func (self *ConversationNotice) ToPB() *iunite_club_srv_core_conversation.NoticePB {
	res := &iunite_club_srv_core_conversation.NoticePB{
		ConversationID: self.ConversationID.Hex(),
		Title:          self.Title,
		Body:           self.Body,
	}

	return res
}
