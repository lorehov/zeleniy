package httpapi

import (
	"github.com/lorehov/zeleniy"
	"github.com/emicklei/go-restful"
	"net/http"
	"gopkg.in/mgo.v2"
)


type ProjectResource struct {
	App *zeleniy.Application
}


func (p *ProjectResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/projects").
		Doc("Manage Projects").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{project-id}").To(p.getById).
		Doc("get a project by ID").
		Operation("getById").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		Writes(zeleniy.Project{}))

	ws.Route(ws.GET("/").To(p.getList).
		Doc("get list of projects").
		Operation("getList").
		Returns(200, "OK", []zeleniy.Project{}))

	ws.Route(ws.PUT("/{project-id}").To(p.update).
		Doc("update a project").
		Operation("update").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		ReturnsError(409, "duplicate project-id", nil).
		Reads(zeleniy.Project{}))

	ws.Route(ws.POST("").To(p.create).
		Doc("create a project").
		Operation("create").
		Reads(zeleniy.Project{}))

	ws.Route(ws.DELETE("/{project-id}").To(p.removeById).
		Doc("delete a project").
		Operation("remove").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")))

	container.Add(ws)
}


func (p *ProjectResource) getList(request *restful.Request, response *restful.Response) {
	projectService := zeleniy.NewProjectService(p.App.Db())
	projects, err := projectService.GetProjectsList()
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteEntity(projects)
}


func (p *ProjectResource) getById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	projectId, err := zeleniy.ObjectId(id)
	if err != nil {
		handleWrongId(response)
		return
	}
	projectService := zeleniy.NewProjectService(p.App.Db())
	project, err := projectService.GetProjectById(projectId)
	if err != nil {
		if err == mgo.ErrNotFound {
			handle404(response, "404: no such project.")
			return
		} else {
			handle500(response, "500: internal error.")
			return
		}
	}
	response.WriteEntity(project)
}


func (p *ProjectResource) removeById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	projectId, err := zeleniy.ObjectId(id)
	if err != nil {
		handleWrongId(response)
		return
	}
	projectService := zeleniy.NewProjectService(p.App.Db())
	err = projectService.DeleteProject(projectId)
	if err != nil {
		if err == mgo.ErrNotFound {
			handle404(response, "404: no such project.")
			return
		} else {
			handle500(response, "500: internal error.")
			return
		}
	}
	response.Write([]byte{})
}


func (p *ProjectResource) update(request *restful.Request, response *restful.Response) {
	project := zeleniy.Project{}
	err := request.ReadEntity(&project)
	if err != nil {
		handle400(response, err)
		return
	}
	projectService := zeleniy.NewProjectService(p.App.Db())
	_, err = projectService.CreateOrUpdateProject(&project)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, project)

}


func (p *ProjectResource) create(request *restful.Request, response *restful.Response) {
	project := zeleniy.Project{}
	err := request.ReadEntity(&project)
	if err != nil {
		handle400(response, err)
		return
	}
	projectService := zeleniy.NewProjectService(p.App.Db())
	_, err = projectService.CreateOrUpdateProject(&project)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, project)
}


func handleWrongId(response *restful.Response) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusBadRequest, "400: wrong id.")
}


func handle400(response *restful.Response, err error) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusBadRequest, err.Error())

}


func handle404(response *restful.Response, msg string) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusNotFound, msg)

}


func handle500(response *restful.Response, msg string) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusInternalServerError, msg)
}
