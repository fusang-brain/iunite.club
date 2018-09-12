package services

import (
	"github.com/iron-kit/go-ironic/micro-assistant"
	"github.com/iron-kit/go-ironic/utils"
	"iunite.club/models"
	pb "iunite.club/srv/organization-srv/proto/school"
)

type SchoolService struct {
	assistant.Service
	Error *assistant.ErrorManager
}

func (s *SchoolService) CreateSchool(req *pb.CreateSchoolRequest) (*models.School, error) {
	SchoolModel := s.Connection.M("School")
	newSchool := models.School{
		Name:        req.Name,
		Description: req.Description,
		SlugName:    utils.Hans2Pinyin(req.Name, "_"),
	}

	if err := SchoolModel.Create(&newSchool); err != nil {
		return nil, s.Error.InternalServerError(err.Error())
	}

	return &newSchool, nil
}

func (s *SchoolService) GetSchoolList(req *pb.ListRequest, resp *pb.ListResponse) error {
	SchoolModel := s.Connection.M("School")
	resp.Schools = make([]*pb.School, 0, 1)
	schools := []models.School{}
	total := SchoolModel.Count()
	resp.Total = int64(total)
	skipNum := 0
	if req.Page > 0 {
		skipNum = int((req.Page - 1) * req.Limit)
	}
	if err := SchoolModel.Find().Skip(skipNum).Limit(int(req.Limit)).Exec(&schools); err != nil {
		return s.Error.InternalServerError(err.Error())
	}
	for _, v := range schools {
		resp.Schools = append(resp.Schools, &pb.School{
			ID:          v.ID.Hex(),
			Name:        v.Name,
			SlugName:    v.SlugName,
			SchoolCode:  v.SchoolCode,
			Description: v.Description,
		})
	}

	return nil
}
