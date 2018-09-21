package dto

type User struct {
	ID        string  `json:"ID,omitempty"`
	CreatedAt int64   `json:"CreatedAt,omitempty"`
	UpdatedAt int64   `json:"UpdatedAt,omitempty"`
	IsTeacher bool    `json:"isTeacher,omitempty"`
	IsAdmin   bool    `json:"ee,omitempty"`
	Username  string  `json:"Username,omitempty"`
	Mobile    string  `json:"Mobile,omitempty"`
	AreaCode  string  `json:"AreaCode,omitempty"`
	Email     string  `json:"Email,omitempty"`
	Enabled   bool    `json:"Enabled,omitempty"`
	School    *School `json:"School,omitempty"`
}

type School struct {
	ID          string `json:"ID,omitempty"`
	CreatedAt   int64  `json:"CreatedAt,omitempty"`
	UpdatedAt   int64  `json:"UpdatedAt,omitempty"`
	Name        string `json:"Name,omitempty"`
	SlugName    string `json:"SlugName,omitempty"`
	SchoolCode  string `json:"SchoolCode,omitempty"`
	Description string `json:"Description,omitempty"`
}

type Organization struct {
	ID          string  `json:"ID,omitempty"`
	CreatedAt   int64   `json:"CreatedAt,omitempty"`
	UpdatedAt   int64   `json:"UpdatedAt,omitempty"`
	Name        string  `json:"Name,omitempty"`
	SlugName    string  `json:"SlugName,omitempty"`
	Logo        string  `json:"Logo,omitempty"`
	Scale       uint    `json:"Scale,omitempty"`
	SchoolRefer string  `json:"SchoolRefer,omitempty"`
	School      *School `json:"School,omitempty"`
	Description string  `json:"Description,omitempty"`
}
