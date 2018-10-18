package main

import (
	"net/http"

	"go.uber.org/ratelimit"

	restful "github.com/emicklei/go-restful"
	restfulApi "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"github.com/micro/go-micro/client"

	"github.com/micro/go-log"

	web "github.com/micro/go-web"
	"iunite.club/services/restful/handler"
)

func getWebService(path, doc string, consumes ...string) *restful.WebService {
	// consumes := restful.MIME_JSON
	// if len(consumes) == 0 {
	// 	consumes = restful.DefaultRequestContentType
	// }
	ws := new(restful.WebService)
	ws.Path(path).
		Doc(doc).
		// Consumes(restful.MIME_JSON, restful.MIME_OCTET).
		Produces(restful.MIME_JSON)
	if len(consumes) > 0 {
		ws.Consumes(consumes...)
	}

	return ws
}

func getWebContainer() *restful.Container {
	c := client.DefaultClient
	approvedHandler := handler.NewApprovedHandler(c)
	fileHandler := handler.NewFileHandler(c)
	schoolHandler := handler.NewSchoolHandler(c)
	authHandler := handler.NewAuthHandler(c)
	smsHandler := handler.NewSMSHandler(c)
	organizationHandler := handler.NewOrganizationHandler(c)
	departmentHandler := handler.NewDepartmentHandler(c)
	jobHandler := handler.NewJobService(c)
	userHandler := handler.NewUserHandler(c)
	activityHandler := handler.NewActivityHandler(c)
	fundingHandler := handler.NewFundingHandler(c)
	borrowHandler := handler.NewBorrowHandler(c)
	cloudHandler := handler.NewCloudService(c)
	otherHandler := new(handler.OtherHandler)

	rc := restful.NewContainer()

	// Test

	res := getWebService("/", "images")
	// res.Route(res.GET("/A/{id}").To(cloudHandler.ShowFile))
	res.Route(res.GET("/avatar/{id}").To(cloudHandler.ShowFile).Doc("获取头像"))
	res.Route(res.GET("/images/{id}").To(cloudHandler.ShowFile).Doc("显示图片"))

	// Base service
	base := getWebService("/v1", "base", restful.MIME_JSON)
	setRoute(base, authHandler.Login, "POST", "/login", "base", func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		type P struct {
			Mobile   string `json:"mobile" description:"用户手机"`
			Password string `json:"password" description:"密码"`
		}
		return rb.
			Doc("登录接口").Reads(P{})

		// Param(base.BodyParameter("password", "用户密码").DataType("string").DefaultValue("123456"))
	})
	setRoute(base, authHandler.Register, "POST", "register", "base")
	setRoute(base, otherHandler.GetUnreadCount, "GET", "/notify/unreadCount", "base")
	setRoute(base, otherHandler.GetAllCanUseRoles, "GET", "/role/getAllCanUseRoles", "base")
	// setRouteTag(base, []string{"base"})
	// base.Route(base.POST("/login").To(authHandler.Login)).Doc("登录")
	// base.Route(base.POST("/register").To(authHandler.Register)).Doc("注册")
	// base.Route(base.GET("/notify/unreadCount").To(otherHandler.GetUnreadCount)).Doc("Get unread message count")
	// base.Route(base.GET("/role/getAllCanUseRoles").To(otherHandler.GetAllCanUseRoles)).Doc("Get all can use roles")
	// base.Route(base.GET("/images/{id}").To(cloudHandler.ShowFile))

	// File service
	file := getWebService("/v1/file", "File service")
	setRoute(file, fileHandler.UploadSingleFile, "POST", "/uploadSingleFile", "file")
	setRoute(file, fileHandler.UploadMutipartFile, "POST", "/uploadMultipartFile", "file")
	// setRouteTag(file, []string{"file"})
	// file.Route(file.POST("/uploadSingleFile").To(fileHandler.UploadSingleFile)).Doc("上传单文件")
	// file.Route(file.POST("/uploadMultipartFile").To(fileHandler.UploadMutipartFile)).Doc("上传多文件")

	// School Serivce
	school := getWebService("/v1/school", "School service")
	setRoute(school, schoolHandler.SchoolList, "GET", "/allSchools", "school")
	setRoute(school, schoolHandler.SearchSchools, "GET", "/searchSchools", "school")
	setRoute(school, schoolHandler.Create, "POST", "/create", "school")
	// setRouteTag(school, []string{"school"})
	// school.Route(school.GET("/allSchools").To(schoolHandler.SchoolList)).Doc("获取学校列表")
	// school.Route(school.GET("/searchSchools").To(schoolHandler.SearchSchools)).Doc("搜索学校")
	// school.Route(school.POST("/create").To(schoolHandler.Create)).Doc("创建学校")

	// SMS Service
	sms := getWebService("/v1/sms", "SMS service")
	sms.Route(
		sms.
			POST("/sendVerifyCode").
			To(smsHandler.SendVerifyCode).
			Metadata(restfulApi.KeyOpenAPITags, []string{"sms"}).
			Doc("发送短信验证码"),
	)
	sms.Route(
		sms.
			POST("/validateSimpleCode").
			To(smsHandler.ValidateSimpleCode).
			Metadata(restfulApi.KeyOpenAPITags, []string{"sms"}).
			Doc("验证短信验证码"),
	)
	// setRouteTag(sms, []string{"sms"})

	// Organization Service
	// Organization Handler
	organization := getWebService("/v1/organization", "Organization service")
	organization.Route(
		organization.
			POST("/createOrganization").
			To(organizationHandler.CreateOrganization).
			Metadata(restfulApi.KeyOpenAPITags, []string{"organization"}).
			Doc("创建社团组织"),
	)
	setRoute(organization, organizationHandler.GetAllOrgByUserID, "GET", "/getAllOrgByUserID", "organization").Doc("通过用户ID获取所有组织")
	setRoute(organization, organizationHandler.GetAllOrgUsersByUserID, "GET", "/getAllOrgUsersByUserID", "organization").Doc("通过用户ID获取用户社团信息")
	setRoute(organization, organizationHandler.SearchHotOrganization, "GET", "/searchHotOrganization", "organization").Doc("搜索热门社团")
	setRoute(organization, organizationHandler.AcceptJoin, "POST", "/acceptJoin", "organization").Doc("同意社团申请")
	setRoute(organization, organizationHandler.AgreeJoin, "POST", "/agreeJoin", "organization")
	setRoute(organization, organizationHandler.RefuseJoin, "POST", "/refuseJoin", "organization")
	setRoute(organization, organizationHandler.FindRefusedAccept, "GET", "/findRefusedAccept", "organization")
	setRoute(organization, organizationHandler.GetDepartmentsAndUsers, "GET", "/getDepartmentDetails", "organization")
	setRoute(organization, organizationHandler.Info, "GET", "/info", "organization")
	setRoute(organization, organizationHandler.UploadLogo, "POST", "/uploadLogo", "organization")
	setRoute(organization, organizationHandler.UpdateOrganizationDescription, "POST", "/updateOrganizationDescription", "organization")
	setRoute(organization, organizationHandler.GetAllOrganizationBySchool, "GET", "/getAllOrganizationBySchool", "organization")
	setRoute(organization, organizationHandler.GetOrganizationDetails, "GET", "/getOrganizationDetails", "organization")
	setRoute(organization, organizationHandler.GetOrganizationUserInfoDetails, "GET", "/getOrganizationUserInfoDetails", "organization")
	setRoute(organization, organizationHandler.SelectOrganizations, "POST", "/selectOrganizations", "organization")
	// setRouteTag(organization, []string{"organization"})

	// User Service
	user := getWebService("/v1/user", "User Service")
	setRoute(user, userHandler.Info, "GET", "/info", "user")
	setRoute(user, userHandler.UpdateCurrentOrg, "POST", "/updateCurrentOrg", "user")
	setRoute(user, userHandler.ForgetPassword, "POST", "/forgetPassword", "user")
	setRoute(user, userHandler.AllUser, "GET", "/allUser", "user")
	setRoute(user, userHandler.GetCurrentOrganization, "GET", "/getCurrentOrganization", "user")
	setRoute(user, userHandler.GetAllMembers, "GET", "/getAllMembers", "user")
	setRoute(user, userHandler.CreateMember, "POST", "/createMember", "user")
	setRoute(user, userHandler.RemoveMemberFromOrg, "POST", "/removeMemberFromOrg", "user")
	setRoute(user, userHandler.UpdateMember, "POST", "/updateMember", "user")
	setRoute(user, userHandler.GetMemberDetails, "GET", "/getMemberDetails", "user")
	setRoute(user, userHandler.RemoveOrg, "POST", "/removeOrg", "user")
	setRoute(user, userHandler.UpdateUserInfo, "POST", "/updateUserInfo", "user")
	setRoute(user, userHandler.FlagMemberState, "POST", "/flagMemberState", "user")
	setRoute(user, userHandler.GetHotUsers, "GET", "/getHotUsers", "user")
	setRoute(user, userHandler.UploadAvatar, "POST", "/uploadAvatar", "user")
	setRoute(user, userHandler.ExportList, "GET", "/exportList", "user")
	setRoute(user, userHandler.DownloadExportTemplate, "GET", "/downloadExportTemplate", "user")
	setRoute(user, userHandler.UploadUserList, "POST", "/uploadUserList", "user")
	// setRouteTag(user, []string{"user"})

	job := getWebService("/v1/job", "Job Service")
	setRoute(job, jobHandler.CreateJob, "POST", "/createJob", "job")
	setRoute(job, jobHandler.GetUsersWithJob, "GET", "/getUsersWithJob", "job")
	setRoute(job, jobHandler.AllCanSelectedUsers, "GET", "/allCanSelectedUsers", "job")
	setRoute(job, jobHandler.AddUsersToJob, "POST", "/addUsersToJob", "job")
	setRoute(job, jobHandler.RemoveUsersFromJob, "POST", "/removeUsersFromJob", "job")
	setRoute(job, jobHandler.GetAllJobs, "GET", "/getAllJobs", "job")
	setRoute(job, jobHandler.Update, "POST", "/update", "job")
	setRoute(job, jobHandler.Remove, "POST", "/deleteJob", "job")
	// setRouteTag(job, []string{"job"})

	department := getWebService("/v1/department", "Department Service")
	setRoute(department, departmentHandler.AddDept, "POST", "/addDept", "department")
	setRoute(department, departmentHandler.GetDepartmentByOrg, "GET", "/getDepartmentByOrg", "department")
	setRoute(department, departmentHandler.GetAllDepartmentByOrg, "GET", "/getAllDepartmentByOrg", "department")
	setRoute(department, departmentHandler.SearchDepartment, "GET", "/searchDepartment", "department")
	setRoute(department, departmentHandler.AddUserToDepartment, "POST", "/addUserToDepartment", "department")
	setRoute(department, departmentHandler.RemoveUserFromDepartment, "POST", "/removeUserFromDepartment", "department")
	setRoute(department, departmentHandler.AllCanSelectedUsers, "GET", "/allCanSelectedUsers", "department")
	setRoute(department, departmentHandler.GetAllUsersWithDepartment, "GET", "/getAllUsersWithDepartment", "department")
	setRoute(department, departmentHandler.Update, "POST", "/update", "department")
	setRoute(department, departmentHandler.Remove, "POST", "/remove", "department")
	// setRouteTag(department, []string{"department"})

	activity := getWebService("/v1/activity", "Activity Service")
	setRoute(activity, activityHandler.CreateActivity, "POST", "/create", "activity")
	setRoute(activity, activityHandler.GetActivities, "GET", "/getActivities", "activity")
	setRoute(activity, activityHandler.Details, "GET", "/details", "activity")
	setRoute(activity, activityHandler.PublishActivity, "POST", "/publishActivity", "activity")
	setRoute(activity, activityHandler.DismissActivity, "POST", "/dismissActivity", "activity")

	funding := getWebService("/v1/funding", "Funding Service")
	setRoute(funding, fundingHandler.Create, "POST", "/create", "funding")

	borrow := getWebService("/v1/goodsBorrow", "Borrow Service")
	setRoute(borrow, borrowHandler.Create, "POST", "/create", "borrow")

	cloud := getWebService("/v1/cloud", "Cloud Service")
	setRoute(cloud, cloudHandler.DownloadFile, "GET", "/downloadFile", "cloud")
	setRoute(cloud, cloudHandler.ShowFile, "GET", "/stream/{id}", "cloud")
	setRoute(cloud, cloudHandler.List, "GET", "/list", "cloud")
	// setRoute(cloud, cloudHandler., "POST", "/deleteOne")
	setRoute(cloud, cloudHandler.CreateDIR, "POST", "/createDIR", "cloud")
	setRoute(cloud, cloudHandler.UploadFile, "POST", "/uploadFile", "cloud")
	setRoute(cloud, cloudHandler.Details, "GET", "/details", "cloud")
	setRoute(cloud, cloudHandler.UpdatePermission, "POST", "/updatePermission", "cloud")

	approved := getWebService("/v1/approved", "Approved service")
	setRoute(approved, approvedHandler.List, "GET", "/list", "approved")
	setRoute(approved, approvedHandler.Details, "GET", "/details", "approved")
	setRoute(approved, approvedHandler.ExecuteOne, "POST", "/executeOne", "approved")

	approvedV2 := getWebService("/v2/approved", "Approved v2 service")
	setRoute(approvedV2, approvedHandler.ListV2, "GET", "/list", "approved")
	setRoute(approvedV2, approvedHandler.ListByPusher, "GET", "/listByPusher", "approved")

	rc.Add(file)
	rc.Add(school)
	rc.Add(base)
	rc.Add(organization)
	rc.Add(sms)
	rc.Add(user)
	rc.Add(job)
	rc.Add(department)
	rc.Add(activity)
	rc.Add(funding)
	rc.Add(borrow)
	rc.Add(cloud)
	rc.Add(approved)
	rc.Add(approvedV2)
	rc.Add(res)

	config := restfulApi.Config{
		WebServices: rc.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject,
	}

	rc.Add(restfulApi.NewOpenAPIService(config))
	// rc.Se

	return rc
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Unite - Navo",
			Description: "Unite API 服务",
			Contact: &spec.ContactInfo{
				Name:  "alixe",
				Email: "alixe.z@foxmail.com",
				URL:   "http://blog.ironkit.xyz",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "Navo v1.0.0",
		},
	}
	swo.Tags = []spec.Tag{
		{
			TagProps: spec.TagProps{
				Name:        "base",
				Description: "基础服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "user",
				Description: "用户服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "organization",
				Description: "组织服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "file",
				Description: "文件服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "school",
				Description: "学校服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "sms",
				Description: "短信服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "job",
				Description: "组织职位服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "department",
				Description: "组织部门服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "activity",
				Description: "活动服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "funding",
				Description: "经费服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "borrow",
				Description: "物品借用服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "cloud",
				Description: "云盘服务",
			},
		},
		{
			TagProps: spec.TagProps{
				Name:        "approved",
				Description: "审批服务",
			},
		},
	}
}

