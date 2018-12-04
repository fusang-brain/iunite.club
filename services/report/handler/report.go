package handler

import (
	"context"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/report/proto"
)

type Report struct {
	ironic.BaseHandler
}

func (r *Report) model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (r *Report) FindReports(ctx context.Context, req *pb.FindReportsRequest, rsp *pb.ReportsResponse) error {
	ReportModel := r.model(ctx, "Report")

	query := ReportModel.Where(bson.M{
		// "user_id": req.UserID,
		"club_id": req.ClubID,
		"$or": []bson.M{
			//{
			//	"user_id": req.UserID,
			//},
			{
				"receivers": req.UserID,
			},
		},
	}).Query()

	total := query.Query().Count()

	reports := make([]models.Report, 0)

	if err := query.
		Query().
		Populate("User", "User.Profile").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&reports); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	pbReports := make([]*pb.ReportPB, 0)
	for _, v := range reports {
		pbReports = append(pbReports, v.ToPB())
	}

	rsp.Reports = pbReports
	rsp.Total = int32(total)

	return nil
}

func (r *Report) PostReport(ctx context.Context, req *pb.PostReportBundle, rsp *pb.PostReportResponse) error {
	// panic("not implemented")
	ReportModel := r.model(ctx, "Report")
	ReportTemplateModel := r.model(ctx, "ReportTemplate")
	newReport := new(models.Report)
	newReport.Body = req.Body
	newReport.ClubID = bson.ObjectIdHex(req.ClubID)
	newReport.UserID = bson.ObjectIdHex(req.UserID)
	// 发布默认汇报
	if req.Kind == pb.PostReportBundle_Default {
		if len(req.Receivers) > 0 {
			receivers := make([]bson.ObjectId, 0, len(req.Receivers))
			for _, v := range req.Receivers {
				receivers = append(receivers, bson.ObjectIdHex(v))
			}
			newReport.Receivers = receivers
		}
		newReport.Kind = "DEFAULT"
		newReport.Title = req.Title
		newReport.Description = req.Description
	}

	if req.Kind == pb.PostReportBundle_Template {
		foundTemplate := new(models.ReportTemplate)

		ReportTemplateModel.Where(bson.M{
			"_id": req.TemplateID,
		}).FindOne(foundTemplate)

		if foundTemplate.IsEmpty() {
			return r.Error(ctx).NotFound("Not found report template")
		}
		newReport.Kind = "TEMPLATE"
		newReport.Title = foundTemplate.Title
		newReport.Description = foundTemplate.Description
		newReport.Receivers = foundTemplate.Receivers
		newReport.Results = hptypes.DecodeToMap(req.Results)
		newReport.TemplateID = bson.ObjectIdHex(req.TemplateID)
	}

	if err := ReportModel.Create(newReport); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	rsp.Report = newReport.ToPB()

	return nil
}

func (r *Report) PostTemplate(ctx context.Context, req *pb.PostTemplateBundle, rsp *pb.PostTemplateResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")
	newReportTemplate := &models.ReportTemplate{
		Title:       req.Title,
		Description: req.Description,
		ClubID:      bson.ObjectIdHex(req.ClubID),
		UserID:      bson.ObjectIdHex(req.UserID),
	}

	if len(req.Receivers) > 0 {
		receivers := make([]bson.ObjectId, 0, len(req.Receivers))
		for _, v := range req.Receivers {
			receivers = append(receivers, bson.ObjectIdHex(v))
		}
		newReportTemplate.Receivers = receivers
	}

	if len(req.CustomFields) > 0 {
		fields := make([]models.TemplateCustomField, 0, len(req.CustomFields))

		for _, v := range req.CustomFields {
			fields = append(fields, models.TemplateCustomField{
				Key:     v.Key,
				Kind:    v.Kind,
				Label:   v.Label,
				Options: hptypes.DecodeToMap(v.Options),
				Sort:    v.Sort,
			})
		}

		newReportTemplate.CustomFields = fields
	}

	if req.Config != nil {
		config := models.TemplateConfig{
			Kind: req.Config.Kind,
		}

		if req.Config.Kind == "custom" {
			config.StartTime = hptypes.Timestamp(req.Config.StartTime)
			config.EndTime = hptypes.Timestamp(req.Config.EndTime)
		} else if req.Config.Kind == "weeks" {
			config.Weeks = req.Config.Weeks
		}

		newReportTemplate.Config = config
	}

	if err := ReportTemplateModel.Create(newReportTemplate); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	rsp.ReportTemplate = newReportTemplate.ToPB()

	return nil
}

