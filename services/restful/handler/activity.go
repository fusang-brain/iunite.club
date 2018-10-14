package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/client"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"

	"gopkg.in/mgo.v2/bson"

	approvedPB "iunite.club/services/core/proto/approved"

	userPB "iunite.club/services/user/proto"

	restful "github.com/emicklei/go-restful"
)

type ActivityHandler struct {
	BaseHandler

	userService     userPB.UserSrvService
	approvedService approvedPB.ApprovedService
}

func NewActivityHandler(c client.Client) *ActivityHandler {
	return &ActivityHandler{
		userService:     userPB.NewUserSrvService(UserSerivce, c),
		approvedService: approvedPB.NewApprovedService(CoreService, c),
	}
}

func (a *ActivityHandler) CreateActivity(req *restful.Request, rsp *restful.Response) {
	ctx := context.Background()
	params := struct {
		OrganizationUser  string   `json:"organizationUser,omitempty"`
		Subject           string   `json:"subject,omitempty"`
		Location          string   `json:"location,omitempty"`
		StartTime         int64    `json:"startTime,omitempty"`
		EndTime           int64    `json:"endTime,omitempty"`
		FoundingAmount    int64    `json:"foundingAmount,omitempty"`
		ParticipantsTotal int      `json:"ParticipantsTotal,omitempty"`
		Pictures          []string `json:"pictures,omitempty"`
		Attach            []string `json:"attach,omitempty"`
		ApprovedPersons   []string `json:"approvedPersons,omitempty"`
		CopyPersons       []string `json:"copyPersons,omitempty"`
	}{}

	if err := a.Bind(req, &params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	if err := a.Validate(&params); err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(err.Error()))
		return
	}

	type userClubProfileResp struct {
		resp *userPB.UserClubProfileResponse
		err  error
	}

	findClubProfileChan := make(chan userClubProfileResp, 1)
	go func() {
		res, err := a.userService.FindUserClubProfileByID(ctx, &userPB.FindUserClubProfileByIDRequest{ID: params.OrganizationUser})
		findClubProfileChan <- userClubProfileResp{
			resp: res,
			err:  err,
		}
	}()

	clubProfileReply := <-findClubProfileChan

	if clubProfileReply.err != nil {
		ErrorResponse(rsp, a.Error().BadRequest(clubProfileReply.err.Error()))
		return
	}
	now := time.Now()
	// objID := bson.NewObjectId()

	// fmt.Println(string(objID.Machine()))
	// bson.Obj
	// bson.M
	// j := make([]byte)
	// j, _ := objID.MarshalJSON()
	// fmt.Println(string(j))
	activityContent := map[string]interface{}{
		"_id":                bson.NewObjectId(),
		"created_at":         now,
		"updated_at":         now,
		"start_time":         time.Unix(params.StartTime/1e3, 0),
		"end_time":           time.Unix(params.EndTime/1e3, 0),
		"amount_funding":     params.FoundingAmount,
		"participants_total": params.ParticipantsTotal,
		"is_publish":         false,
		"pictures":           params.Pictures,
		"attachs":            params.Attach,
	}

	// fmt.Println(activityContent, "content")

	// fmt.Println(hptypes.EncodeToStruct(activityContent), "content")
	_, err := a.approvedService.Create(ctx, &approvedPB.CreateRequest{
		Title:         params.Subject,
		Kind:          "activity",
		Summary:       "无摘要",
		Status:        "pending",
		Description:   params.Subject,
		Content:       hptypes.EncodeToStruct(activityContent),
		ApprovedUsers: params.ApprovedPersons,
		CopyUsers:     params.CopyPersons,
		ClubID:        clubProfileReply.resp.ClubProfile.OrganizationID,
		CreatorID:     a.GetUserIDFromRequest(req),
		DepartmentID:  clubProfileReply.resp.ClubProfile.DepartmentID,
	})

	if err != nil {
		ErrorResponse(rsp, err)
		return
	}

	SuccessResponse(rsp, D{
		"Activity": activityContent,
	})

}
