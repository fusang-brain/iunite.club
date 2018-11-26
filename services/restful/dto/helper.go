package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	conversationPB "iunite.club/services/core/proto/conversation"

	"github.com/iron-kit/go-ironic/protobuf/hptypes"
	"github.com/iron-kit/go-ironic/utils"
	"gopkg.in/mgo.v2/bson"
	announcePB "iunite.club/services/core/proto/announce"
	approvedPB "iunite.club/services/core/proto/approved"
	contactPB "iunite.club/services/core/proto/contacts"
	recruitmentPB "iunite.club/services/core/proto/recruitment"
	orgPB "iunite.club/services/organization/proto"
	schoolPB "iunite.club/services/organization/proto/school"
	storagePB "iunite.club/services/storage/proto"
	cloudPB "iunite.club/services/storage/proto/cloud"
	userPB "iunite.club/services/user/proto"
)

func PBToFriendAccept(pb *contactPB.FriendAccept) *FriendAccept {
	if pb == nil {
		return nil
	}

	result := &FriendAccept{
		ID:            pb.ID,
		SenderRefer:   pb.SenderID,
		ReceiverRefer: pb.ReceiverID,
		Kind:          int(pb.Kind),
		Body:          pb.Body,
		GroupID:       pb.GroupID,
		State:         int(pb.State),
	}

	return result
}

func PBToConversationNotice(pb *conversationPB.NoticePB) *ConversationNotice {
	if pb == nil {
		return nil
	}
	result := &ConversationNotice{
		ID:                pb.ID,
		ConversationRefer: pb.ConversationID,
		Title:             pb.Title,
		Body:              pb.Body,
	}

	return result
}

func PBToConversation(pb *conversationPB.ConversationPB) *Conversation {
	if pb == nil {
		return nil
	}
	result := &Conversation{
		ID:              pb.ID,
		Kind:            pb.Kind,
		Name:            pb.Name,
		Avatar:          pb.Avatar,
		IsStartValidate: pb.IsStartValidate,
		IsTop:           pb.IsTop,
	}

	if len(pb.Members) > 0 {
		members := make([]ConversationUser, 0, len(pb.Members))

		for _, val := range pb.Members {
			members = append(members, ConversationUser{

				UserRefer:         val.UserID,
				ConversationRefer: pb.ID,
				Nickname:          val.Nickname,
				IsTop:             val.IsTop,
			})
		}

		result.Users = members
	}

	return result
}

func PBToAnnounce(pb *announcePB.AnnouncePB) *Announce {
	result := &Announce{
		ID:        pb.ID,
		CreatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Name:      pb.Name,
		Body:      pb.Body,
		Kind:      pb.Kind,
		Org:       pb.ClubID,
	}

	if pb.Options != nil {
		options := hptypes.DecodeToMap(pb.Options)
		if reminderTime, ok := options["ReminderTime"]; ok {
			t := reminderTime.(time.Time)

			result.ReminderTime = utils.Time2MicroUnix(t)
		}

		if reminderTime, ok := options["StartTime"]; ok {
			t := reminderTime.(time.Time)
			result.ReminderTime = utils.Time2MicroUnix(t)
		}
	}

	return result
}

func PBToRecruitmentFormRecords(pb *recruitmentPB.RecruitmentFormRecord) *RecruitmentFormRecords {
	result := &RecruitmentFormRecords{
		ID:              pb.ID,
		Mobile:          pb.Mobile,
		Name:            pb.Name,
		Major:           pb.Major,
		Age:             int(pb.Age),
		SchoolStudentID: pb.SchoolStudentID,
		// AcceptDepartment: pb.DepartmentID
		DepartmentRefer: pb.DepartmentID,
		RecordID:        pb.RecordID,
		Status:          int(pb.Status),
	}

	if len(pb.Answers) > 0 {
		answers := make([]RecruitmentFormAnswer, 0, len(pb.Answers))
		for _, ans := range pb.Answers {
			answers = append(answers, RecruitmentFormAnswer{
				ID:        ans.ID,
				Answer:    ans.Answer,
				ItemRefer: ans.FormID,
			})
		}
		result.Answers = answers
	}

	return result
}

