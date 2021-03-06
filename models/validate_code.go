package models

import (
	"github.com/iron-kit/monger"
)

// ValidateCode 验证码 Model
type ValidateCode struct {
	monger.Schema `json:",inline" bson:",inline"`

	Mobile    string `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Code      string `json:"code,omitempty" bson:"code,omitempty"`
	Usaged    bool   `json:"usaged,omitempty" bson:"usaged"`
	ExpiredAt int64  `json:"expiredAt,omitempty" bson:"expiredAt,omitempty"`
}
