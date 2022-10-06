package database

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/newrelic/go-agent/_integrations/nrpq"
	"github.com/Gealber/outbox/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(ctx context.Context, cfg *config.AppConfig) (*gorm.DB, error) {
	rand.Seed(time.Now().UnixNano())

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SslMode,
		cfg.Database.Timezone,
	)

	conn, err := sql.Open("nrpostgres", dsn)
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
