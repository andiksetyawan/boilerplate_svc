package psql

import (
	"context"
	"os"

	"github.com/andiksetyawan/boilerplate_svc/internal/config"
	"github.com/andiksetyawan/database"
	"github.com/andiksetyawan/database/sqlx"
	"github.com/andiksetyawan/log"
)

func NewSqlx(config config.Config, log log.Logger) sqlx.DB {
	db, err := sqlx.New(sqlx.WithPostgres(database.Config{
		Database: config.DBName,
		User:     config.DBUsername,
		Host:     config.DBHost,
		Port:     config.DBPort,
		Password: config.DBPassword,
	}))
	if err != nil {
		log.Error(context.TODO(), "fail to connecting psql database", "error", err)
		os.Exit(1)
	}

	return db
}