// func setRouteTag(ws *restful.WebService, tag []string) {
// 	for _, v := range ws.Routes() {
// 		// v.Metadata(restfulApi.KeyOpenAPITags, tag)
// 		if v.Metadata == nil {
// 			v.Metadata = make(map[string]interface{})
// 		}

// 		v.Metadata[restfulApi.KeyOpenAPITags] = tag
// 	}
// }

type routeOptionFunc func(rb *restful.RouteBuilder) *restful.RouteBuilder

func setRoute(ws *restful.WebService, fn restful.RouteFunction, method, path, tag string, fns ...routeOptionFunc) *restful.WebService {
	// 增加限流
	// 每秒显示请求数 100
	r := ratelimit.New(20)
	fnn := func(req *restful.Request, rsp *restful.Response) {
		r.Take()
		fn(req, rsp)
	}

	var rb *restful.RouteBuilder
	if method == "POST" {
		rb = ws.POST(path).To(fnn)
	}
	if method == "GET" {
		rb = ws.GET(path).To(fnn)
	}

	rb = rb.Metadata(restfulApi.KeyOpenAPITags, []string{tag})

	for _, fn := range fns {
		rb = fn(rb)
	}
	return ws.Route(rb)
}

func main() {
	// create new web service
	service := web.NewService(
		web.Name("iunite.club.web.restful"),
		web.Version("latest"),
	)

	// service
	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	// service.Handle("/ht/", http.FileServer(http.Dir("html")))
	// service.Handle("/ht/", http.StripPrefix("/ht/", http.FileServer(http.Dir("html"))))
	service.Handle("/apidoc/", http.StripPrefix("/apidoc/", http.FileServer(http.Dir("swigger"))))

	service.Handle("/", getWebContainer())

	// service.Client().
	// service.Ha

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
