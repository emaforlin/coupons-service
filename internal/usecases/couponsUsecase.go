package usecases

import (
	"github.com/emaforlin/coupons-service/internal/entities"
	"github.com/emaforlin/coupons-service/internal/repositories"
	"github.com/emaforlin/coupons-service/pkg/models"
)

type couponsUsecaseImpl struct {
	repository repositories.CouponsRepository
}

// ChangeState implements CouponUsecase.
func (u *couponsUsecaseImpl) ChangeState(in *models.ChangeCouponState) error {
	err := u.repository.UpdateCoupon(in.ID, &entities.InsertCouponDto{
		State: in.State,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *couponsUsecaseImpl) CreateNewCoupon(in *models.AddCouponData) error {
	err := u.repository.InsertCoupon(&entities.InsertCouponDto{
		Title:         in.Title,
		Description:   in.Description,
		Multiplier:    in.Multiplier,
		BussinessName: in.BussinessName,
		State:         in.State,
	})

	if err != nil {
		return err
	}
	return nil
}

func NewCouponsUsecase(rep repositories.CouponsRepository) CouponUsecase {
	return &couponsUsecaseImpl{
		repository: rep,
	}
}
