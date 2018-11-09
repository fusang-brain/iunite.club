package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/approved/proto"
)

type ApprovedTemplate struct {
	monger.Schema `json:",inline" bson:",inline"`

	Enabled                bool                       `json:"enabled,omitempty" bson:"enabled"`
	UserID                 bson.ObjectId              `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ClubID                 bson.ObjectId              `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Title                  string                     `json:"title,omitempty" bson:"title,omitempty"`
	Description            string                     `json:"description,omitempty" bson:"description,omitempty"`
	Body                   string                     `json:"body,omitempty" bson:"body,omitempty"`
	AvailableOrganizations []bson.ObjectId            `json:"available_organization,omitempty" bson:"available_organization,omitempty"` // 可见组织
	CustomFields           []TemplateCustomField      `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"`
	FlowsConfig            TemplateApprovedFlowConfig `json:"flows_config,omitempty" bson:"flows_config,omitempty"`
}

func (at *ApprovedTemplate) ToPB() *pb.ApprovedTemplatePB {
	result := &pb.ApprovedTemplatePB{
		ID:          at.ID.Hex(),
		CreatedAt:   hptypes.TimestampProto(at.CreatedAt),
		UpdatedAt:   hptypes.TimestampProto(at.UpdatedAt),
		UserID:      at.UserID.Hex(),
		ClubID:      at.ClubID.Hex(),
		Title:       at.Title,
		Description: at.Description,
		Body:        at.Body,
		FlowsConfig: at.FlowsConfig.ToPB(),
		Enabled:     at.Enabled,
	}

	if len(at.AvailableOrganizations) > 0 {
		orgs := make([]string, 0, len(at.AvailableOrganizations))
		for _, v := range at.AvailableOrganizations {
			orgs = append(orgs, v.Hex())
		}
		result.AvailableOrganizations = orgs
	}

	if len(at.CustomFields) > 0 {
		fields := make([]*pb.TemplateCustomFieldPB, 0, len(at.CustomFields))
		for _, v := range at.CustomFields {
			fields = append(fields, v.ToApprovedPB())
		}

		result.CustomFields = fields
	}

	return result
}

func (at *ApprovedTemplate) SetByPB(pb *pb.ApprovedTemplatePB) {
	at.UserID = bson.ObjectIdHex(pb.UserID)
	at.ClubID = bson.ObjectIdHex(pb.ClubID)
	at.Title = pb.Title
	at.Description = pb.Description
	at.Body = pb.Body
	if len(pb.AvailableOrganizations) > 0 {
		orgs := make([]bson.ObjectId, 0, len(pb.AvailableOrganizations))
		for _, v := range pb.AvailableOrganizations {
			orgs = append(orgs, bson.ObjectIdHex(v))
		}
		at.AvailableOrganizations = orgs
	}
	if len(pb.CustomFields) > 0 {
		fields := make([]TemplateCustomField, 0, len(pb.CustomFields))
		for _, v := range pb.CustomFields {
			field := TemplateCustomField{}
			field.SetByApprovedPB(v)
			fields = append(fields, field)
		}
		at.CustomFields = fields
	}

	if pb.FlowsConfig != nil {
		at.FlowsConfig = TemplateApprovedFlowConfig{}
		at.FlowsConfig.SetByPB(pb.FlowsConfig)
	}
}
