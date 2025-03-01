package conn

import (
	"fmt"
	"time"

	"github.com/asibulhasanshanto/go_api/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg *config.Config, log *zap.Logger)(*gorm.DB,error) {
	pg := cfg.Postgres
	dsn := pg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get db instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(pg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(pg.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(pg.ConnMaxLifetime) * time.Second)

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info("connected to database", zap.String("dsn", dsn))
	return db, nil
}