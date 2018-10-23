package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/core/proto/recruitment"
)

// RecruitmentRecord 招新记录
type RecruitmentRecord struct {
	monger.Schema `json:",inline" bson:",inline"`

	ClubID       bson.ObjectId    `json:"club_id,omitempty" bson:"club_id,omitempty"`
	CreateUserID bson.ObjectId    `json:"create_user_id,omitempty" bson:"create_user_id,omitempty"`
	Form         *RecruitmentForm `json:"form,omitempty" bson:"form,omitempty" monger:"belongTo,foreignKey=record_id`
	HasStart     bool             `json:"has_start,omitempty" bson:"has_start"`
	HasEnd       bool             `json:"has_end,omitempty" bson:"has_end"`
}

func (self *RecruitmentRecord) ToPB() *pb.RecruitmentRecord {
	result := &pb.RecruitmentRecord{
		ID:           self.ID.Hex(),
		CreatedAt:    hptypes.TimestampProto(self.CreatedAt),
		UpdatedAt:    hptypes.TimestampProto(self.UpdatedAt),
		ClubID:       self.ClubID.Hex(),
		CreateUserID: self.CreateUserID.Hex(),
		HasStart:     self.HasStart,
		HasEnd:       self.HasEnd,
	}

	if self.Form != nil {
		result.Form = self.Form.ToPB()
	}

	return result
}
