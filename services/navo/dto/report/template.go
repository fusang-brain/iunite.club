package dto_report

import (
	"time"
)

type TemplateConfig struct {
	Kind      string    `json:"kind,omitempty"` // custom or weeks
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Weeks     []int32   `json:"weeks,omitempty"` // 星期列表 [0,1,2,3,4,5,6] 星期日 到 星期六
}

type TemplateBundle struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Receivers   []string       `json:"receivers,omitempty"`
	Config      TemplateConfig `json:"config,omitempty"`
}
