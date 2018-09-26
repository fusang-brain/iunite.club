package dto

type User struct {
	ID        string   `json:"ID,omitempty"`
	CreatedAt int64    `json:"CreatedAt,omitempty"`
	UpdatedAt int64    `json:"UpdatedAt,omitempty"`
	IsTeacher bool     `json:"isTeacher,omitempty"`
	IsAdmin   bool     `json:"ee,omitempty"`
	Username  string   `json:"Username,omitempty"`
	Mobile    string   `json:"Mobile,omitempty"`
	AreaCode  string   `json:"AreaCode,omitempty"`
	Email     string   `json:"Email,omitempty"`
	Enabled   bool     `json:"Enabled,omitempty"`
	School    *School  `json:"School,omitempty"`
	Profile   *Profile `json:"Profile,omitempty"`
}

type Profile struct {
	ID               string `json:"ID,omitempty"`
	CreatedAt        int64  `json:"CreatedAt,omitempty"`
	UpdatedAt        int64  `json:"UpdatedAt,omitempty"`
	UserNO           string `json:"UserNO,omitempty"`
	Avatar           string `json:"Avatar,omitempty"`
	FirstName        string `json:"FirstName,omitempty"`
	LastName         string `json:"LastName,omitempty"`
	Gender           string `json:"Gender,omitempty"`
	Birthday         int64  `json:"Birthday,omitempty"`
	Nickname         string `json:"Nickname,omitempty"`
	Status           string `json:"Status,omitempty"`
	UserID           string `json:"UserID,omitempty"`
	SchoolDepartment string `json:"SchoolDepartment,omitempty"`
	Major            string `json:"Major,omitempty"`
	SchoolClass      string `json:"SchoolClass,omitempty"`
	AdvisorMobile    string `json:"AdvisorMobile,omitempty"`
	AdvisorName      string `json:"AdvisorName,omitempty"`
	StudentID        string `json:"StudentID,omitempty"`
	RoomNumber       string `json:"RoomNumber,omitempty"`
}
