package srv

import (
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/utils"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/organization/proto/school"
)

type SchoolService struct {
	ironic.Service
}

func (d *SchoolService) Model(name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(d.Ctx())

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (s *SchoolService) GetSchoolByID(id string) (*models.School, error) {
	SchoolModel := s.Model("School")
	foundSchool := models.School{}

	if !bson.IsObjectIdHex(id) {
		return &foundSchool, s.Error().NotFound("Not found school")
	}

	err := SchoolModel.FindByID(bson.ObjectIdHex(id), &foundSchool)
	if err != nil {
		return &foundSchool, s.Error().InternalServerError(err.Error())
	}

	return &foundSchool, nil

}

func (s *SchoolService) CreateSchool(req *pb.CreateSchoolRequest) (*models.School, error) {
	SchoolModel := s.Model("School")
	newSchool := models.School{
		Name:        req.Name,
		Description: req.Description,
		SlugName:    utils.Hans2Pinyin(req.Name, "_"),
	}

	if err := SchoolModel.Create(&newSchool); err != nil {
		return nil, s.Error().InternalServerError(err.Error())
	}

	return &newSchool, nil
}

func (s *SchoolService) GetSchoolList(req *pb.ListRequest, resp *pb.ListResponse) error {
	SchoolModel := s.Model("School")
	resp.Schools = make([]*pb.School, 0, 1)
	schools := []models.School{}
	total := SchoolModel.Count()
	resp.Total = int64(total)
	skipNum := 0
	if req.Page > 0 {
		skipNum = int((req.Page - 1) * req.Limit)
	}
	if err := SchoolModel.Where().Skip(skipNum).Limit(int(req.Limit)).FindAll(&schools); err != nil {
		return s.Error().InternalServerError(err.Error())
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

func (s *SchoolService) SearchSchools(req *pb.SearchSchoolsRequest, rsp *pb.ListResponse) error {
	SchoolModel := s.Model("School")
	rsp.Schools = make([]*pb.School, 0)
	schools := make([]models.School, 0)
	total := SchoolModel.Count(bson.M{
		"name": bson.RegEx{req.Search, "i"},
	})

	rsp.Total = int64(total)

	if err := SchoolModel.Where(bson.M{
		"name": bson.RegEx{req.Search, "i"},
	}).FindAll(&schools); err != nil {
		return s.Error().InternalServerError(err.Error())
	}

	for _, v := range schools {
		rsp.Schools = append(rsp.Schools, &pb.School{
			ID:          v.ID.Hex(),
			Name:        v.Name,
			SlugName:    v.SlugName,
			SchoolCode:  v.SchoolCode,
			Description: v.Description,
		})
	}

	return nil
}
