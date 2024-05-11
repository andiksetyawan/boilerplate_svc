package usecase

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/google/uuid"
)

type Post interface {
	GetByID(ctx context.Context, id uuid.UUID) (post entity.Post, err error)
}

//type Comment interface {
//	GetByID(ctx context.Context, id uuid.UUID) (comment entity.Comment, err error)
//}