func PBToRecruitmentFormItem(pb *recruitmentPB.RecruitmentFormField) *RecruitmentFormItem {
	if pb == nil {
		return nil
	}

	result := &RecruitmentFormItem{
		ID:      pb.ID,
		Key:     pb.Key,
		Subject: pb.Subject,
		// Options: hptypes.DecodeToMap(pb.Options),
		Kind: pb.Kind,
		// Organization: pb.
	}

	options := hptypes.DecodeToMap(pb.Options)

	if opts, ok := options["Opts"]; ok {
		if optsb, err := json.Marshal(opts); err == nil {
			result.Options = string(optsb)

			if result.Options == "null" {
				result.Options = ""
			}
		}
	} else {
		result.Options = ""
	}

	return result
}

func PBToRecruitmentForm(pb *recruitmentPB.RecruitmentForm) *RecruitmentForm {
	if pb == nil {
		return nil
	}

	result := &RecruitmentForm{
		ID:        pb.ID,
		CreatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		FormName:  pb.Name,
	}

	if len(pb.Fields) > 0 {
		items := make([]*RecruitmentFormItem, 0)
		for _, v := range pb.Fields {
			items = append(items, PBToRecruitmentFormItem(v))
		}
		result.Items = items
	}

	return result
}

func PBToRecruitmentRecord(pb *recruitmentPB.RecruitmentRecord) *RecruitmentRecord {
	if pb == nil {
		return nil
	}

	result := &RecruitmentRecord{
		ID:           pb.ID,
		CreatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Organization: pb.ClubID,
		CreateUser:   pb.CreateUserID,
		HasStart:     pb.HasStart,
		HasEnd:       pb.HasEnd,
	}

	if pb.Form != nil {
		result.RecruitmentForm = PBToRecruitmentForm(pb.Form)
	}

	return result
}

func PBToCloudDisk(pb *cloudPB.CloudPB) *CloudDisk {
	if pb == nil {
		return nil
	}
	res := &CloudDisk{
		ID:           pb.ID,
		CreatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Organization: pb.ClubID,
		Name:         pb.Name,
		OriginalName: pb.OriginalName,
		ParentID:     pb.ParentID,
		Kind:         int(pb.Kind),
		OwnerID:      pb.OwnerID,
		FileID:       pb.FileID,
		EnabledToAll: pb.EnabledToAll,
		// Departments: pb.DepartmentIDS,
		// Users:
	}

	if pb.File != nil {
		res.File = PBToFile(pb.File)
	}

	if pb.Owner != nil {
		res.Owner = PBToUser(pb.Owner)
	}

	return res
}

func PBToUser(pb *userPB.User) *User {
	if pb == nil {
		return &User{}
	}
	mobile := ""
	email := ""
	if pb.Profile != nil {
		mobile = pb.Profile.Mobile
		email = pb.Profile.Email
	}
	return &User{
		ID:        pb.ID,
		CreatedAt: utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt: utils.ISOTime2MicroUnix(pb.UpdatedAt),
		IsTeacher: false,
		IsAdmin:   false,
		Username:  pb.Username,
		Mobile:    mobile,
		AreaCode:  "+86",
		Enabled:   pb.Enabled,
		Profile:   PBToProfile(pb.Profile),
		Email:     email,
	}
}

func PBToProfile(pb *userPB.Profile) *Profile {

	if pb == nil {
		return &Profile{}
	}
	updatedAt := hptypes.Timestamp(pb.UpdatedAt)
	createdAt := hptypes.Timestamp(pb.CreatedAt)
	birthday := hptypes.Timestamp(pb.Birthday)
	return &Profile{
		ID:        pb.ID,
		CreatedAt: utils.Time2MicroUnix(createdAt),
		// CreatedAt: utils.ISOTime2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:        utils.Time2MicroUnix(updatedAt),
		Avatar:           pb.Avatar,
		Gender:           pb.Gender,
		Birthday:         utils.Time2MicroUnix(birthday),
		Nickname:         pb.Nickname,
		UserID:           pb.UserID,
		FirstName:        pb.Firstname,
		LastName:         pb.Lastname,
		UserNO:           "-",
		SchoolClass:      pb.SchoolClass,
		AdvisorMobile:    pb.AdvisorMobile,
		AdvisorName:      pb.AdvisorName,
		StudentID:        pb.StudentID,
		RoomNumber:       pb.RoomNumber,
		Major:            pb.Major,
		SchoolDepartment: pb.SchoolDepartment,
	}
}

func PBToOrganization(pb *orgPB.Organization) *Organization {
	if pb == nil {
		return &Organization{}
	}
	return &Organization{
		ID:          pb.ID,
		Name:        pb.Name,
		SlugName:    pb.Slug,
		Description: pb.Description,
		Logo:        pb.ClubProfile.Logo,
		Scale:       int(pb.ClubProfile.Scale),
		SchoolRefer: pb.SchoolID,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
	}
}

