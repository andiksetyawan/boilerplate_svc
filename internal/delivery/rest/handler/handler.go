package handler

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler/post"
)

type Handler struct {
	Post post.Handler
	//Comment comment.Handler
}

func NewHandler(post post.Handler) Handler {
	return Handler{
		Post: post,
	}
}
