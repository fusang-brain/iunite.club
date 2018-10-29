package dto_report

type ReportTemplateBundle struct {
	Results map[string]interface{} `json:"results,omitempty"` // 模版汇报表单
	ClubID  string                 `json:"club_id,omitempty"`
}
