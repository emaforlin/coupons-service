package repository

import (
	"github.com/emaforlin/coupons-service/internal/entities"
	"github.com/emaforlin/coupons-service/pkg/models"
)

type CouponsRepository interface {
	InsertCoupon(in *models.AddCouponData) error
	SelectCoupon(in *models.GetCouponData) (*entities.Coupon, error)
}
