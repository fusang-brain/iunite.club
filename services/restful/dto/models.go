package dto

type Activity struct {
	ID                string        `json:"ID,omitempty"`
	CreatedAt         int64         `json:"CreatedAt,omitempty"`
	UpdatedAt         int64         `json:"UpdatedAt,omitempty"`
	Subject           string        `json:"Subject,omitempty"`     // 主题
	Location          string        `json:"Location,omitempty"`    // 地点
	ApplicantID       string        `json:"ApplicantID,omitempty"` // 申请人
	Applicant         *User         `json:"Applicant,omitempty"`   // 用户
	ApplyDeptID       string        `json:"ApplyDeptID,omitempty"` // 申请部门
	StartTime         int64         `json:"StartTime,omitempty"`
	EndTime           int64         `json:"EndTime,omitempty"`
	AmountFunding     int64         `json:"AmountFunding,omitempty"` // 活动经费-费用总计(经费)
	ParticipantsTotal int           `json:"ParticipantsTotal,omitempty"`
	ApprovedID        string        `json:"ApprovedID,omitempty"` // 审批关联
	IsUndo            bool          `json:"IsUndo,omitempty"`     // 是否已经撤销
	IsPublish         bool          `json:"IsPublish,omitempty"`
	Approved          *ApprovedTask `json:"Approved,omitempty"`
	Pictures          []*File       `json:"Pictures,omitempty"`
	Attach            []*File       `json:"Attach,omitempty"`
}

type ApprovedTask struct {
	ID                string            `json:"ID,omitempty"`
	CreatedAt         int64             `json:"CreatedAt,omitempty"`
	UpdatedAt         int64             `json:"UpdatedAt,omitempty"`
	Name              string            `json:"Name,omitempty"`         // 审批任务名称
	Description       string            `json:"Description,omitempty"`  // 审批描述
	ApprovedType      string            `json:"ApprovedType,omitempty"` // 审批类型 'activity' 'funding' 'borrow'
	Content           string            `json:"Content,omitempty"`      // 审批内容
	Status            string            `json:"Status,omitempty"`       // 审批状态 pending: 待处理; doing: 处理中; pass: 通过; refuse: 拒绝;
	ApprovedProcesses []ApprovedProcess `json:"ApprovedProcesses,omitempty"`
	Summary           string            `json:"Summary,omitempty"`
}

// ApprovedProcess 审批人员表
type ApprovedProcess struct {
	ID          string        `json:"ID,omitempty"`
	CreatedAt   int64         `json:"CreatedAt,omitempty"`
	UpdatedAt   int64         `json:"UpdatedAt,omitempty"`
	ProcessType string        `json:"ProcessType,omitempty"` // 过程类型 （approved[审批], copy[抄送]）
	Options     string        `json:"Options,omitempty"`     // 审批意见
	ApprovedID  string        `json:"ApprovedID,omitempty"`  // 审批
	Approved    *ApprovedTask `json:"Approved,omitempty"`
	HandlerID   string        `json:"HandlerID,omitempty"`
	Handler     *User         `json:"Handler,omitempty"`
	Archived    bool          `json:"Archived,omitempty"`
	Status      int           `json:"Status,omitempty"` // 0 等待中(未读), 1 处理中(已读), 2 已通过, 3 已拒绝
	ProcessSort int           `json:"ProcessSort,omitempty"`
}

type Announce struct {
	ID           string `json:"ID,omitempty"`
	CreatedAt    int64  `json:"CreatedAt,omitempty"`
	UpdatedAt    int64  `json:"UpdatedAt,omitempty"`
	Name         string `json:"Name,omitempty"`
	Body         string `json:"Body,omitempty"`
	Kind         string `json:"Kind,omitempty"` // instructions: 指令, task: 任务, reminder: 提醒'
	Org          string `json:"Org,omitempty"`
	Receivers    []User `json:"Receivers,omitempty"`
	ReminderTime int64  `json:"ReminderTime,omitempty"`
}

type AnnounceReceiver struct {
	ID         string `json:"ID,omitempty"`
	CreatedAt  int64  `json:"CreatedAt,omitempty"`
	UpdatedAt  int64  `json:"UpdatedAt,omitempty"`
	AnnounceID string `json:"AnnounceID,omitempty"`
	UserID     string `json:"UserID,omitempty"`
	HasRead    bool   `json:"HasRead,omitempty"`
}

