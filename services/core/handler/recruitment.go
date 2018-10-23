package handler

import (
	"context"
	"time"

	ironic "github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/core/proto/recruitment"
)

type Recruitment struct {
	ironic.BaseHandler
}

func (r *Recruitment) model(ctx context.Context, modelName string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		panic(err.Error())
	}

	return conn.M(modelName)
}

func (r *Recruitment) AdjustOnePost(ctx context.Context, req *pb.AdjustOnePostRequest, rsp *pb.Response) error {
	RecruitmentFormRecordModel := r.model(ctx, "RecruitmentFormRecord")
	record := new(models.RecruitmentFormRecord)
	RecruitmentFormRecordModel.Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found this record")
	}

	record.Status = 0
	record.DepartmentID = bson.ObjectIdHex(req.DepartmentID)

	err := RecruitmentFormRecordModel.Update(bson.M{"_id": req.ID}, record)
	if err != nil {
		return r.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (r *Recruitment) PassedOnePost(ctx context.Context, req *pb.PassedOnePostRequest, rsp *pb.Response) error {
	RecruitmentFormRecordModel := r.model(ctx, "RecruitmentFormRecord")
	record := new(models.RecruitmentFormRecord)
	RecruitmentFormRecordModel.Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found this record")
	}

	record.Status = 1
	// record.DepartmentID = bson.ObjectIdHex(req.DepartmentID)

	err := RecruitmentFormRecordModel.Update(bson.M{"_id": req.ID}, record)
	if err != nil {
		return r.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (r *Recruitment) RefusedOnePost(ctx context.Context, req *pb.RefusedOnePostRequest, rsp *pb.Response) error {
	RecruitmentFormRecordModel := r.model(ctx, "RecruitmentFormRecord")
	record := new(models.RecruitmentFormRecord)
	RecruitmentFormRecordModel.Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found this record")
	}

	record.Status = 2

	err := RecruitmentFormRecordModel.Update(bson.M{"_id": req.ID}, record)
	if err != nil {
		return r.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (r *Recruitment) FindLatestRecruitmentRecord(ctx context.Context, req *pb.ByClubIDRequest, rsp *pb.RecruitmentRecordResponse) error {
	RecruitmentRecordModel := r.model(ctx, "RecruitmentRecord")
	record := new(models.RecruitmentRecord)
	RecruitmentRecordModel.Where(bson.M{"has_end": false}).Sort("-updated_at").FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found record")
	}

	rsp.Record = record.ToPB()

	return nil
}

func (r *Recruitment) AddRecruitmentRecord(ctx context.Context, req *pb.ByRecruitmentRecordBundle, rsp *pb.Response) error {
	RecruitmentRecordModel := r.model(ctx, "RecruitmentRecord")
	// RecruitmentRecordFormModel := r.model(ctx, "RecruitmentForm")
	record := new(models.RecruitmentRecord)
	// form := new(models.RecruitmentForm)

	//// 查找招新表
	//	//if err := RecruitmentRecordFormModel.Where(bson.M{"_id": req.FormID}).FindOne(form); err != nil {
	//	//	return r.Error(ctx).InternalServerError(err.Error())
	//	//}

	// 查找招新记录
	//if err := RecruitmentRecordModel.Where(bson.M{"_id": form.RecordID}).FindOne(record); err != nil {
	//	return r.Error(ctx).InternalServerError(err.Error())
	//}

	record.HasStart = true
	record.HasEnd = false
	record.CreateUserID = bson.ObjectIdHex(req.UserID)
	record.ClubID = bson.ObjectIdHex(req.ClubID)

	if err := RecruitmentRecordModel.Create(record); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	//// 更新招新记录
	//if err := RecruitmentRecordModel.Update(bson.M{"_id": form.ID}, record); err != nil {
	//	return r.Error(ctx).InternalServerError(err.Error())
	//}

	rsp.OK = true

	return nil
}

func (r *Recruitment) AddRecruitmentForm(ctx context.Context, req *pb.ByRecruitmentFormBundle, rsp *pb.Response) error {
	RecruitmentRecordFormModel := r.model(ctx, "RecruitmentForm")
	// RecruitmentRecordModel := r.model(ctx, "RecruitmentRecord")
	now := time.Now()
	form := models.RecruitmentForm{
		Name: now.Format("2006/01/02-招新表单"),
		// Fields:
	}
	fields := make([]models.RecruitmentFormField, 0)
	if len(req.RecordForm.Fields) > 0 {
		for _, f := range req.RecordForm.Fields {

			fields = append(fields, models.RecruitmentFormField{
				ID:      bson.NewObjectId(),
				Subject: f.Subject,
				Kind:    f.Kind,
				Key:     f.Key,
				Sort:    f.Sort,
				Options: hptypes.DecodeToMap(f.Options),
			})
		}
	}
	form.Fields = fields
	//record := models.RecruitmentRecord{
	//	ClubID:       bson.ObjectIdHex(req.ClubID),
	//	CreateUserID: bson.ObjectIdHex(req.UserID),
	//	HasStart:     false,
	//	HasEnd:       false,
	//}

	//if err := RecruitmentRecordModel.Create(&record); err != nil {
	//	return r.Error(ctx).BadRequest(err.Error())
	//}

	form.RecordID = bson.ObjectIdHex(req.RecordID)

	if err := RecruitmentRecordFormModel.Create(&form); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	return nil
}

func (r *Recruitment) FindRecruitmentFormDetails(ctx context.Context, req *pb.ByRecruitmentFormID, rsp *pb.RecruitmentRecordFromResponse) error {
	RecruitmentFormModel := r.model(ctx, "RecruitmentForm")
	form := new(models.RecruitmentForm)
	if err := RecruitmentFormModel.Where(bson.M{"_id": req.ID}).FindOne(form); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	//if err := RecruitmentFormModel.Where(bson.M{"_id": req.ID}).Populate("Fields").FindOne(form); err != nil {
	//	return r.Error(ctx).InternalServerError(err.Error())
	//}

	rsp.Form = form.ToPB()
	// panic("not implemented")
	return nil

}

func (r *Recruitment) AddRecruitmentFormRecord(ctx context.Context, req *pb.ByRecruitmentFormRecord, rsp *pb.Response) error {
	RecruitmentFormRecord := r.model(ctx, "RecruitmentForm")
	formRecord := &models.RecruitmentFormRecord{
		Mobile:          req.RecordFormRecord.Mobile,
		Name:            req.RecordFormRecord.Name,
		Major:           req.RecordFormRecord.Major,
		Age:             req.RecordFormRecord.Age,
		SchoolStudentID: req.RecordFormRecord.SchoolStudentID,
		DepartmentID:    bson.ObjectIdHex(req.RecordFormRecord.DepartmentID),
		RecordID:        bson.ObjectIdHex(req.RecordFormRecord.RecordID),
		Status:          req.RecordFormRecord.Status,
	}

	answers := make([]models.RecruitmentAnswer, 0)

	for _, v := range req.RecordFormRecord.Answers {
		answers = append(answers, models.RecruitmentAnswer{
			ID:      bson.NewObjectId(),
			ItemKey: bson.NewObjectId(),
			FormID:  bson.ObjectIdHex(v.FormID),
			Answer:  v.Answer,
		})
	}

	formRecord.Answers = answers

	if err := RecruitmentFormRecord.Create(formRecord); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	return nil

}

func (r *Recruitment) EndRecruitment(ctx context.Context, req *pb.ByRecruitmentID, rsp *pb.Response) error {
	// panic("not implemented")
	RecruitmentRecordModel := r.model(ctx, "RecruitmentRecord")
	record := new(models.RecruitmentRecord)
	RecruitmentRecordModel.Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found this recruitment record")
	}

	record.HasEnd = true

	if err := RecruitmentRecordModel.Update(bson.M{"_id": bson.ObjectIdHex(req.ID)}, record); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	return nil
}

func (r *Recruitment) FindRecruitmentRecordDetails(ctx context.Context, req *pb.ByRecruitmentID, rsp *pb.RecruitmentRecordResponse) error {
	RecruitmentRecordModel := r.model(ctx, "RecruitmentRecord")
	record := new(models.RecruitmentRecord)
	RecruitmentRecordModel.
		Where(bson.M{"_id": req.ID}).
		FindOne(record)

	if record.IsEmpty() {
		return r.Error(ctx).NotFound("Not found record")
	}

	rsp.Record = record.ToPB()

	return nil
}

func (r *Recruitment) FindRecruitmentsFormRecordList(ctx context.Context, req *pb.FindRecruitmentFormRecordRequest, rsp *pb.RecruitmentFormRecordsResponse) error {
	RecruitmentFormRecordModel := r.model(ctx, "RecruitmentFormRecord")

	records := make([]models.RecruitmentFormRecord, 0)
	query := RecruitmentFormRecordModel.Where(bson.M{
		"department_id": req.Department,
		"record_id":     req.RecordID,
		"status":        req.State,
	}).Query()

	total := query.Query().Count()
	if err := query.Query().Skip(int((req.Page - 1) * req.Limit)).Limit(int(req.Limit)).FindAll(&records); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	pbRecords := make([]*pb.RecruitmentFormRecord, 0)
	for _, v := range records {
		pbRecords = append(pbRecords, v.ToPB())
	}
	rsp.Records = pbRecords
	rsp.Total = int32(total)

	return nil
}
