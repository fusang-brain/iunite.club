package dto

import (
	"github.com/iron-kit/go-ironic/utils"
	orgPB "iunite.club/services/organization/proto"
	schoolPB "iunite.club/services/organization/proto/school"
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
	return &Profile{
		ID:        pb.ID,
		CreatedAt: utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt: utils.ISOTime2MicroUnix(pb.UpdatedAt),
		Avatar:    pb.Avatar,
		Gender:    pb.Gender,
		Birthday:  utils.ISOTime2MicroUnix(pb.Birthday),
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
