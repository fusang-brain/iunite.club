package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
)

type TemplateDetails map[string]interface{}

type StoreReportTemplate struct {
	UserID       bson.ObjectId         `json:"user_id,omitempty" bson:"user_id,omitempty"`             // 创建者
	Title        string                `json:"title,omitempty" bson:"title,omitempty"`                 // 标题
	Description  string                `json:"description,omitempty" bson:"description,omitempty"`     // 描述
	Body         string                `json:"body,omitempty" bson:"body,omitempty"`                   // 模板内容
	CustomFields []TemplateCustomField `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"` // 自定义字段
	Config       TemplateConfig        `json:"config,omitempty" bson:"config,omitempty"`               // 模板配置
}

type StoreApprovedTemplate struct {
	UserID       bson.ObjectId              `json:"user_id,omitempty" bson:"user_id,omitempty"`         // 创建者
	Title        string                     `json:"title,omitempty" bson:"title,omitempty"`             // 标题
	Description  string                     `json:"description,omitempty" bson:"description,omitempty"` // 描述
	Body         string                     `json:"body,omitempty" bson:"body,omitempty"`
	CustomFields []TemplateCustomField      `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"`
	FlowsConfig  TemplateApprovedFlowConfig `json:"flows_config,omitempty" bson:"flows_config,omitempty"` // 流程配置
}

type TemplateStore struct {
	monger.Schema          `json:",inline" bson:",inline"`
	CreatorID              bson.ObjectId   `json:"creator_id,omitempty" bson:"creator_id,omitempty"`
	Creator                *User           `json:"creator,omitempty" bson:"creator,omitempty" monger:"belongTo,foreignKey=creator_id"`
	Title                  string          `json:"title,omitempty" bson:"title,omitempty"`                                   // 模板标题
	Description            string          `json:"description,omitempty" bson:"description,omitempty"`                       // 模板描述
	Price                  int64           `json:"price,omitempty" bson:"price,omitempty"`                                   // 价格
	TemplateDetails        TemplateDetails `json:"details,omitempty" bson:"details,omitempty"`                               // 模板详情
	AvailableOrganizations []bson.ObjectId `json:"available_organization,omitempty" bson:"available_organization,omitempty"` // 可见组织
}

func (ts *TemplateStore) SetTemplateDetails(details interface{}) {
	templateDetails := TemplateDetails{}
	if srt, ok := details.(*StoreReportTemplate); ok {
		templateDetails["user_id"] = srt.UserID.Hex()
		templateDetails["title"] = srt.Title
		templateDetails["description"] = srt.Description
		templateDetails["body"] = srt.Body
		templateDetails["custom_fields"] = srt.CustomFields
		templateDetails["config"] = srt.Config
	}

	if sat, ok := details.(*StoreApprovedTemplate); ok {
		templateDetails["user_id"] = sat.UserID.Hex()
		templateDetails["title"] = sat.Title
		templateDetails["description"] = sat.Description
		templateDetails["body"] = sat.Body
		templateDetails["custom_fields"] = sat.CustomFields
		templateDetails["flows_config"] = sat.FlowsConfig
	}

	ts.TemplateDetails = templateDetails
}

func (ts *TemplateStore) SetDetailsOfApprovedTemplate(details *StoreApprovedTemplate) {
	ts.SetTemplateDetails(details)
}

func (ts *TemplateStore) SetDetailsOfReportTemplate(details *StoreReportTemplate) {
	ts.SetTemplateDetails(details)
}
