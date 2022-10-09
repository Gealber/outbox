package database

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/Gealber/outbox/config"
	_ "github.com/newrelic/go-agent/_integrations/nrpq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(ctx context.Context, cfg *config.AppConfig) (*gorm.DB, error) {
	rand.Seed(time.Now().UnixNano())

	dsn := cfg.Cockroach.DSN

	conn, err := sql.Open("nrpostgres", dsn)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
