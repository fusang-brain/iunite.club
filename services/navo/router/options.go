package router

import (
	"github.com/emicklei/go-restful"
)

type Wrapper func(restful.RouteFunction) restful.RouteFunction
type Option func(*Route)
type RouteOption func(*restful.RouteBuilder) *restful.RouteBuilder

func Wrappers(wraps ...Wrapper) Option {
	return func(r *Route) {
		if len(r.wrappers) > 0 {
			r.wrappers = append(r.wrappers, wraps...)
			return
		}

		r.wrappers = wraps
		return
	}
}

// func appendWrappers()
func Description(desc string) Option {
	return func(r *Route) {
		r.description = desc
	}
}

func Name(name string) Option {
	return func(r *Route) {
		r.name = name
	}
}

func Doc(plainText string) Option {
	return func(r *Route) {
		r.ws.Doc(plainText)
	}
}

func APIVersion(apiVersion string) Option {
	return func(r *Route) {
		r.ws.ApiVersion(apiVersion)
	}
}

func Consumes(accepts ...string) Option {
	return func(r *Route) {
		r.ws.Consumes(accepts...)
	}
}

func Produces(contentTypes ...string) Option {
	return func(r *Route) {
		r.ws.Produces(contentTypes...)
	}
}

func RouteDoc(text string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Doc(text)
	}
}

func RouteConsumes(mimeTypes ...string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Consumes(mimeTypes...)
	}
}

func RouteDefaultReturns(message string, model interface{}) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.DefaultReturns(message, model)
	}
}

func RouteMetadata(key string, value interface{}) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Metadata(key, value)
	}
}

func RouteParam(parameter *restful.Parameter) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Param(parameter)
	}
}

func RouteParams(parameters ...*restful.Parameter) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		for _, parameter := range parameters {
			rb = rb.Param(parameter)
		}
		return rb
	}
}

func RouteProduces(mimeTypes ...string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Produces(mimeTypes...)
	}
}

func RouteReads(sample interface{}, optionalDescription ...string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Reads(sample, optionalDescription...)
	}
}

func RouteReturns(code int, message string, model interface{}) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		// rb.Writes()
		return rb.Returns(code, message, model)
	}
}

func RouteReturnsError(code int, message string, model interface{}) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.ReturnsError(code, message, model)
	}
}

func RouteWrites(sample interface{}) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Writes(sample)
	}
}

func RouteNotes(notes string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		return rb.Notes(notes)
	}
}

func RouteOpration(name string) RouteOption {
	return func(rb *restful.RouteBuilder) *restful.RouteBuilder {
		rb.Consumes()
		return rb.Operation(name)
	}
}