func (r *Report) FindTemplates(ctx context.Context, req *pb.FindTemplatesRequest, rsp *pb.TemplatesResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")
	UserModel := r.model(ctx, "User")

	query := ReportTemplateModel.Where(bson.M{
		"club_id": req.ClubID,
	})

	total := query.Query().Count()
	templates := make([]models.ReportTemplate, 0, req.Limit)
	query.
		Query().
		Populate("Creator", "Creator.Profile").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&templates)

	for key, value := range templates {
		// fmt.Println("==== getUser ", value.Receivers)
		users := make([]models.User, 0, len(value.Receivers))
		UserModel.Where(bson.M{"_id": bson.M{"$in": value.Receivers}}).Populate("Profile").FindAll(&users)
		// fmt.Println("users", users)
		// value.SetReceiverInfo(users)
		templates[key].SetReceiverInfo(users)
		// fmt.Println(templates[key].ReceiversInfo)
	}

	// UserModel.Where(bson.M{"_id": bson.M{"$in": }})


	pbTemplates := make([]*pb.ReportTemplatePB, 0, req.Limit)

	for _, v := range templates {
		pbTemplates = append(pbTemplates, v.ToPB())
	}

	rsp.Total = int32(total)
	rsp.Templates = pbTemplates

	return nil
}

func (r *Report) FindPendingTemplates(ctx context.Context, req *pb.FindTemplatesRequest, rsp *pb.PendingTemplateResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")

	query := ReportTemplateModel.Where(bson.M{
		"club_id": req.ClubID,
		"enable": true,
	})

	total := query.Query().Count()
	templates := make([]models.ReportTemplate, 0, req.Limit)
	query.
		Query().
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&templates)


	pbTemplates := make([]*pb.ReportTemplatePB, 0, req.Limit)

	for _, v := range templates {
		pbTemplates = append(pbTemplates, v.ToPB())
	}

	rsp.Total = int32(total)
	rsp.Templates = pbTemplates

	return nil
}

func (r *Report) FindOneReport(ctx context.Context, req *pb.ByIDRequest, rsp *pb.ReportResponse) error {
	ReportModel := r.model(ctx, "Report")
	foundReport := new(models.Report)
	if err := ReportModel.Where(bson.M{
		"_id": req.ID,
	}).Populate("User", "User.Profile").FindOne(foundReport); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.Report = foundReport.ToPB()
	return nil
}

func (r *Report) FindOneTemplate(ctx context.Context, req *pb.ByTemplateIDRequest, rsp *pb.TemplateResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")
	foundReport := new(models.ReportTemplate)
	if err := ReportTemplateModel.Where(bson.M{
		"_id": req.ID,
	}).FindOne(foundReport); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}

	rsp.Template = foundReport.ToPB()
	return nil
}

func (r *Report) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateBundle, rsp *pb.UpdatedTemplateResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")
	willUpdatedFields := hptypes.DecodeToMap(req.Fields)
	if err := ReportTemplateModel.Update(bson.M{"_id": req.ID}, bson.M{
		"$set": willUpdatedFields,
	}); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (r *Report) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest, rsp *pb.DeletedResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")

	if err := ReportTemplateModel.Where(bson.M{"_id": req.ID}).Delete(); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (r *Report) ToggleTemplateEnableState(ctx context.Context, req *pb.ToggleTemplateEnableReq, rsp *pb.UpdatedTemplateResponse) error {
	ReportTemplateModel := r.model(ctx, "ReportTemplate")

	if err := ReportTemplateModel.Update(bson.M{"_id": req.ID}, bson.M{
		"$set": bson.M{
			"enable": req.IsEnabled,
		},
	}); err != nil {
		return r.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}
