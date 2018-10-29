package dto_report

import (
	"time"

	"iunite.club/services/navo/dto/template"
)

type TemplateConfig struct {
	Kind      string    `json:"kind,omitempty" description:"类型, custom: 自定义, weeks: 周"` // custom or weeks
	StartTime time.Time `json:"start_time,omitempty" description:"开始时间"`
	EndTime   time.Time `json:"end_time,omitempty" description:"结束时间"`
	Weeks     []int32   `json:"weeks,omitempty" description:"星期列表 [0,1,2,3,4,5,6] 星期日 到 星期六"`
}

// type TemplateCustomField struct {
// 	Key     string                 `json:"key,omitempty" description:"主键, 客户端生成的UUID"`
// 	Kind    string                 `json:"kind,omitempty" description:"类型: text, checkbox, radio"`
// 	Label   string                 `json:"label,omitempty" description:"字段标签"`
// 	Options map[string]interface{} `json:"options,omitempty" description:"若多选时需要设置选项"`
// 	Sort    int32                  `json:"sort,omitempty" description:"排序"`
// }

type TemplateBundle struct {
	Title        string                             `json:"title,omitempty" description:"预定义标题"`
	Description  string                             `json:"description,omitempty" description:"预定义描述"`
	Receivers    []string                           `json:"receivers,omitempty" description:"预定义接收者"`
	Config       TemplateConfig                     `json:"config,omitempty" description:"配置"`
	ClubID       string                             `json:"club_id,omitempty" description:"社团ID"`
	CustomFields []dto_template.TemplateCustomField `json:"custom_fields,omitempty" description:"自定义字段列表"`
}
