package post

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/andiksetyawan/boilerplate_svc/internal/usecase"
)

type Handler struct {
	resource resource.Resource
	usecase  usecase.Post
}

func New(usecase usecase.Post, resource resource.Resource) Handler {
	return Handler{
		resource: resource,
		usecase:  usecase,
	}
}
