package handler

import (
	"context"
	"fmt"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"

	"iunite.club/models"

	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	iunite_club_srv_core "iunite.club/services/core/proto/approved"
)

type ApprovedHandler struct {
	ironic.BaseHandler
}

func (a *ApprovedHandler) Model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

func (a *ApprovedHandler) List(ctx context.Context, req *iunite_club_srv_core.ListRequest, rsp *iunite_club_srv_core.ListResponse) error {

	ApprovedModel := a.Model(ctx, "Approved")
	condition := bson.M{
		"kind":             req.Kind,
		"status":           req.Status,
		"club_id":          bson.ObjectIdHex(req.ClubID),
		"flows.handler_id": bson.ObjectIdHex(req.HandlerID),
	}

	if req.Search != "" {
		condition["title"] = bson.RegEx{Pattern: req.Search, Options: "i"}
	}

	approveds := make([]models.Approved, 0)

	total := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Count()

	err := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&approveds)

	if err != nil {
		return a.Error(ctx).BadRequest(err.Error())
	}

	rsp.Total = int64(total)
	listSize := len(approveds)

	aps := make([]*iunite_club_srv_core.ApprovedPB, 0, listSize)
	for _, v := range approveds {
		aps = append(aps, v.ToPB())
	}
	rsp.Approveds = aps
	return nil
	// panic("not implemented")
}

func (a *ApprovedHandler) ListV2(ctx context.Context, req *iunite_club_srv_core.ListV2Request, rsp *iunite_club_srv_core.ListResponse) error {
	ApprovedModel := a.Model(ctx, "Approved")
	flowElemMatch := bson.M{
		"handler_id": req.HandlerID,
	}

	// if req.FlowStatus == "pending" {
	// 	flowElemMatch["kind"] = "approved"
	// 	flowElemMatch["status"] = 1
	// }

	switch req.FlowStatus {
	case "pending":
		flowElemMatch["kind"] = "approved"
		flowElemMatch["status"] = 1
	case "finished":
		flowElemMatch["kind"] = "approved"
		flowElemMatch["status"] = bson.M{"$gt": 1}
	case "copy":
		flowElemMatch["kind"] = "copy"
		if req.ReadState == "unread" {
			flowElemMatch["status"] = 0
		} else if req.ReadState == "alread" {
			flowElemMatch["status"] = bson.M{"$gte": 1}
		}
	}

	condition := bson.M{
		"club_id": req.ClubID,
		"flows": bson.M{
			"$elemMatch": flowElemMatch,
		},
	}

	if req.Search != "" {
		condition["title"] = bson.RegEx{Pattern: req.Search, Options: "i"}
	}

	approveds := make([]models.Approved, 0)

	total := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Count()

	err := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&approveds)

	if err != nil {
		return a.Error(ctx).BadRequest(err.Error())
	}

	rsp.Total = int64(total)
	listSize := len(approveds)

	aps := make([]*iunite_club_srv_core.ApprovedPB, 0, listSize)
	for _, v := range approveds {
		aps = append(aps, v.ToPB())
	}
	rsp.Approveds = aps
	return nil
}

func (a *ApprovedHandler) ListByPusher(ctx context.Context, req *iunite_club_srv_core.ListByPusherRequest, rsp *iunite_club_srv_core.ListResponse) error {
	// panic("not implemented")
	ApprovedModel := a.Model(ctx, "Approved")
	condition := bson.M{
		"club_id":   bson.ObjectIdHex(req.ClubID),
		"pusher_id": bson.ObjectIdHex(req.UserID),
	}

	if req.Search != "" {
		condition["title"] = bson.RegEx{Pattern: req.Search, Options: "i"}
	}

	approveds := make([]models.Approved, 0)

	total := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Count()

	err := ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler").
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&approveds)

	if err != nil {
		return a.Error(ctx).BadRequest(err.Error())
	}

	rsp.Total = int64(total)
	listSize := len(approveds)

	aps := make([]*iunite_club_srv_core.ApprovedPB, 0, listSize)
	for _, v := range approveds {
		aps = append(aps, v.ToPB())
	}
	rsp.Approveds = aps
	return nil
}

