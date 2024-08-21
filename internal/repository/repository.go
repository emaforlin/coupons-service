package repository

import (
	"github.com/emaforlin/coupons-service/internal/entities"
)

type CouponsRepository interface {
	InsertCoupon(in *entities.InsertCouponDto) error
	SelectCoupon(in *entities.GetCouponDto) (*entities.Coupon, error)
}
