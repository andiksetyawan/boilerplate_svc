package post

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/andiksetyawan/database/sqlx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (r *postRepository) GetByID(ctx context.Context, q sqlx.Q, id uuid.UUID) (post entity.Post, err error) {
	_, span := r.tracer.Start(ctx, "repository.post.GetByID", trace.WithAttributes(
		attribute.Stringer("postID", id),
	))
	defer span.RecordError(err)
	defer span.End()

	err = q.GetContext(ctx, &post, "SELECT * FROM post WHERE id=$1", id)
	return
}
