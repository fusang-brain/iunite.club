package models

import (
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	recruitmentPB "iunite.club/services/core/proto/recruitment"
)

type RecruitmentAnswer struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	FormID  bson.ObjectId `json:"form_id,omitempty" bson:"form_id,omitempty"`
	ItemKey bson.ObjectId `json:"item_key,omitempty" bson:"item_key,omitempty"`
	Answer  string        `json:"answer,omitempty" bson:"answer,omitempty"`
}

type RecruitmentFormRecord struct {
	monger.Schema `json:",inline" bson:",inline"`

	Mobile          string              `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Name            string              `json:"name,omitempty" bson:"name,omitempty"`
	Major           string              `json:"major,omitempty" bson:"major,omitempty"`
	Age             int32               `json:"age,omitempty" bson:"age,omitempty"`
	SchoolStudentID string              `json:"school_student_id,omitempty" bson:"school_student_id,omitempty"`
	DepartmentID    bson.ObjectId       `json:"department_id,omitempty" bson:"department_id,omitempty"`
	RecordID        bson.ObjectId       `json:"record_id,omitempty" bson:"record_id,omitempty"`
	Status          int32               `json:"status,omitempty" bson:"status,omitempty"` // 0: 报名状态, 1: 已通过 2: 已拒绝
	Answers         []RecruitmentAnswer `json:"answers,omitempty" bson:"answers,omitempty"`
}

func (self *RecruitmentAnswer) ToPB() *recruitmentPB.RecruitmentAnswer {
	result := &recruitmentPB.RecruitmentAnswer{
		ID:      self.ID.Hex(),
		FormID:  self.FormID.Hex(),
		ItemKey: self.ItemKey.Hex(),
		Answer:  self.Answer,
	}

	return result
}

func (self *RecruitmentFormRecord) ToPB() *recruitmentPB.RecruitmentFormRecord {
	result := &recruitmentPB.RecruitmentFormRecord{
		ID:              self.ID.Hex(),
		Mobile:          self.Mobile,
		Name:            self.Name,
		Major:           self.Major,
		Age:             self.Age,
		SchoolStudentID: self.SchoolStudentID,
		DepartmentID:    self.DepartmentID.Hex(),
		RecordID:        self.RecordID.Hex(),
		Status:          self.Status,
		// Answers:
	}

	if len(self.Answers) > 0 {
		answers := make([]*recruitmentPB.RecruitmentAnswer, 0)
		for _, a := range self.Answers {
			answers = append(answers, a.ToPB())
		}

		result.Answers = answers
	}

	return result
}
