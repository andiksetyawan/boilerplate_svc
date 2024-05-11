package post

import (
	"fmt"
	"net/http"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/dto"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	req := dto.GetPostByIdReq{}

	err := c.Bind(&req)
	if err != nil {
		err = fmt.Errorf("fail to bind: %w", err)
		return h.resource.HttpRes.ErrorWithStatus(c, http.StatusBadRequest, err, http.StatusText(http.StatusBadRequest))
	}

	if err = c.Validate(&req); err != nil {
		err = fmt.Errorf("fail to validate: %w", err)
		return h.resource.HttpRes.ErrorWithStatus(c, http.StatusBadRequest, err, http.StatusText(http.StatusBadRequest))
	}

	post, err := h.usecase.GetByID(ctx, req.ID)
	if err != nil {
		return h.resource.HttpRes.Error(c, err, http.StatusText(http.StatusInternalServerError))
	}

	return h.resource.HttpRes.Success(c, "fetch post successfully", post)

}
