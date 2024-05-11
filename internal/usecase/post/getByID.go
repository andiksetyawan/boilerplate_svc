package post

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (p *postUsecase) GetByID(ctx context.Context, id uuid.UUID) (post entity.Post, err error) {
	_, span := p.tracer.Start(ctx, "usecase.post.GetByID", trace.WithAttributes(
		attribute.Stringer("postID", id),
	))
	defer span.RecordError(err)
	defer span.End()

	post, err = p.postRepository.GetByID(ctx, p.resource.DB.GetMaster(), id)
	return
}
