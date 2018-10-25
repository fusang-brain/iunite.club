package dto_report

type SimpleReportBundle struct {
	Title       string   `json:"title,omitempty" description:"title"`
	Description string   `json:"description,omitempty" description:"description"`
	Receivers   []string `json:"receivers,omitempty" description:"receivers"`
	Body        string   `json:"body,omitempty" description:"body"`
}