func (a *ApprovedHandler) WaitingExecuteList(context.Context, *iunite_club_srv_core.ListByCountRequest, *iunite_club_srv_core.ListResponse) error {
	panic("not implemented")
}

func (a *ApprovedHandler) Details(ctx context.Context, req *iunite_club_srv_core.DetailsRequest, rsp *iunite_club_srv_core.ApprovedResponse) error {
	ApprovedModel := a.Model(ctx, "Approved")
	approved := new(models.Approved)
	ApprovedModel.
		Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).
		Populate("Flows", "Flows.Handler", "Flows.Handler.Profile", "Pusher", "Pusher.Profile", "Department").
		Aggregate([]bson.M{
			{
				"$lookup": bson.M{
					"as":   "content.pictureObjects",
					"from": "file",
					"let":  bson.M{"pictures": "$content.pictures"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$pictures"},
										"then": bson.M{"$in": []string{"$_id", "$$pictures"}},
										"else": bson.M{"$eq": []string{"$_id", "$$pictures"}},
									},
								},
							},
						},
					},
				},
			},
			{
				"$lookup": bson.M{
					"as":   "content.attachObjects",
					"from": "file",
					"let":  bson.M{"attachs": "$content.attachs"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$attachs"},
										"then": bson.M{"$in": []string{"$_id", "$$attachs"}},
										"else": bson.M{"$eq": []string{"$_id", "$$attachs"}},
									},
								},
							},
						},
					},
				},
			},
		}).
		FindOne(approved)
	rsp.Approved = approved.ToPB()
	return nil
}

func (a *ApprovedHandler) DetailsByContentID(ctx context.Context, req *iunite_club_srv_core.DetailsRequest, rsp *iunite_club_srv_core.ApprovedResponse) error {
	ApprovedModel := a.Model(ctx, "Approved")
	approved := new(models.Approved)
	ApprovedModel.
		Where(bson.M{"content._id": bson.ObjectIdHex(req.ID)}).
		Populate("Flows", "Flows.Handler", "Flows.Handler.Profile", "Pusher", "Pusher.Profile", "Department").
		Aggregate([]bson.M{
			{
				"$lookup": bson.M{
					"as":   "content.pictureObjects",
					"from": "file",
					"let":  bson.M{"pictures": "$content.pictures"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$pictures"},
										"then": bson.M{"$in": []string{"$_id", "$$pictures"}},
										"else": bson.M{"$eq": []string{"$_id", "$$pictures"}},
									},
								},
							},
						},
					},
				},
			},
			{
				"$lookup": bson.M{
					"as":   "content.attachObjects",
					"from": "file",
					"let":  bson.M{"attachs": "$content.attachs"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$attachs"},
										"then": bson.M{"$in": []string{"$_id", "$$attachs"}},
										"else": bson.M{"$eq": []string{"$_id", "$$attachs"}},
									},
								},
							},
						},
					},
				},
			},
		}).
		FindOne(approved)
	rsp.Approved = approved.ToPB()
	return nil
}

func (a *ApprovedHandler) Execute(ctx context.Context, req *iunite_club_srv_core.ExecuteRequest, rsp *iunite_club_srv_core.Response) error {
	// panic("not implemented")
	ApprovedModel := a.Model(ctx, "Approved")
	ApprovedFlowModel := a.Model(ctx, "ApprovedFlow")
	// OrganizationModel := a.Model(ctx, "Organization")

	foundApproved := new(models.Approved)
	ApprovedModel.Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).Populate("Flows").FindOne(foundApproved)

	if foundApproved.ClubID.Hex() != req.ClubID {
		return a.Error(ctx).BadRequest("Not access")
	}
	executeState := map[bool]int{
		true:  2, // 同意
		false: 3, // 拒绝
	}

	currentKey := 0
	processTotal := len(foundApproved.Flows)
	stateCount := 0

	for k := range foundApproved.Flows {
		vp := &foundApproved.Flows[k]

		if vp.HandlerID.Hex() == req.UserID {
			if vp.Status > 1 {
				rsp.OK = true
				return nil
			}
			currentKey = k + 1
			vp.Options = req.Options
			vp.Status = executeState[req.Result]

			if !req.Result {
				stateCount = stateCount + vp.Status
				break
			}
		}

		// 准备通知下一个人
		if currentKey != 0 && currentKey == k {
			vp.Status = vp.Status + 1
			// TODO 发送通知给下一个人
		}

		if vp.Status == 1 {
			continue
		}

		stateCount = stateCount + vp.Status
	}

	taskStatus := "pending"
	if stateCount > 0 {
		taskStatus = "pending"
	}
	if stateCount == processTotal*2 {
		taskStatus = "pass"
		defer func() {
			// TODO 通知发起人，该审批已经通过
		}()
	}

	if stateCount%2 != 0 {
		taskStatus = "refuse"
		defer func() {
			// TODO 通知发起人，该审批已经被拒绝
		}()
	}

	foundApproved.Status = taskStatus

	if err := ApprovedModel.Update(bson.M{"_id": foundApproved.ID}, foundApproved); err != nil {
		return a.Error(ctx).InternalServerError(err.Error())
	}

	// if err :=
	for _, v := range foundApproved.Flows {
		if err := ApprovedFlowModel.Update(bson.M{"_id": v.ID}, &v); err != nil {
			return a.Error(ctx).InternalServerError(err.Error())
		}
	}

	rsp.OK = true
	return nil
}

