package middleware

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	Resource resource.Resource
}

func NewMiddleware(resource resource.Resource) Middleware {
	return Middleware{
		Resource: resource,
	}
}

func (m *Middleware) BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return nil
}
