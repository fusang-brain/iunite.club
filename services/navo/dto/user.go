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

type OrganizationUser struct {
	ID               string        `json:"id,omitempty"`
	CreatedAt        int64         `json:"created_at,omitempty"`
	UpdatedAt        int64         `json:"updated_at,omitempty"`
	Kind             int           `json:"kind,omitempty"`
	AcceptState      int           `json:"accept_state,omitempty"`
	State            int           `json:"state,omitempty"`
	OrganizationInfo *Organization `json:"organization_info,omitempty"`
	UserInfo         *User         `json:"user_info,omitempty"`
	IsCreator        bool          `json:"is_creator,omitempty"`
	IsSuperManager   bool          `json:"is_super_manager,omitempty"`
	JoinTime         int64         `json:"join_time,omitempty"`
	LeaveTime        int64         `json:"leave_time,omitempty"`
	DepartmentID     string        `json:"department_id,omitempty"`
	JobID            string        `json:"job_id,omitempty"`
}
