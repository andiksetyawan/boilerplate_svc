package repository

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/andiksetyawan/database/sqlx"
	"github.com/google/uuid"
)

type Post interface {
	GetByID(ctx context.Context, q sqlx.Q, id uuid.UUID) (post entity.Post, err error)
}
