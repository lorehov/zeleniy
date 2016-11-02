package httpapi

import (
	"github.com/emicklei/go-restful"
	"net/http"
)


func handle400(response *restful.Response, err error) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusBadRequest, err.Error())

}


func handle404(response *restful.Response, err error) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusNotFound, err.Error())

}


func handle500(response *restful.Response, err error) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusInternalServerError, err.Error())
}
