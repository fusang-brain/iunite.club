package models

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	approvedPB "iunite.club/services/core/proto/approved"
)

type ApprovedFlow struct {
	monger.Schema `json:",inline" bson:",inline"`

	Kind       string        `json:"kind,omitempty" bson:"kind,omitempty"`             // approved copy
	Options    string        `json:"options,omitempty" bson:"options,omitempty"`       // 处理意见
	HandlerID  bson.ObjectId `json:"handler_id,omitempty" bson:"handler_id,omitempty"` // 审批人
	Handler    *User         `json:"handler,omitempty" bson:"handler,omitempty" monger:"belongTo,foreignKey=handler_id"`
	Status     int           `json:"status,omitempty" bson:"status"` // 处理状态 0 等待中(未读), 1 处理中(已读), 2 已通过, 3 已拒绝
	Sort       int           `json:"sort,omitempty" bson:"sort"`     // 处理顺序
	ApprovedID bson.ObjectId `json:"approved_id,omitempty" bson:"approved_id,omitempty"`
}

// Approved 审批表
type Approved struct {
	monger.Schema `json:",inline" bson:",inline"`

	DepartmentID bson.ObjectId          `json:"department_id,omitempty" bson:"department_id,omitempty"`
	PusherID     bson.ObjectId          `json:"pusher_id,omitempty" bson:"pusher_id,omitempty"`
	ClubID       bson.ObjectId          `json:"club_id,omitempty" bson:"club_id,omitempty"`
	Title        string                 `json:"title,omitempty" bson:"title,omitempty"`                                         // 审批主题
	Kind         string                 `json:"kind,omitempty" bson:"kind,omitempty"`                                           // 审批类型 ['activity': 活动, 'funding': 经费, 'borrow': 物品]
	Summary      string                 `json:"summary,omitempty" bson:"summary,omitempty"`                                     // 摘要
	Status       string                 `json:"status,omitempty" bson:"status,omitempty"`                                       // 审批状态 [pending: 待处理; doing: 处理中; pass: 通过; refuse: 拒绝]
	Description  string                 `json:"description,omitempty" bson:"description,omitempty"`                             // 审批描述
	Content      map[string]interface{} `json:"content,omitempty" bson:"content,omitempty"`                                     // 审批内容
	Flows        []ApprovedFlow         `json:"flows,omitempty" bson:"flows,omitempty" monger:"hasMany,foreignKey=approved_id"` // 审批流程
}

func (af *ApprovedFlow) ToPB() *approvedPB.ApprovedFlowPB {
	pb := new(approvedPB.ApprovedFlowPB)
	pb.ID = af.ID.Hex()
	pb.CreatedAt = hptypes.TimestampProto(af.CreatedAt)
	pb.UpdatedAt = hptypes.TimestampProto(af.UpdatedAt)
	pb.Kind = af.Kind
	pb.HandlerID = af.HandlerID.Hex()
	pb.Options = af.Options
	pb.Status = int32(af.Status)
	pb.Sort = int32(af.Sort)
	pb.ApprovedID = af.ApprovedID.Hex()

	return pb
}

func (a *Approved) ToPB() *approvedPB.ApprovedPB {
	pb := new(approvedPB.ApprovedPB)

	pb.ID = a.ID.Hex()
	pb.CreatedAt = hptypes.TimestampProto(a.CreatedAt)
	pb.UpdatedAt = hptypes.TimestampProto(a.UpdatedAt)
	pb.Title = a.Title
	pb.Kind = a.Kind
	pb.Summary = a.Summary
	pb.Status = a.Status
	pb.Description = a.Description
	// pb.Content =
	pb.Content = hptypes.EncodeToStruct(a.Content)
	flowSize := len(a.Flows)

	if flowSize > 0 {
		// pb.Flows =
		flows := make([]*approvedPB.ApprovedFlowPB, 0, flowSize)
		for _, v := range a.Flows {
			flows = append(flows, v.ToPB())
		}
		pb.Flows = flows
	}
	return pb
}
