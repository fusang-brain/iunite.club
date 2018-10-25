package dto_announce

import "time"

type CreateInstructionsBundle struct {
	Name   string `json:"name,omitempty" description:"名称"`
	Body   string `json:"body,omitempty" description:"内容"`
	ClubID string `json:"club_id,omitempty" description:"社团ID"`
}

type CreateTaskBundle struct {
	Name      string    `json:"name,omitempty" description:"名称"`
	Body      string    `json:"body,omitempty" description:"内容"`
	ClubID    string    `json:"club_id,omitempty" description:"社团ID"`
	StartTime time.Time `json:"start_time,omitempty" description:"任务开始时间"`
	EndTime   time.Time `json:"end_time,omitempty" description:"任务结束时间"`
	Users     []string  `json:"users,omitempty" description:"接收人"`
}

type CreateReminderBundle struct {
	Name         string    `json:"name,omitempty" description:"名称"`
	Body         string    `json:"body,omitempty" description:"内容"`
	ClubID       string    `json:"club_id,omitempty" description:"社团ID"`
	ReminderTime time.Time `json:"reminder_time,omitempty" description:"提醒时间"`
	Users        []string  `json:"users,omitempty" description:"接收人"`
}
