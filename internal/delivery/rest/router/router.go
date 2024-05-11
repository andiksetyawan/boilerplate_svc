package router

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler"
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/middleware"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/labstack/echo/v4"
)

type router struct {
	resource   resource.Resource
	handler    handler.Handler
	middleware middleware.Middleware
}

func NewRouter(resource resource.Resource, handler handler.Handler, middleware middleware.Middleware) router {
	return router{
		resource:   resource,
		handler:    handler,
		middleware: middleware,
	}
}

func (r router) RegisterRoutes(ec *echo.Echo) {
	//server.

	ec.GET("/post/:id", r.handler.Post.GetByID)
}
