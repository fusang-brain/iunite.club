package dto_approved

import (
	"gopkg.in/mgo.v2/bson"
	"iunite.club/services/navo/dto/template"
)

type TemplateBundle struct {
	Title                  string                                  `json:"Title,omitempty" description:"预定义标题"`
	Description            string                                  `json:"Description,omitempty" description:"预定义描述"`
	AvailableOrganizations []string                                `json:"AvailableOrganizations,omitempty" description:"预定义接收者"`
	FlowsConfig            dto_template.TemplateApprovedFlowConfig `json:"FlowsConfig,omitempty" description:"流程配置"`
	ClubID                 string                                  `json:"ClubID,omitempty" description:"社团ID"`
	CustomFields           []dto_template.TemplateCustomField      `json:"CustomFields,omitempty" description:"自定义字段列表"`
}

type UpdateTemplateBundle struct {
	Title                  string   `json:"title,omitempty" description:"预定义标题"`
	Description            string   `json:"description,omitempty" description:"预定义描述"`
	AvailableOrganizations []string `json:"available_organizations,omitempty" description:"预定义接收者"`
	FlowsConfig            struct {
		Kind  int32 `json:"kind,omitempty" description:"流程类型, 0: custom"`
		Flows []struct {
			UserID bson.ObjectId `json:"user_id,omitempty" description:"用户ID, 审批人ID"`
			Sort   int32         `json:"sort,omitempty" description:"排序, 流程排序"`
		}
	} `json:"flows_config,omitempty" description:"流程配置"`
	ClubID       string `json:"club_id,omitempty" description:"社团ID"`
	CustomFields []struct {
		Key     string                 `json:"key,omitempty" description:"主键, 客户端生成的UUID"`
		Kind    string                 `json:"kind,omitempty" description:"类型: text, checkbox, radio"`
		Label   string                 `json:"label,omitempty" description:"字段标签"`
		Options map[string]interface{} `json:"options,omitempty" description:"若多选时需要设置选项"`
		Sort    int32                  `json:"sort,omitempty" description:"排序"`
	} `json:"custom_fields,omitempty" description:"自定义字段列表"`
}
