package models

import (
	approvedPB "iunite.club/services/approved/proto"
	"time"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"gopkg.in/mgo.v2/bson"
	pb "iunite.club/services/report/proto"
)

type TemplateApprovedFlowConfig_KindEnum int32

const (
	TemplateApprovedFlowConfig_Custom TemplateApprovedFlowConfig_KindEnum = 0
)

// TemplateConfig 汇报模板配置
type TemplateConfig struct {
	Kind      string    `json:"kind,omitempty" bson:"kind,omitempty"` // custom or weeks
	StartTime time.Time `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty" bson:"end_time,omitempty"`
	Weeks     []int32   `json:"weeks,omitempty" bson:"weeks,omitempty"` // 星期列表 [0,1,2,3,4,5,6] 星期日 到 星期六
}

// TemplateCustomField 模板自定义字段的配置
type TemplateCustomField struct {
	Key     string                 `json:"key,omitempty" bson:"key,omitempty"`         // 字段配置主键, 由前端生成
	Kind    string                 `json:"kind,omitempty" bson:"kind,omitempty"`       // 字段类型 （text, checkbox, radio）
	Label   string                 `json:"label,omitempty" bson:"label,omitempty"`     // 字段名
	Options map[string]interface{} `json:"options,omitempty" bson:"options,omitempty"` // 字段选项 (checkbox or radio时有效)
	Sort    int32                  `json:"sort,omitempty" bson:"sort,omitempty"`       // 字段排序
}

type TemplateApprovedFlow struct {
	UserID bson.ObjectId `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Sort   int32         `json:"sort,omitempty" bson:"sort,omitempty"`
}

// TemplateApprovedFlowConfig 审批模板的自定义流程配置
type TemplateApprovedFlowConfig struct {
	Kind  TemplateApprovedFlowConfig_KindEnum `json:"kind,omitempty" bson:"kind,omitempty"`
	Flows []TemplateApprovedFlow              `json:"flows,omitempty" bson:"flows,omitempty"`
}

func (tc *TemplateConfig) ToPB() *pb.TemplateConfigPB {
	return &pb.TemplateConfigPB{
		Kind:      tc.Kind,
		StartTime: hptypes.TimestampProto(tc.StartTime),
		EndTime:   hptypes.TimestampProto(tc.EndTime),
		Weeks:     tc.Weeks,
	}
}

func (tcf *TemplateCustomField) ToPB() *pb.TemplateCustomFieldPB {
	return &pb.TemplateCustomFieldPB{
		Key:     tcf.Key,
		Kind:    tcf.Kind,
		Label:   tcf.Label,
		Options: hptypes.EncodeToStruct(tcf.Options),
		Sort:    tcf.Sort,
	}
}

func (tcf *TemplateCustomField) SetByPB(pb *pb.TemplateCustomFieldPB) {
	tcf.Key = pb.Key
	tcf.Kind = pb.Kind
	tcf.Label = pb.Label
	tcf.Options = hptypes.DecodeToMap(pb.Options)
	tcf.Sort = pb.Sort
}

func (tcf *TemplateCustomField) SetByApprovedPB(pb *approvedPB.TemplateCustomFieldPB) {
	tcf.Key = pb.Key
	tcf.Kind = pb.Kind
	tcf.Label = pb.Label
	tcf.Options = hptypes.DecodeToMap(pb.Options)
	tcf.Sort = pb.Sort
}

func (tcf *TemplateCustomField) ToApprovedPB() *approvedPB.TemplateCustomFieldPB {
	return &approvedPB.TemplateCustomFieldPB{
		Key:     tcf.Key,
		Kind:    tcf.Kind,
		Label:   tcf.Label,
		Options: hptypes.EncodeToStruct(tcf.Options),
		Sort:    tcf.Sort,
	}
}

func (taf *TemplateApprovedFlow) ToPB() *approvedPB.TemplateApprovedFlowPB {
	return &approvedPB.TemplateApprovedFlowPB{
		UserID: taf.UserID.Hex(),
		Sort:   taf.Sort,
	}
}

func (tafc *TemplateApprovedFlowConfig) ToPB() *approvedPB.TemplateApprovedFlowConfigPB {
	result := &approvedPB.TemplateApprovedFlowConfigPB{
		Kind: int32(tafc.Kind),
	}

	if len(tafc.Flows) > 0 {
		flows := make([]*approvedPB.TemplateApprovedFlowPB, 0, len(tafc.Flows))
		for _, v := range tafc.Flows {
			flows = append(flows, v.ToPB())
		}

		result.Flows = flows
	}

	return result
}

func (tafc *TemplateApprovedFlowConfig) SetByPB(pb *approvedPB.TemplateApprovedFlowConfigPB) {
	tafc.Kind = TemplateApprovedFlowConfig_KindEnum(pb.Kind)
	if len(pb.Flows) > 0 {
		flows := make([]TemplateApprovedFlow, 0, len(pb.Flows))
		for _, v := range pb.Flows {
			flows = append(flows, TemplateApprovedFlow{
				UserID: bson.ObjectIdHex(v.UserID),
				Sort:   v.Sort,
			})
		}
		tafc.Flows = flows
	}
}
