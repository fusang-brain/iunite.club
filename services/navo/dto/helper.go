package dto

import (
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/go-ironic/utils"
	approvedPB "iunite.club/services/core/proto/approved"
	orgPB "iunite.club/services/organization/proto"
	schoolPB "iunite.club/services/organization/proto/school"
	storagePB "iunite.club/services/storage/proto"
	userPB "iunite.club/services/user/proto"
)

func PBToUser(pb *userPB.User) *User {
	if pb == nil {
		return &User{}
	}
	mobile := ""
	if pb.Profile != nil {
		mobile = pb.Profile.Mobile
	}
	return &User{
		ID:        pb.ID,
		CreatedAt: utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt: utils.ISOTime2MicroUnix(pb.UpdatedAt),
		IsTeacher: false,
		IsAdmin:   false,
		Username:  pb.Username,
		Mobile:    mobile,
		AreaCode:  "+86",
		Enabled:   pb.Enabled,
		Profile:   PBToProfile(pb.Profile),
	}
}

func PBToProfile(pb *userPB.Profile) *Profile {

	if pb == nil {
		return &Profile{}
	}
	updatedAt := hptypes.Timestamp(pb.UpdatedAt)
	createdAt := hptypes.Timestamp(pb.CreatedAt)
	birthday := hptypes.Timestamp(pb.Birthday)
	return &Profile{
		ID:        pb.ID,
		CreatedAt: utils.Time2MicroUnix(createdAt),
		// CreatedAt: utils.ISOTime2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt: utils.Time2MicroUnix(updatedAt),
		Avatar:    pb.Avatar,
		Gender:    pb.Gender,
		Birthday:  utils.Time2MicroUnix(birthday),
		Nickname:  pb.Nickname,
		UserID:    pb.UserID,
	}
}

func PBToOrganization(pb *orgPB.Organization) *Organization {
	if pb == nil {
		return &Organization{}
	}
	return &Organization{
		ID:          pb.ID,
		Name:        pb.Name,
		SlugName:    pb.Slug,
		Description: pb.Description,
		Logo:        pb.ClubProfile.Logo,
		Scale:       int(pb.ClubProfile.Scale),
		SchoolRefer: pb.SchoolID,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
	}
}

func PBToSchool(pb *schoolPB.School) *School {
	if pb == nil {
		return new(School)
	}
	return &School{
		ID:          pb.ID,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
		Name:        pb.Name,
		SlugName:    pb.SlugName,
		SchoolCode:  pb.SchoolCode,
		Description: pb.Description,
	}
}

func PBToDepartment(pb *orgPB.Organization) *Department {
	if pb == nil {
		return new(Department)
	}
	return &Department{
		ID:          pb.ID,
		SlugName:    pb.Slug,
		Name:        pb.Name,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
		ParentID:    pb.ParentID,
		Org:         pb.ClubID,
		Description: pb.Description,
	}
}

func PBToJob(pb *orgPB.Job) *Job {
	if pb == nil {
		return new(Job)
	}

	return &Job{
		ID:        pb.ID,
		Name:      pb.Name,
		SlugName:  pb.Slug,
		CreatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Org:       pb.ClubID,
		// CreatedAt: utils.ISOTime2MicroUnix()
		// Org: pb.
	}
}

func PBToOrganizationUser(pb *orgPB.UserClubProfile) *OrganizationUser {
	if pb == nil {
		return new(OrganizationUser)
	}
	rs := &OrganizationUser{
		ID:        pb.ID,
		CreatedAt: utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt: utils.ISOTime2MicroUnix(pb.UpdatedAt),
		// Kind:             1,
		// AcceptState:      1,
		// State:            pb.State,
		UserInfo:         PBToUser(pb.User),
		OrganizationInfo: PBToOrganization(pb.Organization),
		IsCreator:        pb.IsCreator,
		IsSuperManager:   pb.IsMaster,
		JoinTime:         utils.ISOTime2MicroUnix(pb.JoinTime),
		LeaveTime:        utils.ISOTime2MicroUnix(pb.LeaveTime),
		DepartmentID:     pb.DepartmentID,
		JobID:            pb.JobID,
	}
	if pb.State == 1 || pb.State == 2 {
		rs.Kind = 1
		rs.AcceptState = 1
		rs.State = int(pb.State)
	}
	if pb.State == 3 {
		rs.Kind = 0
		rs.AcceptState = 0
		rs.State = 0
	}
	if pb.State == 4 {
		rs.Kind = 0
		rs.AcceptState = 2
		rs.State = 0
	}
	return rs
}

func PBToApprovedProcess(pb *approvedPB.ApprovedFlowPB) *ApprovedProcess {
	if pb == nil {
		return new(ApprovedProcess)
	}

	ap := &ApprovedProcess{
		ID:          pb.ID,
		CreatedAt:   utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:   utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		ProcessSort: int(pb.Sort),
		ProcessType: pb.Kind,
		Options:     pb.Options,
		ApprovedID:  pb.ApprovedID,
		HandlerID:   pb.HandlerID,
		Archived:    false,
		Status:      int(pb.Status),
	}

	return ap
}

// func GetApprovedContent(pb *approvedPB.ApprovedPB) interface{} {
// 	if pb == nil {
// 		return nil
// 	}

// 	var res interface{}
// 	content := hptypes.DecodeToMap(pb.Content)
// 	switch pb.Kind {
// 	case "activity":

// 		activity := &Activity{
// 			ID:        content["ID"].(string),
// 			CreatedAt: utils.ISOTime2MicroUnix(content["CreatedAt"].(string)),
// 			UpdatedAt: utils.ISOTime2MicroUnix(content["UpdatedAt"].(string)),
// 			Subject: content["subject"].(string),
// 			Location: content["location"].(string),
// 			ApplicantID: content["applicant_id"].(string),
// 			Applicant: content["applicant"]
// 		}
// 	case "borrow":
// 	case "funding":

// 	}
// 	// res := interface{}{}

// }

func PBToFile(pb *storagePB.FilePB) *File {
	return &File{
		ID: pb.ID,
		// CreatedAt:
		Filename:         pb.FileKey,
		Path:             pb.Path,
		Ext:              pb.Ext,
		AbstractPath:     pb.Path,
		Host:             pb.Host,
		OriginalFilename: pb.OriginalFilename,
		Size:             pb.Size,
	}
}

func PBToApprovedTask(pb *approvedPB.ApprovedPB) *ApprovedTask {
	if pb == nil {
		return new(ApprovedTask)
	}

	at := &ApprovedTask{
		ID:           pb.ID,
		CreatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Name:         pb.Title,
		Description:  pb.Description,
		ApprovedType: pb.Kind,
		Content:      pb.Description,
		Status:       pb.Status,
		Summary:      pb.Summary,
	}

	if len(pb.Flows) > 0 {
		flows := make([]ApprovedProcess, 0, len(pb.Flows))

		for _, v := range pb.Flows {
			f := PBToApprovedProcess(v)
			flows = append(flows, *f)
		}
		at.ApprovedProcesses = flows
	}

	return at
}
