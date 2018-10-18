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

type BorrowHandler struct {
	BaseHandler

	userService     userPB.UserSrvService
	approvedService approvedPB.ApprovedService
}

type GoodsItem struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func NewBorrowHandler(c client.Client) *BorrowHandler {
	return &BorrowHandler{
		userService:     userPB.NewUserSrvService(UserSerivce, c),
		approvedService: approvedPB.NewApprovedService(CoreService, c),
	}
}

func (b *BorrowHandler) Create(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		OrganizationUser string      `json:"organizationUser"`
		Subject          string      `json:"subject"`
		Description      string      `json:"description"`
		GoodsItems       []GoodsItem `json:"goodsItems"`
		StartTime        int64       `json:"startTime"`
		EndTime          int64       `json:"endTime"`
		Pictures         []string    `json:"pictures"`
		Attach           []string    `json:"attach"`
		ApprovedPersons  []string    `json:"approvedPersons"`
		CopyPersons      []string    `json:"copyPersons"`
	}{}

	if err := b.Bind(req, &params); err != nil {
		ErrorResponse(rsp, b.Error().BadRequest(err.Error()))
		return
	}

	if err := b.Validate(&params); err != nil {
		ErrorResponse(rsp, b.Error().BadRequest(err.Error()))
		return
	}

	type userClubProfileResp struct {
		resp *userPB.UserClubProfileResponse
		err  error
	}

	findClubProfileChan := make(chan userClubProfileResp, 1)
	go func() {
		res, err := b.userService.FindUserClubProfileByID(ctx, &userPB.FindUserClubProfileByIDRequest{ID: params.OrganizationUser})
		findClubProfileChan <- userClubProfileResp{
			resp: res,
			err:  err,
		}
	}()

	clubProfileReply := <-findClubProfileChan

	if clubProfileReply.err != nil {
		ErrorResponse(rsp, b.Error().BadRequest(clubProfileReply.err.Error()))
		return
	}
	now := time.Now()

	content := map[string]interface{}{
		"_id":         bson.NewObjectId(),
		"subject":     params.Subject,
		"description": params.Description,
		"goods":       params.GoodsItems,
		"start_time":  time.Unix(params.StartTime/1e3, 0),
		"end_time":    time.Unix(params.EndTime/1e3, 0),
		"created_at":  now,
		"updated_at":  now,
		"pictures":    params.Pictures,
		"attachs":     params.Attach,
	}

	_, err := b.approvedService.Create(ctx, &approvedPB.CreateRequest{
		Title:         params.Subject,
		Kind:          "borrow",
		Summary:       "",
		Status:        "pending",
		Description:   params.Subject,
		Content:       hptypes.EncodeToStruct(content),
		ApprovedUsers: params.ApprovedPersons,
		CopyUsers:     params.CopyPersons,
		ClubID:        clubProfileReply.resp.ClubProfile.OrganizationID,
		CreatorID:     b.GetUserIDFromRequest(req),
		DepartmentID:  clubProfileReply.resp.ClubProfile.DepartmentID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Goods": content,
	})
}
