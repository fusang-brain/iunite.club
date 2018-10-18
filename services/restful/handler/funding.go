package handler

import (
	"context"
	"time"

	restful "github.com/emicklei/go-restful"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/micro/go-micro/client"
	"gopkg.in/mgo.v2/bson"
	approvedPB "iunite.club/services/core/proto/approved"
	userPB "iunite.club/services/user/proto"
)

type FundingHandler struct {
	BaseHandler

	userService     userPB.UserSrvService
	approvedService approvedPB.ApprovedService
}

func NewFundingHandler(c client.Client) *FundingHandler {
	return &FundingHandler{
		userService:     userPB.NewUserSrvService(UserSerivce, c),
		approvedService: approvedPB.NewApprovedService(CoreService, c),
	}
}

func (f *FundingHandler) Create(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		OrganizationUser string   `json:"organizationUser" validate:"nonzero,objectid"`
		Purpose          string   `json:"purpose"`
		Fee              int64    `json:"fee"`
		Pictures         []string `json:"pictures"`
		Attach           []string `json:"attach"`
		ApprovedPersons  []string `json:"approvedPersons"`
		CopyPersons      []string `json:"copyPersons"`
	}{}

	if err := f.Bind(req, &params); err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
		return
	}

	if err := f.Validate(&params); err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(err.Error()))
		return
	}

	type userClubProfileResp struct {
		resp *userPB.UserClubProfileResponse
		err  error
	}

	findClubProfileChan := make(chan userClubProfileResp, 1)
	go func() {
		res, err := f.userService.FindUserClubProfileByID(ctx, &userPB.FindUserClubProfileByIDRequest{ID: params.OrganizationUser})
		findClubProfileChan <- userClubProfileResp{
			resp: res,
			err:  err,
		}
	}()

	clubProfileReply := <-findClubProfileChan

	if clubProfileReply.err != nil {
		ErrorResponse(rsp, f.Error().BadRequest(clubProfileReply.err.Error()))
		return
	}
	now := time.Now()

	content := map[string]interface{}{
		"_id":              bson.NewObjectId(),
		"created_at":       now,
		"updated_at":       now,
		"apply_purpose":    params.Purpose,
		"amount_apply_fee": params.Fee,
		"pictures":         params.Pictures,
		"attachs":          params.Attach,
	}

	_, err := f.approvedService.Create(ctx, &approvedPB.CreateRequest{
		Title:         params.Purpose,
		Kind:          "funding",
		Summary:       "无摘要",
		Status:        "pending",
		Description:   params.Purpose,
		Content:       hptypes.EncodeToStruct(content),
		ApprovedUsers: params.ApprovedPersons,
		CopyUsers:     params.CopyPersons,
		ClubID:        clubProfileReply.resp.ClubProfile.OrganizationID,
		CreatorID:     f.GetUserIDFromRequest(req),
		DepartmentID:  clubProfileReply.resp.ClubProfile.DepartmentID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Funding": content,
	})
}
