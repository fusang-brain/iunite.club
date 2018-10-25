package router

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

type Route struct {
	name        string
	tag         string
	description string
	ws          *restful.WebService
	wrappers    []Wrapper
}

func newRoute(name, path string, options ...Option) *Route {
	ws := new(restful.WebService)
	ws.Doc(name).Path(path)
	route := &Route{
		ws:   ws,
		name: name,
		tag:  name,
	}

	if len(options) > 0 {
		for _, o := range options {
			o(route)
		}
	}
	return route
}

func (self *Route) WS() *restful.WebService {
	return self.ws
}

func (self *Route) GET(subPath string, handler restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(GET, subPath, handler, opts...)
}

func (self *Route) POST(subPath string, handler restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(POST, subPath, handler, opts...)
}

func (self *Route) PUT(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(PUT, subPath, fn, opts...)
}

func (self *Route) DELETE(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(DELETE, subPath, fn, opts...)
}

func (self *Route) HEAD(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(HEAD, subPath, fn, opts...)
}

func (self *Route) PATCH(subPath string, fn restful.RouteFunction, opts ...RouteOption) {
	self.routeConfig(PATCH, subPath, fn, opts...)
}

func (self *Route) wrapRouteFunction(fn restful.RouteFunction) restful.RouteFunction {
	if len(self.wrappers) > 0 {
		handler := fn
		for _, f := range self.wrappers {
			handler = f(handler)
		}

		return handler
	}

	return fn
}

func (self *Route) routeConfig(method, subPath string, handler restful.RouteFunction, opts ...RouteOption) {
	// self.ws.GET
	var routeBuilder *restful.RouteBuilder
	switch method {
	case GET:
		routeBuilder = self.ws.GET(subPath)
	case POST:
		routeBuilder = self.ws.POST(subPath)
	case PATCH:
		routeBuilder = self.ws.PATCH(subPath)
	case DELETE:
		routeBuilder = self.ws.DELETE(subPath)
	case PUT:
		routeBuilder = self.ws.PUT(subPath)
	case HEAD:
		routeBuilder = self.ws.HEAD(subPath)
	}

	fn := self.wrapRouteFunction(handler)

	routeBuilder = routeBuilder.To(fn).Metadata(restfulspec.KeyOpenAPITags, []string{self.name})

	if len(opts) > 0 {
		for _, fn := range opts {
			routeBuilder = fn(routeBuilder)
		}
	}

	self.ws.Route(routeBuilder)
}