func (a *ApprovedHandler) Create(ctx context.Context, req *iunite_club_srv_core.CreateRequest, rsp *iunite_club_srv_core.ApprovedResponse) error {
	// panic("not implemented")
	ApprovedModel := a.Model(ctx, "Approved")
	ApprovedFlowModel := a.Model(ctx, "ApprovedFlow")
	approvedUsers := req.ApprovedUsers
	auSize := len(approvedUsers)

	if auSize <= 0 {
		return a.Error(ctx).BadRequest("approved users can not be null")
	}

	flows := make([]models.ApprovedFlow, 0, auSize)

	approved := new(models.Approved)
	approved.DepartmentID = bson.ObjectIdHex(req.DepartmentID)
	approved.Kind = req.Kind
	approved.Title = req.Title
	// TODO 暂时不记录摘要
	// approved.Summary = req.Summary
	approved.Summary = ""
	approved.Status = req.Status
	approved.Description = req.Description
	approved.Content = hptypes.DecodeToMap(req.Content)
	approved.ClubID = bson.ObjectIdHex(req.ClubID)
	approved.PusherID = bson.ObjectIdHex(req.CreatorID)

	if err := ApprovedModel.Create(approved); err != nil {
		return a.Error(ctx).BadRequest(err.Error())
	}

	for i, v := range approvedUsers {
		status := 0

		if i == 0 {
			status = 1
		}
		flows = append(flows, models.ApprovedFlow{
			Kind:       "approved",
			HandlerID:  bson.ObjectIdHex(v),
			Status:     status,
			Sort:       i,
			ApprovedID: approved.ID,
		})
	}

	for i, v := range req.CopyUsers {
		flows = append(flows, models.ApprovedFlow{
			Kind:       "copy",
			HandlerID:  bson.ObjectIdHex(v),
			Status:     0,
			Sort:       i,
			ApprovedID: approved.ID,
		})
	}

	for _, f := range flows {
		// fmt.Println(f)
		ApprovedFlowModel.Create(&f)
	}
	// ApprovedModel.Create()
	rsp.OK = true
	return nil
}

