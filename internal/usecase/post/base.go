package post

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/repository"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/andiksetyawan/boilerplate_svc/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type postUsecase struct {
	resource       resource.Resource
	tracer         trace.Tracer
	postRepository repository.Post
}

func NewUsecase(postRepository repository.Post, resource resource.Resource) usecase.Post {
	return &postUsecase{
		tracer:         otel.Tracer("postUsecase"),
		resource:       resource,
		postRepository: postRepository,
	}
}
