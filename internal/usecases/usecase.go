package usecases

import "github.com/emaforlin/coupons-service/pkg/models"

type CouponUsecase interface {
	CreateNewCoupon(in *models.AddCouponData) error
	ChangeState(in *models.ChangeCouponState) error
}
