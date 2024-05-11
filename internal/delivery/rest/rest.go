package rest

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler"
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/middleware"
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/router"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/andiksetyawan/boilerplate_svc/pkg/httpserver"
)

func RegisterServer(resource resource.Resource, http *httpserver.HTTPServer, handler handler.Handler) {
	m := middleware.NewMiddleware(resource)
	r := router.NewRouter(resource, handler, m)
	r.RegisterRoutes(http.Handler())
}