// CloudDisk 云盘表
type CloudDisk struct {
	ID           string       `json:"ID,omitempty"`
	CreatedAt    int64        `json:"CreatedAt,omitempty"`
	UpdatedAt    int64        `json:"UpdatedAt,omitempty"`
	Organization string       `json:"Organization,omitempty"`
	Name         string       `json:"Name,omitempty"`
	OriginalName string       `json:"OriginalName,omitempty"`
	ParentID     string       `json:"ParentID,omitempty"`
	Kind         int          `json:"Kind,omitempty"`    // 0: 目录 1: 文件
	OwnerID      string       `json:"OwnerID,omitempty"` // 创建人(拥有者)
	Owner        *User        `json:"Owner,omitempty"`
	FileID       string       `json:"FileID,omitempty"`
	File         *File        `json:"File,omitempty"`
	EnabledToAll bool         `json:"EnabledToAll,omitempty"`
	Users        []User       `json:"Users,omitempty"`
	Departments  []Department `json:"Departments,omitempty"`
}

type Department struct {
	ID                string             `json:"ID,omitempty"`
	CreatedAt         int64              `json:"CreatedAt,omitempty"`
	UpdatedAt         int64              `json:"UpdatedAt,omitempty"`
	Name              string             `json:"Name,omitempty"`              // 部门架构名
	SlugName          string             `json:"SlugName,omitempty"`          // 部门架构别名 （拼音）
	Org               string             `json:"Org,omitempty"`               // 组织标示
	ParentID          string             `json:"ParentID,omitempty"`          // 父级部门
	TreePath          string             `json:"TreePath,omitempty"`          // 组织架构路径 （dept1_dept2_dept3）
	Description       string             `json:"Description,omitempty"`       // 部门描述
	OrganizationUsers []OrganizationUser `json:"OrganizationUsers,omitempty"` // 该部门的用户列表, 一个部门可以包涵多个用户
}

type File struct {
	ID               string `json:"ID,omitempty"`
	CreatedAt        int64  `json:"CreatedAt,omitempty"`
	UpdatedAt        int64  `json:"UpdatedAt,omitempty"`
	Filename         string `json:"Filename,omitempty"`
	Path             string `json:"Path,omitempty"`
	Ext              string `json:"Ext,omitempty"`
	AbstractPath     string `json:"AbstractPath,omitempty"`
	Host             string `json:"Host,omitempty"`
	OriginalFilename string `json:"OriginalFilename,omitempty"`
	Size             int64  `json:"Size,omitempty"`
}

type FriendAccept struct {
	ID            string `json:"ID,omitempty"`
	CreatedAt     int64  `json:"CreatedAt,omitempty"`
	UpdatedAt     int64  `json:"UpdatedAt,omitempty"`
	SenderRefer   string `json:"SenderRefer,omitempty"`
	Sender        *User  `json:"Sender,omitempty"`
	ReceiverRefer string `json:"ReceiverRefer,omitempty"`
	Receiver      *User  `json:"Receiver,omitempty"`
	Kind          int    `json:"Kind,omitempty"` // 申请类型: 0: 好友申请 1: 加入申请
	GroupID       string `json:"GroupID,omitempty"`
	Body          string `json:"Body,omitempty"`
	State         int    `json:"State,omitempty"` // 好友申请的状态 0: 待处理 1: 已经通过 2: 已经拒绝
}

// Funding Object
type Funding struct {
	ID             string        `json:"ID,omitempty"`
	CreatedAt      int64         `json:"CreatedAt,omitempty"`
	UpdatedAt      int64         `json:"UpdatedAt,omitempty"`
	ApplicantID    string        `json:"ApplicantID,omitempty"` // 申请人
	Applicant      *User         `json:"Applicant,omitempty"`
	ApplyDeptID    string        `json:"ApplyDeptID,omitempty"` // 申请部门
	ApplyDept      *Department   `json:"ApplyDept,omitempty"`
	ApplyPurpose   string        `json:"ApplyPurpose,omitempty"`   // 申请目的
	AmountApplyFee int64         `json:"AmountApplyFee,omitempty"` // 申请费用
	ApprovedID     string        `json:"ApprovedID,omitempty"`     // 审批关联
	Approved       *ApprovedTask `json:"Approved,omitempty"`
	IsUndo         bool          `json:"IsUndo,omitempty"` // 是否已经撤销
	Pictures       []*File       `json:"Pictures,omitempty"`
	Attach         []*File       `json:"Attach,omitempty"`
}

// GoodsBorrow Object 活动
type GoodsBorrow struct {
	ID          string             `json:"ID,omitempty"`
	CreatedAt   int64              `json:"CreatedAt,omitempty"`
	UpdatedAt   int64              `json:"UpdatedAt,omitempty"`
	Subject     string             `json:"Subject,omitempty"`     // 借用主题
	Description string             `json:"Description,omitempty"` // 借用说明
	Goods       []*GoodsBorrowItem `json:"Goods,omitempty"`       // 物品详情
	ApplicantID string             `json:"ApplicantID,omitempty"` // 申请人
	Applicant   *User              `json:"Applicant,omitempty"`
	ApplyDeptID string             `json:"ApplyDeptID,omitempty"` // 申请部门
	ApplyDept   *Department        `json:"ApplyDept,omitempty"`
	StartTime   int64              `json:"StartTime,omitempty"` // 借用时间
	EndTime     int64              `json:"EndTime,omitempty"`
	ApprovedID  string             `json:"ApprovedID,omitempty"` // 审批关联
	Approved    *ApprovedTask      `json:"Approved,omitempty"`
	IsUndo      bool               `json:"IsUndo,omitempty"` // 是否已经撤销
	Pictures    []*File            `json:"Pictures,omitempty"`
	Attach      []*File            `json:"Attach,omitempty"`
}

