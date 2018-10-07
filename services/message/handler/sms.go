package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-log/log"
	"github.com/iron-kit/go-ironic"
	"github.com/iron-kit/go-ironic/utils"
	"gopkg.in/mgo.v2/bson"
	"iunite.club/models"
	smsPB "iunite.club/services/message/proto/sms"
	smsUtils "iunite.club/services/message/utils"
)

type SMSHandler struct {
	ironic.BaseHandler
}

func (s *SMSHandler) SendVerifyCode(ctx context.Context, in *smsPB.SendVerifyCodeRequest, resp *smsPB.SendResponse) error {
	// if in.Type == "Register" {
	// 	s.sendRegisterCode(ctx, in.Mobile, "")
	// }
	var body string
	var err error

	switch in.Type {
	case "Register":
		body, err = s.sendRegisterCode(ctx, in.Mobile, "")
	case "ForgetPassword":
		body, err = s.sendForgetPasswordCode(ctx, in.Mobile, "")
	default:
		err = errors.New("ErrorType")
		body = ""
	}

	log.Logf("send sms body %s", body)

	if err != nil {
		return s.Error(ctx).BadRequest(err.Error())
	}

	resp.OK = true
	resp.Message = body

	return nil
}

func (s *SMSHandler) ValidateMobileCode(ctx context.Context, in *smsPB.ValidateMobileCodeRequest, resp *smsPB.ValidateResponse) error {
	// 验证验证码是否是最新的未使用的验证码

	db, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		return s.Error(ctx).InternalServerError(err.Error())
	}

	ValdiateCodeModel := db.M("ValidateCode")
	verifyCode := models.ValidateCode{}

	ValdiateCodeModel.
		Where(bson.M{
			"mobile": in.Mobile,
			"code":   in.Code,
			"$or": []bson.M{
				{
					"usaged": false,
				},
				{
					"usaged": bson.M{
						"$exists": false,
					},
				},
			},
			// "usaged": false,
		}).FindOne(&verifyCode)

	fmt.Println(verifyCode)
	now := time.Now()

	if verifyCode.IsEmpty() {
		return s.Error(ctx).TemplateBadRequest("ErrorCode")
	}

	if verifyCode.ExpiredAt <= now.Unix() {
		return s.Error(ctx).TemplateBadRequest("CodeHasExpired")
	}
	verifyCode.Usaged = true
	ValdiateCodeModel.UpsertID(verifyCode.ID, verifyCode)
	resp.OK = true

	return nil
}

func (s *SMSHandler) sendVerifyCode(ctx context.Context, mobile, sign string) (string, error) {
	defaultSign := "【Unite】"
	smsTemplate := "%s您的验证码是: %d"
	enableSMSSend := false

	if sign == "" {
		sign = defaultSign
	}
	db, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		return "", err
	}

	ValidateCodeModel := db.M("ValidateCode")
	nowTime := time.Now()
	verifyCount := ValidateCodeModel.Count(bson.M{
		"mobile": mobile,
		// "usaged": bson.M{
		// 	"$or": []bson.M{
		// 		{
		// 			"$exists": false,
		// 		},
		// 		{
		// 			"$eq": false,
		// 		},
		// 	},
		// },
		"createdAt": bson.M{
			"$gte": nowTime.Add(-time.Second * 60),
		},
	})

	if verifyCount > 0 {
		return "", errors.New("Send too many times")
	}

	code := utils.GenerateRangeNum(100000, 999999)
	verifyCode := models.ValidateCode{
		Mobile:    mobile,
		Code:      strconv.Itoa(code),
		ExpiredAt: nowTime.Add(time.Minute * 15).Unix(),
	}

	ValidateCodeModel.Create(&verifyCode)

	msgBody := fmt.Sprintf(smsTemplate, sign, code)

	if !enableSMSSend {
		return msgBody, nil
	}

	if !smsUtils.SendMsg(mobile, msgBody) {
		return "", errors.New("短信发送失败")
	}

	return "", nil
}

func (s *SMSHandler) sendRegisterCode(ctx context.Context, mobile, sign string) (string, error) {

	db, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		return "", err
	}
	ProfileModel := db.M("Profile")

	foundUserCount := ProfileModel.Count(bson.M{
		"mobile": mobile,
	})

	if 0 != foundUserCount {
		return "", errors.New("UserHasRegisted")
	}

	return s.sendVerifyCode(ctx, mobile, sign)
}

func (s *SMSHandler) sendForgetPasswordCode(ctx context.Context, mobile string, sign string) (string, error) {
	db, err := ironic.MongerConnectionFromContext(ctx)
	if err != nil {
		return "", err
	}
	ProfileModel := db.M("Profile")

	foundUserCount := ProfileModel.Count(bson.M{
		"mobile": mobile,
	})

	if 0 == foundUserCount {
		return "", errors.New("UserNotExists")
	}

	return s.sendVerifyCode(ctx, mobile, sign)
}
