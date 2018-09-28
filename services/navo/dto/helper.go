package dto

import (
	"github.com/iron-kit/go-ironic/utils"
	orgPB "iunite.club/services/organization/proto"
	schoolPB "iunite.club/services/organization/proto/school"
)

func PBToOrganization(pb *orgPB.Organization) *Organization {
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
