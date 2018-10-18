package handler

import (
	"context"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"

	"github.com/iron-kit/monger"

	"github.com/iron-kit/go-ironic"

	pb "iunite.club/services/storage/proto/cloud"
)

type Cloud struct {
	ironic.BaseHandler
	conn monger.Connection
}

func (self *Cloud) model(ctx context.Context, field string) monger.Model {
	if self.conn == nil {
		conn, err := ironic.MongerConnectionFromContext(ctx)
		if err != nil {
			panic(err.Error())
		}
		self.conn = conn
	}

	return self.conn.M(field)
}

func (self *Cloud) FindFile(ctx context.Context, req *pb.ByFileID, rsp *pb.FileResponse) error {
	FileModel := self.model(ctx, "File")
	file := new(models.File)
	if err := FileModel.FindByID(bson.ObjectIdHex(req.ID), file); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.File = file.ToPB()
	return nil
}

func (self *Cloud) DeleteOne(ctx context.Context, req *pb.ByID, rsp *pb.IsOK) error {
	CloudModel := self.model(ctx, "Cloud")

	if err := CloudModel.Delete(bson.M{"_id": req.ID}); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (self *Cloud) UpdatePermission(ctx context.Context, req *pb.WithDepartmentsByFileID, rsp *pb.IsOK) error {
	CloudModel := self.model(ctx, "Cloud")
	deptObjectIds := make([]bson.ObjectId, 0)

	foundCloud := new(models.Cloud)

	for _, v := range req.Departments {
		deptObjectIds = append(deptObjectIds, bson.ObjectIdHex(v))
	}
	CloudModel.Where(bson.M{
		"_id":      bson.ObjectIdHex(req.ID),
		"owner_id": bson.ObjectIdHex(req.UserID),
	}).FindOne(foundCloud)

	if foundCloud.IsEmpty() {
		return nil
	}

	if len(deptObjectIds) > 0 {
		foundCloud.EnabledToAll = false
	} else {
		foundCloud.EnabledToAll = true
	}
	foundCloud.DepartmentIDS = deptObjectIds
	if err := CloudModel.Update(bson.M{"_id": bson.ObjectIdHex(req.ID)}, foundCloud); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}

	rsp.OK = true
	return nil
}

func (self *Cloud) CreateItem(ctx context.Context, req *pb.WithItemBundle, rsp *pb.IsCreateOK) error {
	CloudModel := self.model(ctx, "Cloud")
	condition := bson.M{
		"original_name": req.Name,
		"club_id":       bson.ObjectIdHex(req.ClubID),
		// "parent_id":     req.ParentID,
	}
	if req.ParentID != "" {
		condition["parent_id"] = bson.ObjectIdHex(req.ParentID)
	}
	// 重命名同名文件
	total := CloudModel.Where(condition).Count()
	name := req.Name
	if total > 0 {
		s := strings.SplitN(name, ".", 2)
		if len(s) > 1 {
			name = fmt.Sprintf("%s(%d).%s", s[0], total, s[1])
		} else {
			name = fmt.Sprintf("%s(%d)", s[0], total)
		}
	}

	// 创建记录
	item := models.Cloud{
		Name:         name,
		OriginalName: req.Name,
		OwnerID:      bson.ObjectIdHex(req.UserID),
		// FileID: bson.ObjectIdHex(req.FileID),
		EnabledToAll: true,
		ClubID:       bson.ObjectIdHex(req.ClubID),
		// ParentID:     bson.ObjectIdHex(req.ParentID),
		Kind: int(req.Kind),
	}

	if req.ParentID != "" {
		item.ParentID = bson.ObjectIdHex(req.ParentID)
	}

	if req.FileID != "" {
		item.FileID = bson.ObjectIdHex(req.FileID)
	}

	if len(req.Departments) > 0 {
		// item.EnabledToAll = false
		item.DepartmentIDS = make([]bson.ObjectId, 0)
		c := 0
		for _, v := range req.Departments {
			if bson.IsObjectIdHex(v) {
				item.DepartmentIDS = append(item.DepartmentIDS, bson.ObjectIdHex(v))
				c++
			}
		}

		if c > 0 {
			item.EnabledToAll = false
		}
	}

	if err := CloudModel.Create(&item); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}
	rsp.OK = true
	rsp.CloudItem = item.ToPB()
	return nil
}

