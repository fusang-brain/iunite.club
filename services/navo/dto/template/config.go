package dto_template

import (
	"gopkg.in/mgo.v2/bson"
)

type TemplateCustomField struct {
	Key     string                 `json:"Key,omitempty" description:"主键, 客户端生成的UUID"`
	Kind    string                 `json:"Kind,omitempty" description:"类型: text, checkbox, radio"`
	Label   string                 `json:"Label,omitempty" description:"字段标签"`
	Options map[string]interface{} `json:"Options,omitempty" description:"若多选时需要设置选项"`
	Sort    int32                  `json:"Sort,omitempty" description:"排序"`
}

type TemplateApprovedFlow struct {
	UserID bson.ObjectId `json:"UserID,omitempty" description:"用户ID, 审批人ID"`
	Sort   int32         `json:"Sort,omitempty" description:"排序, 流程排序"`
}

// TemplateApprovedFlowConfig 审批模板的自定义流程配置
type TemplateApprovedFlowConfig struct {
	Kind  int32                  `json:"Kind,omitempty" description:"流程类型, 0: custom"`
	Flows []TemplateApprovedFlow `json:"Flows,omitempty" description:"流程配置"`
}
