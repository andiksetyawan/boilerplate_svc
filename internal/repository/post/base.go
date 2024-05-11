package post

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/repository"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type postRepository struct {
	resource resource.Resource
	tracer   trace.Tracer
}

func NewRepository(resource resource.Resource) repository.Post {
	return &postRepository{
		tracer:   otel.Tracer("postRepository"),
		resource: resource,
	}
}
