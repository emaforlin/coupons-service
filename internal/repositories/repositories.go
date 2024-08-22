package repositories

import (
	"github.com/emaforlin/coupons-service/internal/entities"
)

type CouponsRepository interface {
	UpdateCoupon(id uint, in *entities.InsertCouponDto) error
	InsertCoupon(in *entities.InsertCouponDto) error
	SelectCoupon(in *entities.GetCouponDto) (*entities.Coupon, error)
}
