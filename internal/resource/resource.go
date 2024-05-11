package resource

import (
	"github.com/andiksetyawan/boilerplate_svc/internal/config"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource/client/rest/user_client"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource/httpresponse"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource/logger"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource/psql"
	"github.com/andiksetyawan/boilerplate_svc/pkg/response"
	"github.com/andiksetyawan/database/sqlx"
	"github.com/andiksetyawan/log"
)

type Resource struct {
	Config  config.Config
	Log     log.Logger
	DB      sqlx.DB
	HttpRes *response.HttpResponse

	userClient user_client.UserClient
}

func NewResource() Resource {
	cfg := config.NewConfig()
	slog := logger.NewSlog(cfg)
	db := psql.NewSqlx(cfg, slog)
	httpRes := httpresponse.NewHttpResponse(slog)

	return Resource{
		Config:  cfg,
		Log:     slog,
		DB:      db,
		HttpRes: httpRes,

		userClient: nil,
	}
}
