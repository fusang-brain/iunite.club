package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/report/proto"
)

type ReportTemplate struct {
	monger.Schema `json:",inline" bson:",inline"`
	UserID        bson.ObjectId         `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ClubID        bson.ObjectId         `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Title         string                `json:"title,omitempty" bson:"title,omitempty"`
	Description   string                `json:"description,omitempty" bson:"description,omitempty"`
	Body          string                `json:"body,omitempty" bson:"body,omitempty"`
	Receivers     []bson.ObjectId       `json:"receivers,omitempty" bson:"receivers,omitempty"`
	CustomFields  []TemplateCustomField `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"`
	Config        TemplateConfig        `json:"config,omitempty" bson:"config,omitempty"`
	Enable        bool                  `json:"enable,omitempty" bson:"enable"`
}

func (rt *ReportTemplate) ToPB() *pb.ReportTemplatePB {
	result := &pb.ReportTemplatePB{
		ClubID:      rt.ClubID.Hex(),
		Title:       rt.Title,
		Description: rt.Description,
		Body:        rt.Body,
	}

	if len(rt.Receivers) > 0 {
		receivers := make([]string, 0, len(rt.Receivers))
		for _, v := range rt.Receivers {
			receivers = append(receivers, v.Hex())
		}

		result.Receivers = receivers
	}

	if len(rt.CustomFields) > 0 {
		fields := make([]*pb.TemplateCustomFieldPB, 0, len(rt.CustomFields))
		for _, v := range rt.CustomFields {
			fields = append(fields, v.ToPB())
		}
		result.CustomFields = fields
	}

	result.Config = rt.Config.ToPB()
	return result
}
