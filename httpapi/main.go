package httpapi


import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/lorehov/zeleniy"
	"github.com/emicklei/go-restful/swagger"
)


func StartApi(app *zeleniy.Application) {
	wsContainer := restful.NewContainer()

	pr := ProjectResource{app}
	pr.Register(wsContainer)
	mr := MetricResource{app}
	mr.Register(wsContainer)
	cr := ComponentResource{app}
	cr.Register(wsContainer)

	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(),
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",}
		//SwaggerFilePath: "/Users/emicklei/xProjects/swagger-ui/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)

	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	server.ListenAndServe()
}


