package handler

import (
	"context"
	"fmt"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/approved/proto"
)

type Approved struct {
	ironic.BaseHandler
}

func (aprd *Approved) model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (aprd *Approved) FindTemplates(ctx context.Context, req *pb.FindTemplatesRequest, rsp *pb.TemplatesResponse) error {
	ApprovedTemplateModel := aprd.model(ctx, "ApprovedTemplate")

	query := ApprovedTemplateModel.Where(bson.M{
		"club_id": req.ClubID,
	})

	fmt.Println(req)

	if len(req.Populate) > 0 {
		fmt.Println("populate")
		query = query.Populate(req.Populate...)
	}

	total := query.Query().Count()

	templates := make([]models.ApprovedTemplate, 0, req.Limit)

	query.Query().Skip(int((req.Page - 1) * req.Limit)).Limit(int(req.Limit)).FindAll(&templates)

	rsp.Total = int32(total)
	pbTemplates := make([]*pb.ApprovedTemplatePB, 0, len(templates))
	for _, v := range templates {
		pbTemplates = append(pbTemplates, v.ToPB())
	}

	rsp.Templates = pbTemplates
	return nil
}

func (aprd *Approved) PostTemplate(ctx context.Context, req *pb.PostTemplateRequest, rsp *pb.PostedResponse) error {
	ApprovedTemplateModel := aprd.model(ctx, "ApprovedTemplate")
	newTemplate := new(models.ApprovedTemplate)
	newTemplate.SetByPB(req.Template)

	if err := ApprovedTemplateModel.Create(newTemplate); err != nil {
		return aprd.Error(ctx).InternalServerError(err.Error())
	}
	rsp.Template = newTemplate.ToPB()
	rsp.OK = true
	return nil
}

func (aprd *Approved) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest, rsp *pb.UpdatedResponse) error {
	ApprovedTemplateModel := aprd.model(ctx, "ApprovedTemplate")

	willUpdatedFields := hptypes.DecodeToMap(req.Fields)

	if err := ApprovedTemplateModel.Update(bson.M{
		"_id": req.ID,
	}, bson.M{
		"$set": willUpdatedFields,
	}); err != nil {
		return aprd.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (aprd *Approved) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest, rsp *pb.DeletedResponse) error {
	ApprovedTemplateModel := aprd.model(ctx, "ApprovedTemplate")
	if err := ApprovedTemplateModel.Where(bson.M{"_id": req.ID}).Delete(); err != nil {
		return aprd.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}

func (aprd *Approved) ToggleTemplateEnableState(ctx context.Context, req *pb.ToggleEnableStateReq, rsp *pb.UpdatedResponse) error {
	ApprovedTemplateModel := aprd.model(ctx, "ApprovedTemplate")
	if err := ApprovedTemplateModel.Update(bson.M{"_id": req.ID}, bson.M{"$set": bson.M{"enabled": req.IsEnabled}}); err != nil {
		return aprd.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	return nil
}
