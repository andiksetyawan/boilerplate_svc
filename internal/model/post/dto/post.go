package dto

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/google/uuid"
)

type GetPostByIdReq struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type GetPostByIdRes entity.Post
