package handler

const (
	UserSerivce         = "iunite.club.srv.user"
	OrganizationService = "iunite.club.srv.organization"
	SMSService          = "iunite.club.srv.message"
	StorageService      = "iunite.club.srv.storage"
	CoreService         = "iunite.club.srv.core"
)

// func ExampleCall(w http.ResponseWriter, r *http.Request) {
// 	// decode the incoming request as json
// 	var request map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// call the backend service
// 	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
// 	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
// 		Name: request["name"].(string),
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// we want to augment the response
// 	response := map[string]interface{}{
// 		"msg": rsp.Msg,
// 		"ref": time.Now().UnixNano(),
// 	}

// 	// encode and write the response as json
// 	if err := json.NewEncoder(w).Encode(response); err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// }

// func GetUserProducts(req *restful.Request, res *restful.Response) {
// 	id := req.PathParameter("sku")
// 	// req.ReadEntity()
// 	// req.Request.Form
// 	u := struct {
// 		Sku string
// 	}{
// 		Sku: id,
// 	}

// 	res.WriteEntity(u)
// }
