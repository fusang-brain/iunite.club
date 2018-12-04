package models

import (
	"fmt"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/report/proto"
)

type ReportTemplate struct {
	monger.Schema `json:",inline" bson:",inline"`
	UserID        bson.ObjectId   `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Creator       *User           `json:"creator,omitempty" bson:"creator,omitempty" monger:"belongTo,foreignKey=user_id"`
	ClubID        bson.ObjectId   `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Title         string          `json:"title,omitempty" bson:"title,omitempty"`
	Description   string          `json:"description,omitempty" bson:"description,omitempty"`
	Body          string          `json:"body,omitempty" bson:"body,omitempty"`
	Receivers     []bson.ObjectId `json:"receivers,omitempty" bson:"receivers,omitempty"`
	// ReceiversInfo []User                `json:"ReceiversInfo,omitempty"`
	CustomFields  []TemplateCustomField `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"`
	Config        TemplateConfig        `json:"config,omitempty" bson:"config,omitempty"`
	Enable        bool                  `json:"enable,omitempty" bson:"enable"`
	receiversInfo []User
}

func (rt *ReportTemplate) SetReceiverInfo(users []User) {
	rt.receiversInfo = users
}

func (rt *ReportTemplate) ToPB() *pb.ReportTemplatePB {
	result := &pb.ReportTemplatePB{
		ID:          rt.ID.Hex(),
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

	if len(rt.receiversInfo) > 0 {
		infos := make([]*pb.TemplateUser, 0, len(rt.receiversInfo))

		userToTemplateUser := func(u *User) *pb.TemplateUser {
			tu := new(pb.TemplateUser)
			if u != nil {
				tu.ID = u.ID.Hex()
				if u.Profile != nil {
					tu.Name = u.Profile.Nickname
					tu.Nickname = u.Profile.Nickname
					tu.Avatar = u.Profile.Avatar
					tu.RealName = u.Profile.Firstname + u.Profile.Lastname
				}
			}

			return tu
		}

		for _, v := range rt.receiversInfo {
			infos = append(infos, userToTemplateUser(&v))
		}

		result.ReceiverInfo = infos
	}

	if rt.Creator != nil {
		fmt.Println("temp creator profile", rt.Creator.Profile)
		result.Creator = new(pb.TemplateCreator)
		result.Creator.ID = rt.Creator.ID.Hex()
		if rt.Creator.Profile != nil {
			result.Creator.Avatar = rt.Creator.Profile.Avatar
			result.Creator.Name = rt.Creator.Profile.Nickname
			result.Creator.Nickname = rt.Creator.Profile.Nickname
			result.Creator.RealName = rt.Creator.Profile.Firstname + rt.Creator.Profile.Lastname
		}
	}

	result.Config = rt.Config.ToPB()
	return result
}
