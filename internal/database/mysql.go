package database

import (
	"fmt"

	"github.com/emaforlin/coupons-service/internal/config"
	"github.com/emaforlin/coupons-service/internal/entities"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDatabase struct {
	db *gorm.DB
}

// AutoMigrate implements Database.
func (p *mysqlDatabase) AutoMigrate() error {
	return p.db.AutoMigrate(entities.InsertCouponDto{})
}

func NewMySQLDatabase(cfg *config.Config, l hclog.Logger) Database {
	dblogger := logger.New(l.Named("database").StandardLogger(&hclog.StandardLoggerOptions{}), logger.Config{
		LogLevel: logger.Error,
	})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Passwd, cfg.Db.Host, cfg.Db.Name),
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dblogger.LogMode(logger.Silent),
	})

	if err != nil {
		l.Error("Database connection failed")
		panic(err)
	}
	return &mysqlDatabase{db: db}
}

func (p *mysqlDatabase) GetDb() *gorm.DB {
	return p.db
}
