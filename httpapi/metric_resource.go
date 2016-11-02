package httpapi

import (
	"github.com/emicklei/go-restful"
	"github.com/lorehov/zeleniy"
	"net/http"
	"gopkg.in/mgo.v2"
)


type MetricResource struct {
	App *zeleniy.Application
}


func (p *MetricResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/metrics").
		Doc("Manage Metrics").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{metric-id}").To(p.getById).
		Doc("get a metric by ID").
		Operation("getById").
		Param(ws.PathParameter("metric-id", "identifier of the metric").DataType("string")).
		Writes(zeleniy.Metric{}))

	ws.Route(ws.GET("/").To(p.getList).
		Doc("get list of metrics").
		Operation("getList").
		Returns(200, "OK", []zeleniy.Metric{}))

	ws.Route(ws.PUT("/{metric-id}").To(p.update).
		Doc("update a metric").
		Operation("update").
		Param(ws.PathParameter("metric-id", "identifier of the metric").DataType("string")).
		ReturnsError(409, "duplicate metric-id", nil).
		Reads(zeleniy.Metric{}))

	ws.Route(ws.POST("").To(p.create).
		Doc("create a metric").
		Operation("create").
		Reads(zeleniy.Metric{}))

	ws.Route(ws.DELETE("/{metric-id}").To(p.removeById).
		Doc("delete a metric").
		Operation("remove").
		Param(ws.PathParameter("metric-id", "identifier of the metric").DataType("string")))

	container.Add(ws)
}


func (p *MetricResource) getList(request *restful.Request, response *restful.Response) {
	metricService := zeleniy.NewMetricService(p.App.Db())
	metrics, err := metricService.GetMetricsList()
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteEntity(metrics)
}


func (p *MetricResource) getById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("metric-id")
	metricId, err := zeleniy.ObjectId(id)
	if err != nil {
		handle400(response, err)
		return
	}
	metricService := zeleniy.NewMetricService(p.App.Db())
	metric, err := metricService.GetMetricById(metricId)
	if err != nil {
		if err == mgo.ErrNotFound {
			handle404(response, err)
			return
		} else {
			handle500(response, err)
			return
		}
	}
	response.WriteEntity(metric)
}


func (p *MetricResource) removeById(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("metric-id")
	metricId, err := zeleniy.ObjectId(id)
	if err != nil {
		handle400(response, err)
		return
	}
	metricService := zeleniy.NewMetricService(p.App.Db())
	err = metricService.DeleteMetric(metricId)
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


func (p *MetricResource) update(request *restful.Request, response *restful.Response) {
	metric := zeleniy.Metric{}
	err := request.ReadEntity(&metric)
	if err != nil {
		handle400(response, err)
		return
	}
	metricService := zeleniy.NewMetricService(p.App.Db())
	_, err = metricService.CreateOrUpdateMetric(&metric)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, metric)

}


func (p *MetricResource) create(request *restful.Request, response *restful.Response) {
	metric := zeleniy.Metric{}
	err := request.ReadEntity(&metric)
	if err != nil {
		handle400(response, err)
		return
	}
	metricService := zeleniy.NewMetricService(p.App.Db())
	_, err = metricService.CreateOrUpdateMetric(&metric)
	if err != nil {
		handle500(response, err.Error())
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, metric)
}

