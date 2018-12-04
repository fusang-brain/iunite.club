package handler

import (
	"context"

	ironic "github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/monger"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	pb "iunite.club/services/core/proto/announce"
)

type Announce struct {
	ironic.BaseHandler
}

func (announce *Announce) model(ctx context.Context, name string) monger.Model {
	conn, err := ironic.MongerConnectionFromContext(ctx)

	if err != nil {
		panic(err.Error())
	}

	return conn.M(name)
}

// CreateInstructions 创建社长令
func (announce *Announce) CreateInstructions(ctx context.Context, req *pb.CreateInstructionsRequest, rsp *pb.CreatedResponse) error {
	// AnnounceModel := announce.model(ctx, "Announce")
	AnnounceModel := announce.model(ctx, "Announce")
	UserClubProfileModel := announce.model(ctx, "UserClubProfile")
	newAnnounce := &models.Announce{
		Name:   req.Name,
		Body:   req.Body,
		Kind:   models.KindAnnounceInstructions,
		ClubID: bson.ObjectIdHex(req.ClubID),
	}

	users := make([]models.User, 0)

	UserClubProfileModel.Where(bson.M{
		"organization_id": req.ClubID,
		"state": 1,
	}).FindAll(&users)


	for _, value := range users {
		announceReceiver := new(models.AnnounceReceiver)
		announceReceiver.HasRead = false
		announceReceiver.UserID = value.ID
		newAnnounce.Receivers = append(newAnnounce.Receivers, *announceReceiver)
	}

	if err := AnnounceModel.Create(newAnnounce); err != nil {
		return announce.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true
	rsp.Announce = newAnnounce.ToPB()
	return nil
}



func (announce *Announce) CreateTask(ctx context.Context, req *pb.CreateTaskRequest, rsp *pb.CreatedResponse) error {
	AnnounceModel := announce.model(ctx, "Announce")
	newTask := &models.Announce{
		Name:   req.Name,
		Body:   req.Body,
		ClubID: bson.ObjectIdHex(req.ClubID),
		Kind:   models.KindAnnounceTask,
		Options: map[string]interface{}{
			"StartTime": hptypes.Timestamp(req.StartTime),
			"EndTime":   hptypes.Timestamp(req.EndTime),
		},
	}

	if len(req.Users) > 0 {
		receivers := make([]models.AnnounceReceiver, 0, len(req.Users))

		for _, u := range req.Users {
			receivers = append(receivers, models.AnnounceReceiver{
				UserID:  bson.ObjectIdHex(u),
				HasRead: false,
			})
		}

		newTask.Receivers = receivers
	}

	if err := AnnounceModel.Create(newTask); err != nil {
		return announce.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	rsp.Announce = newTask.ToPB()
	return nil
}

func (announce *Announce) CreateReminder(ctx context.Context, req *pb.CreateReminderRequest, rsp *pb.CreatedResponse) error {
	AnnounceModel := announce.model(ctx, "Announce")
	newReminder := &models.Announce{
		Name:   req.Name,
		Body:   req.Body,
		ClubID: bson.ObjectIdHex(req.ClubID),
		Kind:   models.KindAnnounceReminder,
		Options: map[string]interface{}{
			"ReminderTime": hptypes.Timestamp(req.ReminderTime),
		},
	}

	if len(req.Users) > 0 {
		receivers := make([]models.AnnounceReceiver, 0, len(req.Users))

		for _, u := range req.Users {
			receivers = append(receivers, models.AnnounceReceiver{
				UserID:  bson.ObjectIdHex(u),
				HasRead: false,
			})
		}

		newReminder.Receivers = receivers
	}

	if err := AnnounceModel.Create(newReminder); err != nil {
		return announce.Error(ctx).InternalServerError(err.Error())
	}
	rsp.OK = true
	rsp.Announce = newReminder.ToPB()
	return nil
}

func (announce *Announce) GetAnnounces(ctx context.Context, req *pb.GetAnnouncesRequest, rsp *pb.AnnounceResponse) error {
	AnnounceModel := announce.model(ctx, "Announce")

	condition := bson.M{
		"club_id": req.ClubID,
		"kind":    req.Kind,
	}
	announces := make([]models.Announce, 0, int(req.Limit))
	//if req.Kind != models.KindAnnounceInstructions {
	//	condition["receivers.user_id"] = req.UserID
	//}
	condition["receivers.user_id"] = req.UserID
	query := AnnounceModel.Where(condition).Query()
	total := query.Query().Count()
	query.
		Query().
		Skip(int((req.Page - 1) * req.Limit)).
		Limit(int(req.Limit)).
		FindAll(&announces)

	rsp.Total = int32(total)
	pbAnnounces := make([]*pb.AnnouncePB, 0, len(announces))
	for _, v := range announces {
		pbAnnounces = append(pbAnnounces, v.ToPB())
	}

	rsp.Announces = pbAnnounces

	return nil
}

func (announce *Announce) GetUnreadCountByUserID(ctx context.Context, req *pb.ByUserID, rsp *pb.UnreadCountResponse) error {
	AnnounceModel := announce.model(ctx, "Announce")
	condition := bson.M{
		"club_id":            req.ClubID,
		"receivers.user_id":  req.UserID,
		"receivers.has_read": false,
	}

	count := AnnounceModel.Where(condition).Count()
	rsp.Count = int32(count)

	return nil
}

func (announce *Announce) MarkedOneToRead(ctx context.Context, req *pb.MarkedOneToReadRequest, rsp *pb.Response) error {
	AnnounceModel := announce.model(ctx, "Announce")

	if err := AnnounceModel.Update(bson.M{
		"_id":               req.ID,
		"receivers.user_id": req.UserID,
	}, bson.M{
		"$set": bson.M{"receivers.$.has_read": true},
	}); err != nil {
		return announce.Error(ctx).InternalServerError(err.Error())
	}

	rsp.OK = true

	return nil
}