package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/core/proto/recruitment"
)

type RecruitmentFormField struct {
	ID      bson.ObjectId          `json:"id,omitempty" bson:"id,omitempty"`
	Subject string                 `json:"subject,omitempty" bson:"subject,omitempty"`
	Kind    string                 `json:"kind,omitempty" bson:"kind,omitempty"`
	Options map[string]interface{} `json:"options,omitempty" bson:"options,omitempty"`
	Key     string                 `json:"key,omitempty" bson:"key,omitempty"`
	Sort    int32                  `json:"sort,omitempty" bson:"sort,omitempty"`
}

type RecruitmentForm struct {
	monger.Schema `json:",inline" bson:",inline"`

	Name     string        `json:"name,omitempty" bson:"name,omitempty"`
	RecordID bson.ObjectId `json:"record_id,omitempty" bson:"record_id,omitempty"`
	// ClubID bson.ObjectId          `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Fields []RecruitmentFormField `json:"fields,omitempty" bson:"fields,omitempty"`
}

func (self *RecruitmentFormField) ToPB() *pb.RecruitmentFormField {
	return &pb.RecruitmentFormField{
		ID:      self.ID.Hex(),
		Subject: self.Subject,
		Kind:    self.Kind,
		Options: hptypes.EncodeToStruct(self.Options),
		Key:     self.Key,
		Sort:    self.Sort,
	}
}

func (self *RecruitmentForm) ToPB() *pb.RecruitmentForm {
	result := &pb.RecruitmentForm{
		ID:       self.ID.Hex(),
		Name:     self.Name,
		RecordID: self.RecordID.Hex(),
	}

	fields := make([]*pb.RecruitmentFormField, 0)
	if len(self.Fields) > 0 {
		for _, f := range self.Fields {
			fields = append(fields, f.ToPB())
		}
	}
	result.Fields = fields
	return result
}
