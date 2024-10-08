package main

import (
	"log"

	"github.com/emaforlin/coupons-service/internal/config"
	"github.com/emaforlin/coupons-service/internal/database"
	"github.com/emaforlin/coupons-service/internal/entities"
	"github.com/emaforlin/coupons-service/internal/repositories"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.FromStandardLogger(log.Default(), &hclog.LoggerOptions{
		Name:  "coupons-service",
		Level: hclog.Debug,
	})
	logger.Info("Load configurations")

	config.InitViper("config")
	conf := config.LoadConfig()

	logger.Info("Connect to database")
	db := database.NewMySQLDatabase(conf, logger)

	if err := db.AutoMigrate(); err != nil {
		logger.Error("could not migrate database", err)
	}
	logger.Info("successfully migrated database")

	repo := repositories.NewCouponsRepository(db)
	repo.InsertCoupon(&entities.InsertCouponDto{
		Title:         "Otro cupon",
		Description:   "Descripcion del cupon",
		BussinessName: "No se flaco",
	})

}
