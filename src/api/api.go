package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"

	"hubimage/src/model"
	"hubimage/src/pkg/images"
	"hubimage/src/pkg/projects"
	"hubimage/src/pkg/users"
	"os"
	"path/filepath"
	"strconv"
)

const FILEMAXSIZE = 1024 * 1024 * 1024

type Size interface {
	Size() int64
}

type Stat interface {
	Stat() (os.FileInfo, error)
}

func GetFilePath(r *http.Request) (string, error) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(256 << 20)

	file, handler, err := r.FormFile("file") //retrieve the file from data
	if err != nil {
		return "", err

	}
	defer file.Close() //close the file when we finish

	// get size
	if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		log.Printf("upload file: %s size: %d", handler.Filename, fileInfo.Size())
		if fileInfo.Size() > FILEMAXSIZE {
			msg := fmt.Sprintf("image \"%s\" size \"%d\" too large, it should less than 1GB.", handler.Filename, fileInfo.Size())
			return "", errors.New(msg)
		}
	}

	// delete file
	path, _ := filepath.Abs("./")
	fpath := path + handler.Filename
	os.Remove(fpath)

	//this is path which  we want to store the file
	f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	//here we save our file to our path
	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return fpath, nil
}

// HandleCreateImage create Image
func HandleCreateImage(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateImage")
	requestRaw := new(model.RequestImageCreate)
	requestRaw.ProjectName = request.PathParameter("project")
	requestRaw.Name = request.PathParameter("name")
	requestRaw.Tag = request.QueryParameter("tag")
	var err error
	requestRaw.EnableWait, err = strconv.ParseBool(request.QueryParameter("enableWait"))

	requestRaw.File = request.QueryParameter("file")
	if requestRaw.File == "" {
		requestRaw.File, err = GetFilePath(request.Request)
		if err != nil {
			handleInternalError(response, err)
			return
		}
	}

	result, err := images.CreateImage(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetImageList get image list
func HandleGetImageList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetImageList")
	requestRaw := new(model.RequestImageList)
	if request.QueryParameter("projectId") != "" {
		projectId, err := strconv.Atoi(request.QueryParameter("projectId"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.ProjectId = int32(projectId)
	}

	requestRaw.ProjectName = request.QueryParameter("projectName")
	if request.QueryParameter("page") != "" {
		page, err := strconv.Atoi(request.QueryParameter("page"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.Page = int32(page)
	}

	if request.QueryParameter("pageSize") != "" {
		pageSize, err := strconv.Atoi(request.QueryParameter("pageSize"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.PageSize = int32(pageSize)
	}

	result, res, err := images.GetImageList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.AddHeader("X-Total-Count", res.Response.Header.Get("X-Total-Count"))
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetImageDetail get detail image info
func HandleGetImageDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetImageDetail")
	requestRaw := new(model.RequestImageDetail)
	requestRaw.ProjectName = request.PathParameter("project")
	requestRaw.ImageName = request.PathParameter("name")
	requestRaw.Tag = request.QueryParameter("tag")

	result, err := images.GetImageDetail(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteImage delete app
func HandleDeleteImage(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteImage")
	name := request.PathParameter("name")
	requestRaw := new(model.RequestImageDelete)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	requestRaw.ImageName = name
	err = images.DeleteImage(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

/******************* projects *********************/
// HandleCreateProject create Project
func HandleCreateProject(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateProject")
	requestRaw := new(model.RequestProjectCreate)
	requestRaw.Name = request.PathParameter("name")

	var err error
	requestRaw.Public, err = strconv.ParseBool(request.QueryParameter("public"))
	if err != nil {
		handleInternalError(response, err)
		return
	}

	result, err := projects.CreateProject(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetProjectList get Project list
func HandleGetProjectList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetProjectList")
	requestRaw := new(model.RequestProjectList)
	requestRaw.Name = request.QueryParameter("name")
	if request.QueryParameter("all") != "" {
		all, err := strconv.ParseBool(request.QueryParameter("all"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.All = all
	}

	if request.QueryParameter("public") != "" {
		public, err := strconv.ParseBool(request.QueryParameter("public"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.Public = public
	}

	if request.QueryParameter("public") == "" && request.QueryParameter("all") == "" {
		requestRaw.All = true
	}

	requestRaw.Owner = request.QueryParameter("owner")
	if request.QueryParameter("page") != "" {
		page, err := strconv.Atoi(request.QueryParameter("page"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.Page = int32(page)
	}

	if request.QueryParameter("pageSize") != "" {
		pageSize, err := strconv.Atoi(request.QueryParameter("pageSize"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.PageSize = int32(pageSize)
	}

	result, res, err := projects.GetProjectList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.AddHeader("X-Total-Count", res.Response.Header.Get("X-Total-Count"))
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetProjectDetail get detail Project info
func HandleGetProjectDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetProjectDetail")
	name := request.PathParameter("name")
	result, err := projects.GetProjectDetail(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteProject delete Project
func HandleDeleteProject(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteProject")
	name := request.PathParameter("name")
	err := projects.DeleteProject(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

/******************* users *********************/
// HandleCreateUser create user
func HandleCreateUser(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleCreateUser")
	requestRaw := new(model.RequestUserCreate)
	err := request.ReadEntity(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	result, err := users.CreateUser(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetUserList get user list
func HandleGetUserList(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetUserList")
	requestRaw := new(model.RequestUserList)
	requestRaw.Username = request.QueryParameter("username")
	requestRaw.Email = request.QueryParameter("email")
	if request.QueryParameter("page") != "" {
		page, err := strconv.Atoi(request.QueryParameter("page"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.Page = int32(page)
	}

	if request.QueryParameter("pageSize") != "" {
		pageSize, err := strconv.Atoi(request.QueryParameter("pageSize"))
		if err != nil {
			handleInternalError(response, err)
			return
		}
		requestRaw.PageSize = int32(pageSize)
	}

	result, err := users.GetUserList(requestRaw)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleGetUserDetail get detail user info
func HandleGetUserDetail(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleGetUserDetail")
	name := request.PathParameter("name")
	result, err := users.GetUserDetail(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// HandleDeleteUser delete user
func HandleDeleteUser(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleDeleteUser")
	name := request.PathParameter("name")
	err := users.DeleteUser(name)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeader(http.StatusOK)
}

// test
func HandleTest(request *restful.Request, response *restful.Response) {
	log.Printf("call HandleTest")

	response.WriteHeader(http.StatusOK)
}