func (self *ApprovedHandler) ListActivity(ctx context.Context, req *iunite_club_srv_core.ListActivityRequest, rsp *iunite_club_srv_core.ListResponse) error {

	ApprovedModel := self.Model(ctx, "Approved")

	condition := bson.M{
		"club_id": req.ClubID, // 当前社团的
		"content.is_hidden": bson.M{
			"$ne": true,
		},
	}

	switch req.Kind {
	case "all":
		// 所有已经发布的活动
		condition["content.is_publish"] = true

	case "other":
		fallthrough
	case "mine":
		condition["pusher_id"] = req.UserID
		condition["content.is_publish"] = true
		condition["status"] = "pass"
	case "pending":
		condition["pusher_id"] = req.UserID
		condition["content.is_publish"] = false
		condition["status"] = "pass"
	}

	fmt.Println(condition)

	approveds := make([]models.Approved, 0)

	ApprovedModel.
		Where(condition).
		Populate("Flows", "Flows.Handler", "Flows.Handler.Profile", "Pusher", "Pusher.Profile", "Department").
		Aggregate([]bson.M{
			{
				"$lookup": bson.M{
					"as":   "content.pictureObjects",
					"from": "file",
					"let":  bson.M{"pictures": "$content.pictures"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$pictures"},
										"then": bson.M{"$in": []string{"$_id", "$$pictures"}},
										"else": bson.M{"$eq": []string{"$_id", "$$pictures"}},
									},
								},
							},
						},
					},
				},
			},
			{
				"$lookup": bson.M{
					"as":   "content.attachObjects",
					"from": "file",
					"let":  bson.M{"attachs": "$content.attachs"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$attachs"},
										"then": bson.M{"$in": []string{"$_id", "$$attachs"}},
										"else": bson.M{"$eq": []string{"$_id", "$$attachs"}},
									},
								},
							},
						},
					},
				},
			},
		}).
		FindAll(&approveds)

	rsp.Total = int64(ApprovedModel.Where(condition).Count())
	listSize := len(approveds)

	aps := make([]*iunite_club_srv_core.ApprovedPB, 0, listSize)
	for _, v := range approveds {
		aps = append(aps, v.ToPB())
	}
	rsp.Approveds = aps
	return nil

}

func (self *ApprovedHandler) PublishActivity(ctx context.Context, req *iunite_club_srv_core.PublishActivityRequest, rsp *iunite_club_srv_core.Response) error {
	ApprovedModel := self.Model(ctx, "Approved")

	condition := bson.M{
		"content._id": req.ID,
	}

	err := ApprovedModel.Update(condition, bson.M{
		"$set": bson.M{
			"content.is_publish": true,
		},
	})

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (self *ApprovedHandler) DismissActivity(ctx context.Context, req *iunite_club_srv_core.DismissActivityRequest, rsp *iunite_club_srv_core.Response) error {
	ApprovedModel := self.Model(ctx, "Approved")

	condition := bson.M{
		"content._id": req.ID,
	}

	err := ApprovedModel.Update(condition, bson.M{
		"$set": bson.M{
			"content.is_hidden": true,
		},
	})

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (self *ApprovedHandler) GetPendingApprovedCountByUserID(ctx context.Context, req *iunite_club_srv_core.GetPendingApprovedCountRequest, rsp *iunite_club_srv_core.PendingApprovedCountResponse) error {
	ApprovedModel := self.Model(ctx, "Approved")

	flowElemMatch := bson.M{
		"handler_id": req.UserID,
		"kind":       "approved",
		"status":     1,
	}

	condition := bson.M{
		"flows": bson.M{
			"$elemMatch": flowElemMatch,
		},
	}
	fmt.Println(condition, "condition")
	type CountScan struct {
		Count int64         `bson:"count"`
		ID    bson.ObjectId `bson:"_id"`
	}

	scan := make([]CountScan, 0)
	// scan := make(map[string]interface{})

	err := ApprovedModel.
		Aggregate([]bson.M{}).
		Populate("Flows").
		Where(bson.M{
			"$or": []bson.M{
				{
					"flows": bson.M{
						"$elemMatch": bson.M{
							"handler_id": bson.ObjectIdHex(req.UserID),
							"kind":       "approved",
							"status":     1,
						},
					},
				},
				{
					"flows": bson.M{
						"$elemMatch": bson.M{
							"handler_id": req.UserID,
							"kind":       "copy",
							"status":     0,
						},
					},
				},
			},
		}).
		Pipe(
			bson.M{
				"$group": bson.M{
					"_id": "$club_id",
					"count": bson.M{
						"$sum": 1,
					},
				},
			},
		).
		All(&scan)

	if err != nil {
		fmt.Println(err.Error())
		return self.Error(ctx).BadRequest(err.Error())
	}

	fmt.Println(scan, "scan")

	counts := make([]*iunite_club_srv_core.ApprovedCountObject, 0)
	for _, v := range scan {
		counts = append(counts, &iunite_club_srv_core.ApprovedCountObject{
			Count:  v.Count,
			ClubID: v.ID.Hex(),
		})
	}

	rsp.Counts = counts

	return nil
}