// GoodsBorrowItem 借用物品项目Model
type GoodsBorrowItem struct {
	ID            string `json:"ID,omitempty"`
	CreatedAt     int64  `json:"CreatedAt,omitempty"`
	UpdatedAt     int64  `json:"UpdatedAt,omitempty"`
	Name          string `json:"Name,omitempty"`
	Count         int    `json:"Count,omitempty"`
	GoodsBorrowID string `json:"GoodsBorrowID,omitempty"`
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
	Scale       int     `json:"Scale,omitempty"`
	SchoolRefer string  `json:"SchoolRefer,omitempty"`
	School      *School `json:"School,omitempty"`
	Description string  `json:"Description,omitempty"`
}

type OrganizationUser struct {
	ID               string        `json:"ID,omitempty"`
	CreatedAt        int64         `json:"CreatedAt,omitempty"`
	UpdatedAt        int64         `json:"UpdatedAt,omitempty"`
	Kind             int           `json:"Kind"`
	AcceptState      int           `json:"AcceptState"`
	State            int           `json:"State"` // 用户在社团的状态 (0: 未知, 1: 在职, 2: 离职)
	OrganizationInfo *Organization `json:"OrganizationInfo,omitempty"`
	UserInfo         *User         `json:"UserInfo,omitempty"`
	IsCreator        bool          `json:"IsCreator"`
	IsSuperManager   bool          `json:"IsSuperManager"`
	JoinTime         int64         `json:"JoinTime,omitempty"`
	LeaveTime        int64         `json:"LeaveTime,omitempty"`
	DepartmentID     string        `json:"DepartmentID,omitempty"`
	Department       *Department   `json:"Department,omitempty"`
	JobID            string        `json:"JobID,omitempty"`
	Job              *Job          `json:"UserJob,omitempty"`
}

type OrganizationPaperwork struct {
	OrganizationUserID string `json:"organization_user_id,omitempty"`
	FileID             string `json:"file_id,omitempty"`
}

type Job struct {
	ID        string             `json:"ID,omitempty"`
	CreatedAt int64              `json:"CreatedAt,omitempty"`
	UpdatedAt int64              `json:"UpdatedAt,omitempty"`
	Name      string             `json:"Name,omitempty"`
	SlugName  string             `json:"SlugName,omitempty"`
	Org       string             `json:"Org,omitempty"`
	Users     []OrganizationUser `json:"Users,omitempty"` // 该职位的用户列表, 一个职位可以包涵多个用户
}

// Notification 通知Model
type Notification struct {
	ID          string `json:"ID,omitempty"`
	CreatedAt   int64  `json:"CreatedAt,omitempty"`
	UpdatedAt   int64  `json:"UpdatedAt,omitempty"`
	UserID      string `json:"UserID,omitempty"`
	Title       string `json:"Title,omitempty"`
	Type        string `json:"Type,omitempty"`
	Description string `json:"Description,omitempty"`
	Extra       string `json:"Extra,omitempty"`
	Status      string `json:"Status,omitempty"`
	Unread      bool   `json:"Unread,omitempty"`
}

// RecruitmentForm 招新表单
type RecruitmentForm struct {
	ID        string                 `json:"ID,omitempty"`
	CreatedAt int64                  `json:"CreatedAt,omitempty"`
	UpdatedAt int64                  `json:"UpdatedAt,omitempty"`
	FormName  string                 `json:"FormName,omitempty"` // 招新表名称
	Items     []*RecruitmentFormItem `json:"Items,omitempty"`    // 表单的项目
}

// RecruitmentFormItem 招新表单项目model
type RecruitmentFormItem struct {
	ID           string                 `json:"ID,omitempty"`
	CreatedAt    int64                  `json:"CreatedAt,omitempty"`
	UpdatedAt    int64                  `json:"UpdatedAt,omitempty"`
	Key          string                 `json:"Key,omitempty"` // 题目的标识，由前端生成
	Organization string                 `json:"Organization,omitempty"`
	Subject      string                 `json:"Subject,omitempty"`
	Options      string `json:"Options,omitempty"` // json object string
	Kind         string                 `json:"Kind,omitempty"`      // 种类 单选：radio 多选: checkbox 填空: words
	FormRefer    string                 `json:"FormRefer,omitempty"` // 所属表单的关联字段
}

