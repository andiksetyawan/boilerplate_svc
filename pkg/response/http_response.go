package response

import (
	"net/http"
	"strings"

	"github.com/andiksetyawan/log"
	"github.com/labstack/echo/v4"
)

type HttpResponse struct {
	log      log.Logger
	logAttrs []any
}

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Response
	Errors []string `json:"errors,omitempty"`
}

type OptFunc func(o *HttpResponse) error

func WithErrLogger(logger log.Logger, attrs ...any) OptFunc {
	return func(h *HttpResponse) (err error) {
		h.log = logger
		h.logAttrs = append(h.logAttrs, attrs...)
		return
	}
}

func New(opt ...OptFunc) (h *HttpResponse, err error) {
	h = new(HttpResponse)

	for _, fn := range opt {
		err = fn(h)
		if err != nil {
			return
		}
	}

	return
}

func (h HttpResponse) Error(ctx echo.Context, err error, msg string) error {
	return h.ErrorWithStatus(ctx, http.StatusInternalServerError, err, msg)
}

func (h HttpResponse) ErrorWithStatus(ctx echo.Context, code int, err error, msg string) error {
	if h.log != nil {
		var attrs []any
		if len(h.logAttrs) != 0 {
			attrs = append(attrs, h.logAttrs)
		}

		attrs = append(attrs, "error", err)
		h.log.Error(ctx.Request().Context(), "http response error", attrs...)
	}

	return ctx.JSON(code, ErrorResponse{
		Response: Response{
			Status:     "error", //TODO
			StatusCode: code,
			Message:    msg,
		},
		Errors: strings.Split(err.Error(), "\n"),
	})
}

func (h HttpResponse) Success(ctx echo.Context, msg string, data interface{}) error {
	return h.SuccessWithStatus(ctx, http.StatusOK, msg, data)
}

func (h HttpResponse) SuccessWithStatus(ctx echo.Context, code int, msg string, data interface{}) error {
	return ctx.JSON(code, Response{
		Status:     "OK", //TODO
		StatusCode: code,
		Message:    msg,
		Data:       data,
	})
}
