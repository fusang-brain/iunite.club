package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Activity struct {
	// monger.Schemer `bson:",inline" json:",inline"`
	ID                bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Subject           string        `json:"subject,omitempty" bson:"subject,omitempty"`
	Location          string        `json:"location,omitempty" bson:"location,omitempty"`
	StartTime         time.Time     `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime           time.Time     `json:"end_time,omitempty" bson:"end_time,omitempty"`
	AmountFunding     int64         `json:"amount_funding,omitempty" bson:"amount_funding,omitempty"`
	ParticipantsTotal int32         `json:"participants_total,omitempty" bson:"participants_total,omitempty"`
	IsPublish         bool          `json:"is_publish,omitempty" bson:"is_publish"`
	CreatedAt         time.Time     `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
}