// RecruitmentFormRecords 招新表单提交记录表
type RecruitmentFormRecords struct {
	ID               string                  `json:"ID,omitempty"`
	CreatedAt        int64                   `json:"CreatedAt,omitempty"`
	UpdatedAt        int64                   `json:"UpdatedAt,omitempty"`
	Mobile           string                  `json:"Mobile,omitempty"`           // 手机号
	Name             string                  `json:"Name,omitempty"`             // 姓名
	Major            string                  `json:"Major,omitempty"`            // 专业
	Age              int                     `json:"Age,omitempty"`              // 年龄
	SchoolStudentID  string                  `json:"SchoolStudentID,omitempty"`  // 学号
	AcceptDepartment Department              `json:"AcceptDepartment,omitempty"` // 申请加入的部门
	DepartmentRefer  string                  `json:"DepartmentRefer,omitempty"`
	Record           RecruitmentRecord       `json:"Record,omitempty"` // 归属的招新记录
	RecordID         string                  `json:"RecordID,omitempty"`
	Answers          []RecruitmentFormAnswer `json:"Answers,omitempty"`
	Status           int                     `json:"Status,omitempty"` // 0: 报名状态, 1: 已通过 2: 已拒绝
}

// RecruitmentRecord 招新记录表
type RecruitmentRecord struct {
	ID                   string           `json:"ID,omitempty"`
	CreatedAt            int64            `json:"CreatedAt,omitempty"`
	UpdatedAt            int64            `json:"UpdatedAt,omitempty"`
	Organization         string           `json:"Organization,omitempty"`         // 关联的组织
	CreateUser           string           `json:"CreateUser,omitempty"`           // 招新开启人
	RecruitmentForm      *RecruitmentForm `json:"RecruitmentForm,omitempty"`      // 本次招新的招新表
	RecruitmentFormRefer string           `json:"RecruitmentFormRefer,omitempty"` // 招新报名表的ID
	HasStart             bool             `json:"HasStart,omitempty"`             // 招新是否已经开始，一般在招新表提交后就正式开始
	HasEnd               bool             `json:"HasEnd,omitempty"`               // 招新是否已经结束
}

// RecruitmentFormAnswer 招新表单答案记录表
type RecruitmentFormAnswer struct {
	ID           string              `json:"ID,omitempty"`
	CreatedAt    int64               `json:"CreatedAt,omitempty"`
	UpdatedAt    int64               `json:"UpdatedAt,omitempty"`
	FormItem     RecruitmentFormItem `json:"FormItem,omitempty"` // 本答案关联的问题
	ItemRefer    string              `json:"ItemRefer,omitempty"`
	Answer       string              `json:"Answer,omitempty"` // 本问题的答案
	FormRecordID string              `json:"FormRecordID,omitempty"`
}

// RoleGroup 角色组
type RoleGroup struct {
	ID               string `json:"ID,omitempty"`
	CreatedAt        int64  `json:"CreatedAt,omitempty"`
	UpdatedAt        int64  `json:"UpdatedAt,omitempty"`
	GroupName        string `json:"GroupName,omitempty"`
	GroupDescription string `json:"GroupDescription,omitempty"`
	Org              string `json:"Org,omitempty"`
	Roles            []Role `json:"roles,omitempty"`
}

// Role 角色
type Role struct {
	ID                string             `json:"ID,omitempty"`
	CreatedAt         int64              `json:"CreatedAt,omitempty"`
	UpdatedAt         int64              `json:"UpdatedAt,omitempty"`
	Name              string             `json:"Name,omitempty"`
	Level             string             `json:"Level,omitempty"` // 角色等级 (user: 用户级, system: 系统级别, super: 超级)
	GroupID           string             `json:"GroupID,omitempty"`
	Org               string             `json:"Org,omitempty"`
	Group             *RoleGroup         `json:"Group,omitempty"`
	OrganizationUsers []OrganizationUser `json:"OrganizationUsers,omitempty"`
}

// OrganizationUserRole 组织用户和角色的关联表
type OrganizationUserRole struct {
	OrganizationUserID string `json:"OrganizationUserID,omitempty"`
	RoleID             string `json:"RoleID,omitempty"`
}

type Upgrade struct {
	ID          string `json:"ID,omitempty"`
	CreatedAt   int64  `json:"CreatedAt,omitempty"`
	UpdatedAt   int64  `json:"UpdatedAt,omitempty"`
	Version     string `json:"Version,omitempty"`
	Description string `json:"Description,omitempty"`
	Platform    string `json:"Platform,omitempty"`
	Body        string `json:"Body,omitempty"`
	URL         string `json:"URL,omitempty"`
}
