package main

import (
	"iunite.club/services/navo/routers"
	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
	"github.com/micro/go-log" // "net/http"
	// "github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"iunite.club/services/navo/handler"
	"iunite.club/services/navo/router"
)

func getRouter() *router.Router {
	// c := client.DefaultClient

	// announceHandler := handler.NewAnnounceHandler(c)

	r := router.NewRouter(
		router.Name("Navo"),
    router.Description("Navo API"),
    router.Produces(restful.MIME_JSON),
	)

	r.Title("Unite - Navo").
		Description("Unite API 服务").
		Contact(&spec.ContactInfo{
			Name:  "alixe",
			Email: "alixe.z@foxmail.com",
			URL:   "http://blog.ironkit.xyz",
		}).
		License(&spec.License{
			Name: "MIT",
			URL:  "http://mit.org",
		}).
    Version("Navo v0.0.1")

	// 通告
	routers.AnnounceRoute(r)
	routers.ReportRoute(r)
	routers.ApprovedRoute(r)
	return r
}

func main() {
	// create new web service
	service := web.NewService(
		web.Name("iunite.club.web.navo"),
		web.Version("latest"),
	)

	r := getRouter()

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// service.Handle("/", )

	// register html handler
	// service.Handle("/", http.FileServer(http.Dir("html")))
	container := r.GetContainer()
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "U-Access-Token"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      container,
	}
	container.Filter(cors.Filter)
	service.Handle("/", container)

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
	// service.Init(web.DefaultAddress)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
