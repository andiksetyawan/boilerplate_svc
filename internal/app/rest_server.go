package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest"
	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler"
	"github.com/andiksetyawan/log"

	handlerPost "github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler/post"
	repositoryPost "github.com/andiksetyawan/boilerplate_svc/internal/repository/post"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	usecasePost "github.com/andiksetyawan/boilerplate_svc/internal/usecase/post"
	"github.com/andiksetyawan/boilerplate_svc/pkg/httpserver"
	"github.com/andiksetyawan/boilerplate_svc/pkg/opentelemetry"
)

type RestServer struct {
	shutdownFn []func(context.Context) error
	logger     log.Logger

	httpserver    *httpserver.HTTPServer
	opentelemetry *opentelemetry.Otel
}

func NewRestServer(ctx context.Context) (*RestServer, error) {
	rsc := resource.NewResource()
	postRepo := repositoryPost.NewRepository(rsc)
	postUc := usecasePost.NewUsecase(postRepo, rsc)
	postHandler := handlerPost.New(postUc, rsc)
	handlers := handler.NewHandler(postHandler)

	// setup opentelemetry
	var closers []func(context.Context) error
	otel, err := opentelemetry.New(
		opentelemetry.WithServiceName(rsc.Config.ServiceName),
		opentelemetry.WithLogger(rsc.Log),
		opentelemetry.WithJaegerTracerProvider(rsc.Config.OtelJaegerEndpoint),
	)
	if err != nil {
		rsc.Log.Error(ctx, "fail to setup opentelemetry", "error", err)
		return nil, err
	}

	//setup http server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", rsc.Config.ServicePort))
	if err != nil {
		rsc.Log.Error(ctx, "fail to start listener", "error", err)
		return nil, err
	}

	httpServer, err := httpserver.New(
		httpserver.WithTracer(rsc.Config.ServiceName),
		httpserver.WithAppName(rsc.Config.ServiceName),
		httpserver.WithLogger(rsc.Log),
		httpserver.WithListener(listener),
	)
	if err != nil {
		rsc.Log.Error(ctx, "fail to instantiate http server", "error", err)
		return nil, err
	}

	rest.RegisterServer(rsc, httpServer, handlers)

	return &RestServer{
		httpserver:    httpServer,
		opentelemetry: otel,
		logger:        rsc.Log,
		shutdownFn:    closers,
	}, nil
}

func (s *RestServer) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := s.httpserver.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(ctx, "fail to start http server", "error", err)
			return
		}
	}()

	s.opentelemetry.Setup()
	s.shutdownFn = append(s.shutdownFn, s.opentelemetry.Close)
	s.shutdownFn = append(s.shutdownFn, s.httpserver.Close)

	<-ctx.Done()

	s.close(context.TODO())
}

func (s *RestServer) close(ctx context.Context) {
	for _, fn := range s.shutdownFn {
		err := fn(ctx)
		if err != nil {
			s.logger.Error(ctx, "fail to shutdown", "error", err)
		}
	}
}
