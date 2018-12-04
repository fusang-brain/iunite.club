package dto_report

type SimpleReportBundle struct {
	Title       string   `json:"title,omitempty" description:"标题"`
	Description string   `json:"description,omitempty" description:"描述"`
	Receivers   []string `json:"receivers,omitempty" description:"接收者"`
	Body        string   `json:"body,omitempty" description:"方法体"`
	ClubID      string   `json:"club_id,omitempty" description:"社团ID"`
}
