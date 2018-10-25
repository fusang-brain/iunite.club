package router

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
	HEAD   = "HEAD"
)

type Router struct {
	title        string
	description  string
	contact      *spec.ContactInfo
	license      *spec.License
	version      string
	container    *restful.Container
	defaultRoute *Route
	groups       map[string]*Route
}

func NewRouter(options ...Option) *Router {
	route := newRoute("default", "/", options...)
	container := restful.NewContainer()
	container.Add(route.ws)
	return &Router{
		container:    container,
		defaultRoute: route,
		groups:       make(map[string]*Route),
	}
}

func (self *Router) Title(title string) *Router {
	self.title = title

	return self
}

func (self *Router) Description(description string) *Router {
	self.description = description
	return self
}

func (self *Router) Contact(info *spec.ContactInfo) *Router {
	self.contact = info
	return self
}

func (self *Router) License(license *spec.License) *Router {
	self.license = license
	return self
}

func (self *Router) Version(version string) *Router {
	self.version = version
	return self
}

func (self *Router) GetWS() *restful.WebService {
	return self.defaultRoute.ws
}

func (self *Router) getSwaggerSpecInfo() *spec.Info {
	return &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       self.title,
			Description: self.description,
			Contact:     self.contact,
			License:     self.license,
			Version:     self.version,
		},
	}
}

func (self *Router) getSwaggerTags() []spec.Tag {
	tags := make([]spec.Tag, 0, len(self.groups)+1)
	for _, r := range self.groups {
		tags = append(tags, spec.Tag{
			TagProps: spec.TagProps{
				Name:        r.name,
				Description: r.description,
			},
		})
	}
	tags = append(tags, spec.Tag{
		TagProps: spec.TagProps{
			Name:        self.defaultRoute.name,
			Description: self.defaultRoute.description,
		},
	})
	return tags
}

func (self *Router) getSwaggerConfig() restfulspec.Config {
	swiggerConfig := restfulspec.Config{
		WebServices: self.container.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: func(swo *spec.Swagger) {
			swo.Info = self.getSwaggerSpecInfo()
			swo.Tags = self.getSwaggerTags()
			swo.Schemes = []string{"http"}
		},
	}

	return swiggerConfig
}

func (self *Router) GetContainer() *restful.Container {

	for _, v := range self.groups {
		self.container.Add(v.ws)
	}

	// add swagger support
	self.container.Add(restfulspec.NewOpenAPIService(self.getSwaggerConfig()))
	return self.container
}

func (self *Router) Group(name, path string, options ...Option) *Route {
	options = append(options, Wrappers(self.defaultRoute.wrappers...))
	if self.groups == nil {
		self.groups = make(map[string]*Route)
	}
	if route, found := self.groups[name]; found {
		return route
	}

	route := newRoute(name, path, options...)

	self.groups[name] = route

	return route
}

func (self *Router) GET(subPath string, handler restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.GET(subPath, handler, opts...)
}

func (self *Router) POST(subPath string, handler restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.POST(subPath, handler, opts...)
}

func (self *Router) PUT(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.PUT(subPath, fn, opts...)
}

func (self *Router) DELETE(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.DELETE(subPath, fn, opts...)
}

func (self *Router) HEAD(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.HEAD(subPath, fn, opts...)
}

func (self *Router) PATCH(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.defaultRoute.PATCH(subPath, fn, opts...)
}
