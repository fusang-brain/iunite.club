package main

import (
	restful "github.com/emicklei/go-restful"
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
	fileHandler := handler.NewFileHandler(c)
	schoolHandler := handler.NewSchoolHandler(c)
	authHandler := handler.NewAuthHandler(c)
	smsHandler := handler.NewSMSHandler(c)
	organizationHandler := handler.NewOrganizationHandler(c)
	departmentHandler := handler.NewDepartmentHandler(c)
	jobHandler := handler.NewJobService(c)
	userHandler := handler.NewUserHandler(c)
	activityHandler := handler.NewActivityHandler(c)

	rc := restful.NewContainer()

	// Base service
	base := getWebService("/v1", "Base service")
	base.Route(base.POST("/login").To(authHandler.Login)).Doc("登录")
	base.Route(base.POST("/register").To(authHandler.Register)).Doc("注册")

	// File service
	file := getWebService("/v1/file", "File service")
	file.Route(file.POST("/uploadSingleFile").To(fileHandler.UploadSingleFile)).Doc("上传单文件")
	file.Route(file.POST("/uploadMultipartFile").To(fileHandler.UploadMutipartFile)).Doc("上传多文件")

	// School Serivce
	school := getWebService("/v1/school", "School service")
	school.Route(school.GET("/allSchools").To(schoolHandler.SchoolList)).Doc("获取学校列表")
	school.Route(school.GET("/searchSchools").To(schoolHandler.SearchSchools)).Doc("搜索学校")
	school.Route(school.POST("/create").To(schoolHandler.Create)).Doc("创建学校")

	// SMS Service
	sms := getWebService("/v1/sms", "SMS service")
	sms.Route(sms.POST("/sendVerifyCode").To(smsHandler.SendVerifyCode)).Doc("发送短信验证码")
	sms.Route(sms.POST("/validateSimpleCode").To(smsHandler.ValidateSimpleCode)).Doc("验证短信验证码")

	// Organization Service
	// Organization Handler
	organization := getWebService("/v1/organization", "Organization service")
	organization.Route(organization.POST("/createOrganization").To(organizationHandler.CreateOrganization)).Doc("创建社团组织")
	setRoute(organization, organizationHandler.GetAllOrgByUserID, "GET", "/getAllOrgByUserID").Doc("通过用户ID获取所有组织")
	setRoute(organization, organizationHandler.GetAllOrgUsersByUserID, "GET", "/getAllOrgUsersByUserID").Doc("通过用户ID获取用户社团信息")
	setRoute(organization, organizationHandler.SearchHotOrganization, "GET", "/searchHotOrganization").Doc("搜索热门社团")
	setRoute(organization, organizationHandler.AcceptJoin, "POST", "/acceptJoin").Doc("同意社团申请")
	setRoute(organization, organizationHandler.AgreeJoin, "POST", "/agreeJoin")
	setRoute(organization, organizationHandler.RefuseJoin, "POST", "/refuseJoin")
	setRoute(organization, organizationHandler.FindRefusedAccept, "GET", "/findRefusedAccept")
	setRoute(organization, organizationHandler.GetDepartmentDetails, "GET", "/getDepartmentDetails")
	setRoute(organization, organizationHandler.Info, "GET", "/info")
	setRoute(organization, organizationHandler.UploadLogo, "POST", "/uploadLogo")
	setRoute(organization, organizationHandler.UpdateOrganizationDescription, "POST", "/updateOrganizationDescription")
	setRoute(organization, organizationHandler.GetAllOrganizationBySchool, "GET", "/getAllOrganizationBySchool")
	setRoute(organization, organizationHandler.GetOrganizationDetails, "GET", "/getOrganizationDetails")
	setRoute(organization, organizationHandler.GetOrganizationUserInfoDetails, "GET", "/getOrganizationUserInfoDetails")
	setRoute(organization, organizationHandler.SelectOrganizations, "POST", "/selectOrganizations")

	// User Service
	user := getWebService("/v1/user", "User Service")
	setRoute(user, userHandler.Info, "GET", "/info")
	setRoute(user, userHandler.UpdateCurrentOrg, "POST", "/updateCurrentOrg")
	setRoute(user, userHandler.ForgetPassword, "POST", "/forgetPassword")
	setRoute(user, userHandler.AllUser, "GET", "/allUser")
	setRoute(user, userHandler.GetCurrentOrganization, "GET", "/getCurrentOrganization")
	setRoute(user, userHandler.GetAllMembers, "GET", "/getAllMembers")
	setRoute(user, userHandler.CreateMember, "POST", "/createMember")
	setRoute(user, userHandler.RemoveMemberFromOrg, "POST", "/removeMemberFromOrg")
	setRoute(user, userHandler.UpdateMember, "POST", "/updateMember")
	setRoute(user, userHandler.GetMemberDetails, "GET", "/getMemberDetails")
	setRoute(user, userHandler.RemoveOrg, "POST", "/removeOrg")
	setRoute(user, userHandler.UpdateUserInfo, "POST", "/updateUserInfo")
	setRoute(user, userHandler.FlagMemberState, "POST", "/flagMemberState")
	setRoute(user, userHandler.GetHotUsers, "GET", "/getHotUsers")
	setRoute(user, userHandler.UploadAvatar, "POST", "/uploadAvatar")
	setRoute(user, userHandler.ExportList, "GET", "/exportList")
	setRoute(user, userHandler.DownloadExportTemplate, "GET", "/downloadExportTemplate")
	setRoute(user, userHandler.UploadUserList, "POST", "/uploadUserList")

	job := getWebService("/v1/job", "Job Service")
	setRoute(job, jobHandler.CreateJob, "POST", "/createJob")
	setRoute(job, jobHandler.GetUsersWithJob, "GET", "/getUsersWithJob")
	setRoute(job, jobHandler.AllCanSelectedUsers, "GET", "/allCanSelectedUsers")
	setRoute(job, jobHandler.AddUsersToJob, "POST", "/addUsersToJob")
	setRoute(job, jobHandler.RemoveUsersFromJob, "POST", "/removeUsersFromJob")
	setRoute(job, jobHandler.GetAllJobs, "GET", "/getAllJobs")
	setRoute(job, jobHandler.Update, "POST", "/update")
	setRoute(job, jobHandler.Remove, "POST", "/deleteJob")

	department := getWebService("/v1/department", "Department Service")
	setRoute(department, departmentHandler.AddDept, "POST", "/addDept")
	setRoute(department, departmentHandler.GetDepartmentByOrg, "GET", "/getDepartmentByOrg")
	setRoute(department, departmentHandler.GetAllDepartmentByOrg, "GET", "/getAllDepartmentByOrg")
	setRoute(department, departmentHandler.SearchDepartment, "GET", "/searchDepartment")
	setRoute(department, departmentHandler.AddUserToDepartment, "POST", "/addUserToDepartment")
	setRoute(department, departmentHandler.RemoveUserFromDepartment, "POST", "/removeUserFromDepartment")
	setRoute(department, departmentHandler.AllCanSelectedUsers, "GET", "/allCanSelectedUsers")
	setRoute(department, departmentHandler.GetAllUsersWithDepartment, "GET", "/getAllUsersWithDepartment")
	setRoute(department, departmentHandler.Update, "POST", "/update")
	setRoute(department, departmentHandler.Remove, "POST", "/remove")

	activity := getWebService("/v1/activity", "Activity Service")
	setRoute(activity, activityHandler.CreateActivity, "POST", "/create")

	rc.Add(file)
	rc.Add(school)
	rc.Add(base)
	rc.Add(organization)
	rc.Add(sms)
	rc.Add(user)
	rc.Add(job)
	rc.Add(department)
	rc.Add(activity)

	return rc
}

func setRoute(ws *restful.WebService, fn restful.RouteFunction, method, path string) *restful.WebService {
	var routerBuilder *restful.RouteBuilder
	if method == "POST" {
		routerBuilder = ws.POST(path).To(fn)
	}
	if method == "GET" {
		routerBuilder = ws.GET(path).To(fn)
	}
	return ws.Route(routerBuilder)
}

func main() {
	// create new web service
	service := web.NewService(
		web.Name("iunite.club.web.restful"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	// service.Handle("/api", http.FileServer(http.Dir("html")))
	service.Handle("/", getWebContainer())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
