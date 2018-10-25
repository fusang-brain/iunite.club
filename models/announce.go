package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	announcePB "iunite.club/services/core/proto/announce"
)

const (
	KindAnnounceInstructions = "instructions"
	KindAnnounceTask         = "task"
	KindAnnounceReminder     = "reminder"
)

type AnnounceReceiver struct {
	UserID  bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	HasRead bool          `json:"has_read,omitempty" bson:"has_read"`
}

type Announce struct {
	monger.Schema `json:",inline" bson:",inline"`

	Name      string                 `json:"name,omitempty" bson:"name,omitempty"`
	Body      string                 `json:"body,omitempty" bson:"body,omitempty"`
	Kind      string                 `json:"kind,omitempty" bson:"kind,omitempty"` // instructions: 指令, task: 任务, reminder: 提醒
	ClubID    bson.ObjectId          `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Receivers []AnnounceReceiver     `json:"receivers,omitempty" bson:"receivers,omitempty"`
	Options   map[string]interface{} `json:"options,omitempty" bson:"options,omitempty"`
}

func (self *Announce) ToPB() *announcePB.AnnouncePB {
	result := &announcePB.AnnouncePB{
		ID:        self.ID.Hex(),
		Name:      self.Name,
		Body:      self.Body,
		Kind:      self.Kind,
		ClubID:    self.ClubID.Hex(),
		Options:   hptypes.EncodeToStruct(self.Options),
		CreatedAt: hptypes.TimestampProto(self.CreatedAt),
		UpdatedAt: hptypes.TimestampProto(self.UpdatedAt),
	}

	if len(self.Receivers) > 0 {
		receivers := make([]*announcePB.AnnounceReceiverPB, 0, len(self.Receivers))
		for _, v := range self.Receivers {
			receivers = append(receivers, &announcePB.AnnounceReceiverPB{
				UserID:  v.UserID.Hex(),
				HasRead: v.HasRead,
			})
		}

		result.Receivers = receivers
	}

	return result
}
