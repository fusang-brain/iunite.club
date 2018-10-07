package main

import (
	"log"

	"github.com/iron-kit/go-ironic"
	api "github.com/micro/go-api"
	apiHandler "github.com/micro/go-api/handler/api"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"iunite.club/services/navo/client"
	"iunite.club/services/navo/handler"
	core "iunite.club/services/navo/proto"
	auth "iunite.club/services/navo/proto/auth"
	school "iunite.club/services/navo/proto/school"
	// example "iunite.club/services/navo/proto/example"
)

func apiEndpoint(name, method, path string) server.HandlerOption {
	return api.WithEndpoint(&api.Endpoint{
		Name:    name,
		Method:  []string{method},
		Path:    []string{path},
		Handler: apiHandler.Handler,
	})
}

func main() {
	// New Service
	service := ironic.NewService(
		micro.Name("iunite.club.api.navo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		// micro.WrapHandler(client.ExampleWrapper(service)),
		micro.WrapHandler(
			client.UserServiceWrapper(service),
			// client.OrganizationServiceWrapper(service),
		),
		micro.WrapHandler(
			client.OrganizationServiceWrapper(service),
			client.MessageServiceWrapper(service),
		),
	)

	// Register Handler
	auth.RegisterAuthHandlerHandler(
		service.Server(),
		new(handler.AuthHandler),

		api.WithEndpoint(
			&api.Endpoint{
				Name:    "AuthHandler.Login",
				Path:    []string{"/v1/login"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "AuthHandler.Register",
				Path:    []string{"/v1/register"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
	)

	core.RegisterSMSHandlerHandler(
		service.Server(),
		new(handler.SMSHandler),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "SMSHandler.SendVerifyCode",
				Path:    []string{"/v1/sms/sendVerifyCode"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "SMSHandler.ValidateSimpleCode",
				Path:    []string{"/v1/sms/validateSimpleCode"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
	)

	core.RegisterOrganizationHandlerHandler(
		service.Server(),
		new(handler.OrganizationHandler),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "OrganizationHandler.CreateOrganization",
				Path:    []string{"/v1/organization/createOrganization"},
				Method:  []string{"POST"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "OrganizationHandler.GetAllOrgByUserID",
				Path:    []string{"/v1/organization/getAllOrgByUserID"},
				Method:  []string{"GET"},
				Handler: apiHandler.Handler,
			},
		),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "OrganizationHandler.GetAllOrgUsersByUserID",
				Path:    []string{"/v1/organization/getAllOrgUsersByUserID"},
				Method:  []string{"GET"},
				Handler: apiHandler.Handler,
			},
		),
		apiEndpoint("OrganizationHandler.SearchHotOrganization", "GET", "/v1/organization/searchHotOrganization"),
		apiEndpoint("OrganizationHandler.AcceptJoin", "POST", "/v1/organization/acceptJoin"),
		apiEndpoint("OrganizationHandler.AgreeJoin", "POST", "/v1/organization/agreeJoin"),
		apiEndpoint("OrganizationHandler.RefuseJoin", "POST", "/v1/organization/refuseJoin"),
		apiEndpoint("OrganizationHandler.FindRefusedAccept", "GET", "/v1/organization/findRefusedAccept"),
		apiEndpoint("OrganizationHandler.GetDepartmentDetails", "GET", "/v1/organization/getDepartmentDetails"),
		apiEndpoint("OrganizationHandler.Info", "GET", "/v1/organization/info"),
		apiEndpoint("OrganizationHandler.UploadLogo", "POST", "/v1/organization/uploadLogo"),
		apiEndpoint("OrganizationHandler.UpdateOrganizationDescription", "POST", "/v1/organization/updateOrganizationDescription"),
		apiEndpoint("OrganizationHandler.GetAllOrganizationBySchool", "GET", "/v1/organization/getAllOrganizationBySchool"),
		apiEndpoint("OrganizationHandler.GetOrganizationDetails", "GET", "/v1/organization/getOrganizationDetails"),
		apiEndpoint("OrganizationHandler.GetOrganizationUserInfoDetails", "GET", "/v1/organization/getOrganizationUserInfoDetails"),
		apiEndpoint("OrganizationHandler.SelectOrganizations", "POST", "/v1/organization/selectOrganizations"),
	)

	school.RegisterSchoolHandlerHandler(
		service.Server(),
		new(handler.SchoolSrv),
		apiEndpoint("SchoolHandler.SchoolList", "GET", "/v1/school/allSchools"),
		apiEndpoint("SchoolHandler.SearchSchools", "GET", "/v1/school/searchSchools"),
		apiEndpoint("SchoolHandler.Create", "POST", "/v1/school/create"),
	)

	core.RegisterUserHandlerHandler(
		service.Server(),
		new(handler.UserHandler),
		apiEndpoint("UserHandler.Info", "GET", "/v1/user/info"),
		apiEndpoint("UserHandler.UpdateCurrentOrg", "POST", "/v1/user/UpdateCurrentOrg"),
		apiEndpoint("UserHandler.ForgetPassword", "POST", "/v1/user/forgetPassword"),
		apiEndpoint("UserHandler.AllUser", "GET", "/v1/user/allUser"),
		apiEndpoint("UserHandler.GetCurrentOrganization", "", "/v1/user/getCurrentOrganization"),
		apiEndpoint("UserHandler.GetAllMembers", "GET", "/v1/user/getAllMembers"),
		apiEndpoint("UserHandler.CreateMember", "POST", "/v1/user/createMember"),
		apiEndpoint("UserHandler.RemvoeMemberFromOrg", "POST", "/v1/user/removeMemberFromOrg"),
		apiEndpoint("UserHandler.UpdateMember", "POST", "/v1/user/updateMember"),
		apiEndpoint("UserHandler.GetMemberDetails", "GET", "/v1/user/getMemberDetails"),
		apiEndpoint("UserHandler.RemoveOrg", "POST", "/v1/user/removeOrg"),
		apiEndpoint("UserHandler.UpdateUserInfo", "POST", "/v1/user/updateUserInfo"),
		apiEndpoint("UserHandler.FlagMemberState", "POST", "/v1/user/flagMemberState"),
		apiEndpoint("UserHandler.GetHotUsers", "GET", "/v1/user/getHotUsers"),
		apiEndpoint("UserHandler.UploadAvatar", "POST", "/v1/user/uploadAvatar"),
		apiEndpoint("UserHandler.ExportList", "GET", "/v1/user/exportList"),
		apiEndpoint("UserHandler.DownloadExportTemplate", "GET", "/v1/user/downloadExportTemplate"),
		apiEndpoint("UserHandler.UploadUserList", "POST", "/v1/user/uploadUserList"),
	)
	// example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
