package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/andiksetyawan/log"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type HTTPServer struct {
	appName    string
	listener   net.Listener
	handler    *echo.Echo
	validator  echo.Validator
	server     *http.Server
	address    string
	tracerName string
	logger     log.Logger
}

type OptFunc func(h *HTTPServer) error

func WithAppName(name string) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.appName = name
		return
	}
}

func WithListener(l net.Listener) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.listener = l
		return
	}
}

func WithTracer(name string) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.tracerName = name
		return
	}
}

func WithLogger(logger log.Logger) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.logger = logger
		return
	}
}

func WithAddress(address string) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.address = address
		return
	}
}

func WithHandler(echo *echo.Echo) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.handler = echo
		return
	}
}

func WithValidator(validator echo.Validator) OptFunc {
	return func(h *HTTPServer) (err error) {
		h.validator = validator
		return
	}
}

func New(opts ...OptFunc) (h *HTTPServer, err error) {
	h = &HTTPServer{handler: echo.New()}

	for _, opt := range opts {
		if err = opt(h); err != nil {
			return
		}
	}

	if h.appName == "" {
		return nil, fmt.Errorf("missing app name")
	}

	if h.logger == nil {
		return nil, fmt.Errorf("missing logger")
	}

	err = h.setup()
	return
}

func (h *HTTPServer) setup() (err error) {
	if h.tracerName != "" {
		h.handler.Use(otelecho.Middleware(h.tracerName))
	}

	h.handler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	if h.validator == nil {
		h.handler.Validator = NewValidator(validator.New())
	}

	h.handler.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, h.appName)
	})

	h.server = &http.Server{
		Handler: h.handler,
	}
	return
}

func (h *HTTPServer) Start(ctx context.Context) (err error) {
	if h.listener == nil {
		if h.address == "" {
			h.server.Addr = ":9999"
		}

		h.logger.Info(ctx, "starting http server", "address", h.server.Addr)
		return h.server.ListenAndServe()
	}

	h.logger.Info(ctx, "starting http server", "address", h.listener.Addr().String())
	return errors.Join(err, h.server.Serve(h.listener))
}

func (h *HTTPServer) Close(ctx context.Context) (err error) {
	err = h.server.Shutdown(ctx)
	if err != nil {
		return
	}

	h.logger.Info(ctx, "http server is shutdown correctly", "address", h.listener.Addr().String())
	return
}

func (h *HTTPServer) Handler() *echo.Echo {
	return h.handler
}
