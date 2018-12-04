package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/report/proto"
)

type Report struct {
	monger.Schema `json:",inline" bson:",inline"`

	Kind        string                 `json:"kind,omitempty" bson:"kind,omitempty"`
	ClubID      bson.ObjectId          `json:"club_id,omitempty" bson:"club_id,omitempty"`
	UserID      bson.ObjectId          `json:"user_id,omitempty" bson:"user_id,omitempty"`
	User        *User                  `json:"user,omitempty" bson:"user,omitempty" monger:"belongTo,foreignKey=user_id"`
	Title       string                 `json:"title,omitempty" bson:"title,omitempty"`
	Description string                 `json:"description,omitempty" bson:"description,omitempty"`
	Body        string                 `json:"body,omitempty" bson:"body,omitempty"`
	Receivers   []bson.ObjectId        `json:"receivers,omitempty" bson:"receivers,omitempty"`
	Results     map[string]interface{} `json:"results,omitempty" bson:"results,omitempty"`         // 自定义字段内容
	TemplateID  bson.ObjectId          `json:"template_id,omitempty" bson:"template_id,omitempty"` // 模板ID
}

func (r *Report) ToPB() *pb.ReportPB {
	// fmt.Println(r, "report")
	result := &pb.ReportPB{
		ID:          r.ID.Hex(),
		Title:       r.Title,
		Description: r.Description,
		Body:        r.Body,
		Kind:        r.Kind,
		CreatedAt:   hptypes.TimestampProto(r.CreatedAt),
		UpdatedAt:   hptypes.TimestampProto(r.UpdatedAt),
		Results:     hptypes.EncodeToStruct(r.Results),
		TemplateID:  r.TemplateID.Hex(),
	}

	if len(r.Receivers) > 0 {
		receivers := make([]string, 0)
		for _, v := range r.Receivers {
			receivers = append(receivers, v.Hex())
		}
		result.Receivers = receivers
	}

	if r.User != nil {
		result.User = r.User.ToPB()
	}

	return result
}