func PBToSchool(pb *schoolPB.School) *School {
	if pb == nil {
		return new(School)
	}
	return &School{
		ID:          pb.ID,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
		Name:        pb.Name,
		SlugName:    pb.SlugName,
		SchoolCode:  pb.SchoolCode,
		Description: pb.Description,
	}
}

func PBToDepartment(pb *orgPB.Organization) *Department {
	if pb == nil {
		return new(Department)
	}
	return &Department{
		ID:          pb.ID,
		SlugName:    pb.Slug,
		Name:        pb.Name,
		CreatedAt:   utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt:   utils.ISOTime2MicroUnix(pb.UpdatedAt),
		ParentID:    pb.ParentID,
		Org:         pb.ClubID,
		Description: pb.Description,
	}
}

func PBToJob(pb *orgPB.Job) *Job {
	if pb == nil {
		return new(Job)
	}

	return &Job{
		ID:        pb.ID,
		Name:      pb.Name,
		SlugName:  pb.Slug,
		CreatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt: utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Org:       pb.ClubID,
		// CreatedAt: utils.ISOTime2MicroUnix()
		// Org: pb.
	}
}

func PBToOrganizationUser(pb *orgPB.UserClubProfile) *OrganizationUser {
	if pb == nil || pb.ID == "" {
		return nil
	}
	rs := &OrganizationUser{
		ID:        pb.ID,
		CreatedAt: utils.ISOTime2MicroUnix(pb.CreatedAt),
		UpdatedAt: utils.ISOTime2MicroUnix(pb.UpdatedAt),
		// Kind:             1,
		// AcceptState:      1,
		// State:            pb.State,
		UserInfo:         PBToUser(pb.User),
		OrganizationInfo: PBToOrganization(pb.Organization),
		IsCreator:        pb.IsCreator,
		IsSuperManager:   pb.IsMaster,
		JoinTime:         utils.ISOTime2MicroUnix(pb.JoinTime),
		LeaveTime:        utils.ISOTime2MicroUnix(pb.LeaveTime),
		DepartmentID:     pb.DepartmentID,
		JobID:            pb.JobID,
	}
	if pb.Department != nil {
		// rs.Depart
		rs.Department = PBToDepartment(pb.Department)
	}
	if pb.Job != nil {
		rs.Job = PBToJob(pb.Job)
	}
	if pb.State == 1 || pb.State == 2 {
		rs.Kind = 1
		rs.AcceptState = 1
		rs.State = int(pb.State)
	}
	if pb.State == 3 {
		rs.Kind = 0
		rs.AcceptState = 0
		rs.State = 0
	}
	if pb.State == 4 {
		rs.Kind = 0
		rs.AcceptState = 2
		rs.State = 0
	}
	return rs
}

func PBToApprovedProcess(pb *approvedPB.ApprovedFlowPB) *ApprovedProcess {
	if pb == nil {
		return new(ApprovedProcess)
	}

	ap := &ApprovedProcess{
		ID:          pb.ID,
		CreatedAt:   utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:   utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		ProcessSort: int(pb.Sort),
		ProcessType: pb.Kind,
		Options:     pb.Options,
		ApprovedID:  pb.ApprovedID,
		HandlerID:   pb.HandlerID,
		Archived:    false,
		Status:      int(pb.Status),
	}

	if pb.Handler != nil {
		ap.Handler = PBToUser(pb.Handler)
	}

	return ap
}

