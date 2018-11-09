package handler

import (
	"context"
	"fmt"

	"github.com/iron-kit/go-ironic"

	restful "github.com/emicklei/go-restful"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/micro/go-micro/client"
	"github.com/skip2/go-qrcode"
	recruitmentPB "iunite.club/services/core/proto/recruitment"
	"iunite.club/services/restful/dto"
)

type RecruitmentHandler struct {
	BaseHandler

	recruitmentService recruitmentPB.RecruitmentService
}

func NewRecruitmentHandler(c client.Client) *RecruitmentHandler {
	return &RecruitmentHandler{
		recruitmentService: recruitmentPB.NewRecruitmentService(CoreService, c),
	}
}

func (self *RecruitmentHandler) AdjustOnePost(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID           string `json:"id"`
		DepartmentID string `json:"dept"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	_, err := self.recruitmentService.AdjustOnePost(ctx, &recruitmentPB.AdjustOnePostRequest{
		ID:           params.ID,
		DepartmentID: params.DepartmentID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (self *RecruitmentHandler) PassedOnePost(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string `json:"id,omitempty"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	passedResp, err := self.recruitmentService.PassedOnePost(ctx, &recruitmentPB.PassedOnePostRequest{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
	}

	if !passedResp.OK {
		ErrorResponse(rsp, self.Error().BadRequest("error action"))
	}

	SuccessResponse(rsp, D{})
	return
}

func (self *RecruitmentHandler) RefusedOnePost(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string `json:"id,omitempty"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	_, err := self.recruitmentService.RefusedOnePost(ctx, &recruitmentPB.RefusedOnePostRequest{
		ID: params.ID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (self *RecruitmentHandler) GetLastestRecruitmentRecord(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ClubID string `query:"org"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	recordResp, err := self.recruitmentService.FindLatestRecruitmentRecord(ctx, &recruitmentPB.ByClubIDRequest{ClubID: params.ClubID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Record": dto.PBToRecruitmentRecord(recordResp.Record),
	})

	return
}

func (self *RecruitmentHandler) AddRecruitmentRecord(req *restful.Request, rsp *restful.Response) {
	params := struct {
		Creator      string `json:"userID,omitempty"`
		Organization string `json:"organization,omitempty"`
		FormID       string `json:"formID,omitempty"`
	}{}
	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if _, err := self.recruitmentService.AddRecruitmentRecord(ctx, &recruitmentPB.ByRecruitmentRecordBundle{
		ClubID: params.Organization,
		UserID: self.GetUserIDFromRequest(req),
	}); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (self *RecruitmentHandler) AddRecruitmentForm(req *restful.Request, rsp *restful.Response) {
	type Option struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	}

	type FormItem struct {
		Key     string   `json:"key"`
		Label   string   `json:"label"`
		Kind    string   `json:"kind"`
		Options []Option `json:"options"`
	}

	params := struct {
		RecordID     string     `json:"recordID"`
		Organization string     `json:"organization"`
		FormConfig   []FormItem `json:"formConfig"`
	}{}

	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	fields := make([]*recruitmentPB.RecruitmentFormField, 0)
	if len(params.FormConfig) > 0 {
		for index, f := range params.FormConfig {
			opts := map[string]interface{}{
				"Opts": f.Options,
			}
			fields = append(fields, &recruitmentPB.RecruitmentFormField{
				Subject: f.Label,
				Key:     f.Key,
				Options: hptypes.EncodeToStruct(opts),
				Sort:    int32(index),
			})
		}

	}

	_, err := self.recruitmentService.AddRecruitmentForm(ctx, &recruitmentPB.ByRecruitmentFormBundle{
		RecordID: params.RecordID,
		ClubID:   params.Organization,
		UserID:   self.GetUserIDFromRequest(req),
		RecordForm: &recruitmentPB.RecruitmentForm{
			Fields: fields,
		},
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
}

func (self *RecruitmentHandler) GetRecruitmentFormDetails(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string `query:"id`
	}{}

	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	formRsp, err := self.recruitmentService.FindRecruitmentFormDetails(ctx, &recruitmentPB.ByRecruitmentFormID{ID: params.ID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"FormDetails": dto.PBToRecruitmentForm(formRsp.Form),
	})

	return
}

func (self *RecruitmentHandler) GetRecruitmentRecordDetails(req *restful.Request, rsp *restful.Response) {
	params := struct {
		ID string `query:"id`
	}{}

	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}
	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	detailsResp, err := self.recruitmentService.FindRecruitmentRecordDetails(ctx, &recruitmentPB.ByRecruitmentID{ID: params.ID})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"RecordDetails": dto.PBToRecruitmentRecord(detailsResp.Record),
	})
	return
}
func (self *RecruitmentHandler) AddRecruitmentFormRecord(req *restful.Request, rsp *restful.Response) {

	type AnswerItem struct {
		ItemID string `json:"itemID"`
		Answer string `json:"answer"`
	}

	params := struct {
		Mobile          string       `json:"mobile,omitempty"`
		Name            string       `json:"name,omitempty"`
		Major           string       `json:"major,omitempty"`
		Age             int32        `json:"age,omitempty"`
		SchoolStudentID string       `json:"school_student_id,omitempty"`
		RecordID        string       `json:"record_id,omitempty"`
		Department      string       `json:"department,omitempty"`
		Answers         []AnswerItem `json:"answers,omitempty"`
	}{}

	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	answers := make([]*recruitmentPB.RecruitmentAnswer, 0)
	for _, v := range params.Answers {
		answers = append(answers, &recruitmentPB.RecruitmentAnswer{
			Answer:  v.Answer,
			ItemKey: v.ItemID,
		})
	}

	_, err := self.recruitmentService.AddRecruitmentFormRecord(ctx, &recruitmentPB.ByRecruitmentFormRecord{
		RecordFormRecord: &recruitmentPB.RecruitmentFormRecord{
			Mobile:          params.Mobile,
			Name:            params.Name,
			Major:           params.Major,
			Age:             params.Age,
			SchoolStudentID: params.SchoolStudentID,
			RecordID:        params.RecordID,
			DepartmentID:    params.Department,
			Answers:         answers,
		},
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}

func (self *RecruitmentHandler) GetRecruitmentFormRecordList(req *restful.Request, rsp *restful.Response) {
	params := struct {
		Page       int32  `json:"page,omitempty"`
		Limit      int32  `json:"limit,omitempty"`
		Department string `json:"department,omitempty"`
		State      int32  `json:"state,omitempty"`
		RecordID   string `json:"recordID,omitempty"`
	}{}

	ctx := context.Background()

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	listResp, err := self.recruitmentService.FindRecruitmentsFormRecordList(ctx, &recruitmentPB.FindRecruitmentFormRecordRequest{
		Page:       params.Page,
		Limit:      params.Limit,
		Department: params.Department,
		State:      params.State,
		RecordID:   params.RecordID,
	})

	// listResp.

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	records := make([]*dto.RecruitmentFormRecords, 0)
	for _, record := range listResp.Records {
		records = append(records, dto.PBToRecruitmentFormRecords(record))
	}

	SuccessResponseWithPage(rsp, int64(params.Page), int64(params.Limit), int64(listResp.Total), records)
	return
}

func (self *RecruitmentHandler) DownloadQRCode(req *restful.Request, rsp *restful.Response) {
	// self.recruitment
	params := struct {
		ID string `json:"id,omitempty"`
	}{}

	if err := self.BindAndValidate(req, &params); err != nil {
		ErrorResponse(rsp, err)
		return
	}

	ctx := context.Background()
	detailsResp, err := self.recruitmentService.FindRecruitmentRecordDetails(ctx, &recruitmentPB.ByRecruitmentID{
		ID: params.ID,
	})
	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	record := detailsResp.Record

	png, err := qrcode.Encode(fmt.Sprintf("%s/common/recruitment?recordID=%s", "im.iunite.club", record.ID), qrcode.Medium, 256)

	if err != nil {
		ErrorResponse(rsp, self.Error().InternalServerError(err.Error()))
		return
	}

	rsp.AddHeader(ironic.HeaderContentDisposition, fmt.Sprintf("%s, filename=%s_qr.png", "attachment", record.ID))
	rsp.Write(png)

	return
}

func (self *RecruitmentHandler) EndRecruitment(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		RecordID string `query:"recordID"`
	}{}

	if err := self.Bind(req, &params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	if err := self.Validate(&params); err != nil {
		ErrorResponse(rsp, self.Error().BadRequest(err.Error()))
		return
	}

	_, err := self.recruitmentService.EndRecruitment(ctx, &recruitmentPB.ByRecruitmentID{ID: params.RecordID})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{})
	return
}
