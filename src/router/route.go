package router

import (
	"hubimage/src/api"
	"hubimage/src/model"
	"hubimage/src/router/filter"
	"hubimage/swagger"

	restful "github.com/emicklei/go-restful"
)

// CreateHTTPRouter create http router for app
func CreateHTTPRouter() *restful.Container {

	wsContainer := restful.NewContainer()
	wsContainer.Filter(filter.LogRequestAndResponse)
	ws := new(restful.WebService)

	ws.Path("/api/v1").
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	/************** images *************/
	// create image
	// curl   -F "file=@busybox.tgz" -F "tag=test" -F "enableWait=true"  http://10.130.177.25:30909/api/v1/image/library/busybox
	ws.Route(ws.POST("/image/{project}/{name}").
		To(api.HandleCreateImage).
		Writes(model.ImageDetail{}))

	// get image list
	// curl http://10.110.214.232:80/api/v1/image?projectName=library
	ws.Route(ws.GET("/image").
		To(api.HandleGetImageList).
		Writes(model.ResponseImageList{}))

	// get image detail
	ws.Route(ws.GET("/image/{project}/{name}").
		To(api.HandleGetImageDetail).
		Writes(model.ResponseImageDetail{}))

	// delete image
	ws.Route(ws.DELETE("/image/{name}").
		To(api.HandleDeleteImage).
		Reads(model.RequestImageDelete{}))

	/************** projects *************/
	// create project
	// curl http://10.130.177.25:30909/api/v1/project/test
	ws.Route(ws.POST("/project/{name}").
		To(api.HandleCreateProject).
		Writes(swagger.Project{}))

	// get project list
	// curl -X GET http://10.130.177.25:30909/api/v1/project
	ws.Route(ws.GET("/project").
		To(api.HandleGetProjectList).
		Writes(model.ResponseProjectList{}))

	// get project detail
	ws.Route(ws.GET("/project/{name}").
		To(api.HandleGetProjectDetail).
		Writes(swagger.Project{}))

	// delete project
	ws.Route(ws.DELETE("/project/{name}").
		To(api.HandleDeleteProject))

	/************** users *************/
	// create user
	ws.Route(ws.POST("/user").
		To(api.HandleCreateUser).
		Reads(model.RequestUserCreate{}).
		Writes(swagger.User{}))

	// get user list
	ws.Route(ws.GET("/user").
		To(api.HandleGetUserList).
		Writes(model.ResponseUserList{}))

	// get user detail
	ws.Route(ws.GET("/user/{name}").
		To(api.HandleGetUserDetail).
		Writes(swagger.User{}))

	// delete user
	ws.Route(ws.DELETE("/user/{name}").
		To(api.HandleDeleteUser))

	// test
	ws.Route(ws.GET("/test").
		To(api.HandleTest))

	wsContainer.Add(ws)
	return wsContainer
}