func GetApprovedContent(pb *approvedPB.ApprovedPB) interface{} {
	if pb == nil {
		return nil
	}
	content := hptypes.DecodeToMap(pb.Content)
	switch pb.Kind {
	case "borrow":
		borrow := new(GoodsBorrow)
		if pb.Pusher != nil {
			borrow.Applicant = PBToUser(pb.Pusher)
			borrow.ApplicantID = pb.Pusher.ID
		}

		for k, val := range content {
			switch k {
			case "goods":
				goods := make([]interface{}, 0)
				goodItems := make([]*GoodsBorrowItem, 0)
				if v, ok := val.([]interface{}); ok {
					goods = v
				}
				for _, g := range goods {
					item := func() *GoodsBorrowItem {
						bitem := new(GoodsBorrowItem)
						gi := g.(map[string]interface{})
						for gk, gv := range gi {
							if gk == "Name" {
								bitem.Name = gv.(string)
							}

							if gk == "Count" {
								c := fmt.Sprintf("%v", gv)
								ci, _ := strconv.Atoi(c)
								bitem.Count = ci
							}
						}

						return bitem
					}()

					goodItems = append(goodItems, item)
				}
				borrow.Goods = goodItems
			case "description":
				borrow.Description = val.(string)
			case "subject":
				borrow.Subject = val.(string)
			case "start_time":
				borrow.StartTime = utils.Time2MicroUnix(val.(time.Time))
			case "end_time":
				borrow.EndTime = utils.Time2MicroUnix(val.(time.Time))
			case "created_at":
				borrow.CreatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "updated_at":
				borrow.UpdatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "_id":
				id := val.(bson.ObjectId)
				borrow.ID = id.Hex()
			case "pictureObjects", "attachObjects":
				pics := make([]interface{}, 0)
				files := make([]*File, 0)
				if v, ok := val.([]interface{}); ok {
					pics = v
				}

				for _, vi := range pics {
					v := vi.(map[string]interface{})
					file := func() *File {
						f := new(File)
						for k, pic := range v {
							if k == "original_filename" {
								f.OriginalFilename = pic.(string)
							}

							if k == "_id" {
								id := pic.(bson.ObjectId)
								f.ID = id.Hex()
								f.Filename = f.ID
							}
							if k == "path" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path
							}
							if k == "ext" {
								f.Ext = pic.(string)
							}
							if k == "file_key" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path

							}
							if k == "size" {
								s, _ := strconv.Atoi(pic.(string))
								f.Size = int64(s)
							}

						}
						return f
					}()

					files = append(files, file)
				}
				// activity.Pictures = files

				if k == "pictureObjects" {
					borrow.Pictures = files
				} else {
					borrow.Attach = files
				}
			}
		}

		return borrow
	case "funding":
		funding := new(Funding)
		if pb.Pusher != nil {
			funding.Applicant = PBToUser(pb.Pusher)
			funding.ApplicantID = pb.Pusher.ID
		}

		for k, val := range content {
			switch k {
			case "apply_purpose":
				funding.ApplyPurpose = val.(string)
			case "amount_apply_fee":
				funding.AmountApplyFee = int64(val.(float64))
			case "created_at":
				funding.CreatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "updated_at":
				funding.UpdatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "_id":
				id := val.(bson.ObjectId)
				funding.ID = id.Hex()
			case "pictureObjects", "attachObjects":
				pics := make([]interface{}, 0)
				files := make([]*File, 0)
				if v, ok := val.([]interface{}); ok {
					pics = v
				}

				for _, vi := range pics {
					v := vi.(map[string]interface{})
					file := func() *File {
						f := new(File)
						for k, pic := range v {
							if k == "original_filename" {
								f.OriginalFilename = pic.(string)
							}

							if k == "_id" {
								id := pic.(bson.ObjectId)
								f.ID = id.Hex()
								f.Filename = f.ID
							}
							if k == "path" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path
							}
							if k == "ext" {
								f.Ext = pic.(string)
							}
							if k == "file_key" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path

							}
							if k == "size" {
								s, _ := strconv.Atoi(pic.(string))
								f.Size = int64(s)
							}

						}
						return f
					}()

					files = append(files, file)
				}
				// activity.Pictures = files

				if k == "pictureObjects" {
					funding.Pictures = files
				} else {
					funding.Attach = files
				}
			}
		}
		return funding
	case "activity":
		activity := new(Activity)
		if pb.Pusher != nil {
			activity.Applicant = PBToUser(pb.Pusher)
			activity.ApplicantID = pb.Pusher.ID
		}

		for k, val := range content {
			switch k {
			case "subject":
				activity.Subject = val.(string)
			case "location":
				activity.Location = val.(string)
			case "start_time":
				activity.StartTime = utils.Time2MicroUnix(val.(time.Time))
			case "end_time":
				activity.EndTime = utils.Time2MicroUnix(val.(time.Time))
			case "created_at":
				activity.CreatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "updated_at":
				activity.UpdatedAt = utils.Time2MicroUnix(val.(time.Time))
			case "_id":
				id := val.(bson.ObjectId)
				activity.ID = id.Hex()
			case "amount_funding":
				activity.AmountFunding = int64(val.(float64))
			case "is_publish":
				activity.IsPublish = val.(bool)
			case "participants_total":
				activity.ParticipantsTotal = int(val.(float64))
			case "pictureObjects", "attachObjects":
				pics := make([]interface{}, 0)
				files := make([]*File, 0)
				if v, ok := val.([]interface{}); ok {
					pics = v
				}

				for _, vi := range pics {
					v := vi.(map[string]interface{})
					file := func() *File {
						f := new(File)
						for k, pic := range v {
							if k == "original_filename" {
								f.OriginalFilename = pic.(string)
							}

							if k == "_id" {
								id := pic.(bson.ObjectId)
								f.ID = id.Hex()
								f.Filename = f.ID
							}
							if k == "path" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path
							}
							if k == "ext" {
								f.Ext = pic.(string)
							}
							if k == "file_key" {
								path := pic.(string)
								f.Path = path
								f.AbstractPath = path

							}
							if k == "size" {
								s, _ := strconv.Atoi(pic.(string))
								f.Size = int64(s)
							}

						}
						return f
					}()

					files = append(files, file)
				}
				// activity.Pictures = files

				if k == "pictureObjects" {
					activity.Pictures = files
				} else {
					activity.Attach = files
				}
			}
		}
		return activity
	default:
		return nil
	}
}

