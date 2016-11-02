package httpapi

import (
	"github.com/emicklei/go-restful"
	"github.com/lorehov/zeleniy"
	"net/http"
	"gopkg.in/mgo.v2"
)


type ComponentResource struct {
	App *zeleniy.Application
}


func (p *ComponentResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/components").
		Doc("Manage Components").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{component-id}").To(p.getById).
		Doc("get a component by ID").
		Operation("getById").
		Param(ws.PathParameter("component-id", "identifier of the component").DataType("string")).
		Writes(zeleniy.Component{}))

	ws.Route(ws.GET("/").To(p.getList).
		Doc("get list of components").
		Operation("getList").
		Returns(200, "OK", []zeleniy.Component{}))

	ws.Route(ws.PUT("/{component-id}").To(p.update).
		Doc("update a component").
		Operation("update").
		Param(ws.PathParameter("component-id", "identifier of the component").DataType("string")).
		ReturnsError(409, "duplicate component-id", nil).
		Reads(zeleniy.Component{}))

	ws.Route(ws.POST("").To(p.create).
		Doc("create a component").
		Operation("create").
		Reads(zeleniy.Component{}))

	ws.Route(ws.DELETE("/{component-id}").To(p.removeById).
		Doc("delete a component").
		Operation("remove").
		Param(ws.PathParameter("component-id", "identifier of the component").DataType("string")))

	container.Add(ws)
}


func (p *ComponentResource) getList(request *restful.Request, response *restful.Response) {
	componentService := zeleniy.NewComponentService(p.App.Db())
	components, err := componentService.GetComponentsList()
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteEntity(components)
}


func (p *ComponentResource) getById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("component-id")
	componentId, err := zeleniy.ObjectId(id)
	if err != nil {
		handle400(response, err)
		return
	}
	componentService := zeleniy.NewComponentService(p.App.Db())
	component, err := componentService.GetComponentById(componentId)
	if err != nil {
		if err == mgo.ErrNotFound {
			handle404(response, err)
			return
		} else {
			handle500(response, err)
			return
		}
	}
	response.WriteEntity(component)
}


func (p *ComponentResource) removeById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("component-id")
	componentId, err := zeleniy.ObjectId(id)
	if err != nil {
		handle400(response, err)
		return
	}
	componentService := zeleniy.NewComponentService(p.App.Db())
	err = componentService.DeleteComponent(componentId)
	if err != nil {
		if err == mgo.ErrNotFound {
			handle404(response, err)
			return
		} else {
			handle500(response, err)
			return
		}
	}
	response.Write([]byte{})
}


func (p *ComponentResource) update(request *restful.Request, response *restful.Response) {
	component := zeleniy.Component{}
	err := request.ReadEntity(&component)
	if err != nil {
		handle400(response, err)
		return
	}
	componentService := zeleniy.NewComponentService(p.App.Db())
	_, err = componentService.CreateOrUpdateComponent(&component)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, component)

}


func (p *ComponentResource) create(request *restful.Request, response *restful.Response) {
	component := zeleniy.Component{}
	err := request.ReadEntity(&component)
	if err != nil {
		handle400(response, err)
		return
	}
	componentService := zeleniy.NewComponentService(p.App.Db())
	_, err = componentService.CreateOrUpdateComponent(&component)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, component)
}