func (self *Cloud) List(ctx context.Context, req *pb.ByParentAndClubID, rsp *pb.ListResponse) error {
	// panic("not implemented")

	CloudModel := self.model(ctx, "Cloud")
	ClubProfile := self.model(ctx, "UserClubProfile")
	fmt.Println(req.UserID)
	fmt.Println(req.ClubID)
	userID := bson.ObjectIdHex(req.UserID)
	clubID := bson.ObjectIdHex(req.ClubID)

	clubProfile := new(models.UserClubProfile)
	ClubProfile.Where(bson.M{
		"user_id":         userID,
		"organization_id": clubID,
	}).FindOne(clubProfile)

	if clubProfile.IsEmpty() {
		return self.Error(ctx).BadRequest("not found user %s's club profile", req.UserID)
	}

	deptID := func() interface{} {
		if len(clubProfile.DepartmentID) > 0 {
			return clubProfile.DepartmentID
		}

		return ""
	}()

	condition := bson.M{
		"club_id": clubID,
		"$or": []bson.M{
			// 或对文件的拥有者可见
			{
				"owner_id": userID,
			},
			// 或对所有人可见
			{
				"enabled_to_all": true,
			},
			// 或对归属部门内或归属部门的下级部门的人可见
			{
				"department_ids": deptID,
			},
			{
				"departments": bson.M{
					"$elemMatch": bson.M{
						"pathindexs": deptID,
					},
				},
			},
		},
	}

	if req.ParentID != "" {
		condition["parent_id"] = bson.ObjectIdHex(req.ParentID)
	} else {
		condition["parent_id"] = bson.M{
			"$exists": false,
		}
	}

	items := make([]models.Cloud, 0)
	query := CloudModel.
		Where(condition).
		Populate("File", "Owner", "Owner.Profile").
		Aggregate([]bson.M{
			{
				"$lookup": bson.M{
					"as":   "departments",
					"from": "organization",
					"let":  bson.M{"deptids": "$department_ids"},
					"pipeline": []bson.M{
						{
							"$match": bson.M{
								"$expr": bson.M{
									"$cond": bson.M{
										"if":   bson.M{"$isArray": "$$deptids"},
										"then": bson.M{"$in": []string{"$_id", "$$deptids"}},
										"else": bson.M{"$eq": []string{"$_id", "$$deptids"}},
									},
								},
							},
						},
						{
							"$graphLookup": bson.M{
								"from":             "organization",
								"startWith":        "$parent_id",
								"connectFromField": "parent_id",
								"connectToField":   "_id",
								"as":               "pathindexs",
							},
						},
					},
				},
			},
		})

	if err := query.Pipe().All(&items); err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}
	resItems := make([]*pb.CloudPB, 0)
	for _, v := range items {
		resItems = append(resItems, v.ToPB())
	}
	rsp.Items = resItems
	rsp.Total = int64(query.Count())
	return nil
}

func (self *Cloud) GetDetails(ctx context.Context, req *pb.ByID, rsp *pb.ItemResponse) error {
	// panic("not implemented")
	CloudModel := self.model(ctx, "Cloud")
	cloud := new(models.Cloud)
	// if err := CloudModel.FindByID(bson.ObjectIdHex(req.ID), cloud); err != nil {
	// 	return self.Error(ctx).BadRequest(err.Error())
	// }
	err := CloudModel.
		Where(bson.M{"_id": bson.ObjectIdHex(req.ID)}).
		Populate("File").
		FindOne(cloud)

	if err != nil {
		return self.Error(ctx).BadRequest(err.Error())
	}
	rsp.Item = cloud.ToPB()
	return nil

}
