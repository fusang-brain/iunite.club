package models

import (
	"github.com/iron-kit/monger"
)

// School class 学校类
type School struct {
	monger.Document `json:",inline" bson:",inline"`
	// SchoolID string `gorm:"primay_key"`
	// 学校名
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// 学校标示名
	SlugName string `json:"slug_name,omitempty" bson:"slug_name,omitempty"`

	// 学校编码
	SchoolCode string `json:"school_code,omitempty" bson:"school_code,omitempty"`
	// Logo

	// 描述
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}