// func GetApprovedContent(pb *approvedPB.ApprovedPB) interface{} {
// 	if pb == nil {
// 		return nil
// 	}

// 	var res interface{}
// 	content := hptypes.DecodeToMap(pb.Content)
// 	switch pb.Kind {
// 	case "activity":

// 		activity := &Activity{
// 			ID:        content["ID"].(string),
// 			CreatedAt: utils.ISOTime2MicroUnix(content["CreatedAt"].(string)),
// 			UpdatedAt: utils.ISOTime2MicroUnix(content["UpdatedAt"].(string)),
// 			Subject: content["subject"].(string),
// 			Location: content["location"].(string),
// 			ApplicantID: content["applicant_id"].(string),
// 			Applicant: content["applicant"]
// 		}
// 	case "borrow":
// 	case "funding":

// 	}
// 	// res := interface{}{}

// }

func PBToFile(pb *storagePB.FilePB) *File {
	return &File{
		ID: pb.ID,
		// CreatedAt:
		Filename:         pb.FileKey,
		Path:             pb.Path,
		Ext:              pb.Ext,
		AbstractPath:     pb.Path,
		Host:             pb.Host,
		OriginalFilename: pb.OriginalFilename,
		Size:             pb.Size,
	}
}

func PBToApprovedTask(pb *approvedPB.ApprovedPB) *ApprovedTask {
	if pb == nil {
		return new(ApprovedTask)
	}

	at := &ApprovedTask{
		ID:           pb.ID,
		CreatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.CreatedAt)),
		UpdatedAt:    utils.Time2MicroUnix(hptypes.Timestamp(pb.UpdatedAt)),
		Name:         pb.Title,
		Description:  pb.Description,
		ApprovedType: pb.Kind,
		Content:      pb.Description,
		Status:       pb.Status,
		Summary:      pb.Summary,
	}

	if len(pb.Flows) > 0 {
		flows := make([]ApprovedProcess, 0, len(pb.Flows))

		for _, v := range pb.Flows {
			f := PBToApprovedProcess(v)
			flows = append(flows, *f)
		}
		at.ApprovedProcesses = flows
	}

	// if pb.Pusher != nil

	return at
}

func PBToUserMetaData(pb *conversationPB.UserMetaData) *UserMetaData {
	return &UserMetaData{
		ID: pb.ID,
		RealName: pb.RealName,
		Avatar: pb.Avatar,
		Nickname: pb.Nickname,
		RemarkName: pb.RemarkName,
		GroupNickname: pb.GroupNickname,
		Email: pb.Email,
	}
}

func PBToConversatoinMetaData(pb *conversationPB.ConversationMetaData) *ConversationMetaData {
	if pb == nil {
		return new(ConversationMetaData)
	}

	res := &ConversationMetaData{
		UniteConversationID: pb.UniteConversationID,
		Kind: pb.Kind,
		ConversationName: pb.ConversationName,
		ConversationAvatar: pb.ConversationAvatar,
		// MemberMapper: pb.MemberMapper,
		TopMembers: pb.TopMembers,
		IsTop: pb.IsTop,
	}

	if pb.MemberMapper != nil {
		res.MemberMapper = make(map[string]*UserMetaData)
		for key, value := range pb.MemberMapper {
			res.MemberMapper[key] = PBToUserMetaData(value)
		}
	}

	return res
}