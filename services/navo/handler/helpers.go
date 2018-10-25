package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/errors"
)

type D map[string]interface{}

func WriteError(rsp *restful.Response, err error) {
	parsedErr := errors.Parse(err.Error())

	rsp.WriteAsJson(parsedErr)
}

func pbToMap(pb proto.Message) map[string]interface{} {
	str := pbToJsonString(pb)
	fmt.Println(str)
	result := make(map[string]interface{})
	json.Unmarshal([]byte(str), &result)
	return result
}

func pbToJsonString(pb proto.Message) string {
	marshal := new(jsonpb.Marshaler)
	// out := make([]byte, 0)
	// b := new(bytes.Buffer)
	str := new(strings.Builder)
	// b := new(bytes.Buffer)

	marshal.Marshal(str, pb)

	return str.String()
}

func pbToBytes(pb proto.Message) []byte {
	marshal := new(jsonpb.Marshaler)
	// out := make([]byte, 0)
	// b := new(bytes.Buffer)
	str := new(strings.Builder)
	// b := new(bytes.Buffer)

	marshal.Marshal(str, pb)

	return []byte(str.String())
}

func WriteJsonResponse(rsp *restful.Response, data interface{}) {

	if d, ok := data.(proto.Message); ok {
		rsp.WriteAsJson(pbToMap(d))
		return
	}

	if d, ok := data.(D); ok {
		rsp.WriteAsJson(d)
		return
	}

	rsp.WriteAsJson(data)
	return
}
