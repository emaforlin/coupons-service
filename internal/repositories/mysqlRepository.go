package repositories

import (
	"errors"
	"fmt"

	"github.com/emaforlin/coupons-service/internal/database"
	"github.com/emaforlin/coupons-service/internal/entities"
	"gorm.io/gorm"
)

type couponsMysqlRepository struct {
	db database.Database
}

// UpdateCoupon implements CouponsRepository.
func (r *couponsMysqlRepository) UpdateCoupon(id uint, in *entities.InsertCouponDto) error {
	err := r.db.GetDb().Model(&entities.InsertCouponDto{}).Where("id = ?", id).Updates(in).Error
	if err != nil {
		fmt.Printf("DEBUG: %v\n", err)
		return err
	}
	return nil
}

// SelectCoupon implements CouponsRepository.
func (r *couponsMysqlRepository) SelectCoupon(in *entities.GetCouponDto) (*entities.Coupon, error) {
	found := new(entities.Coupon)
	err := r.db.GetDb().Model(entities.GetCouponDto{}).First(&found, in).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("coupon not found")
		}
	}

	return found, nil
}

// InsertCoupon implements CouponsRepository.
func (r *couponsMysqlRepository) InsertCoupon(in *entities.InsertCouponDto) error {
	res := r.db.GetDb().Create(in)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("coupon already exists")
		}
		return errors.New("something went wrong creating new coupon")
	}
	return nil
}

func NewCouponsRepository(d database.Database) CouponsRepository {
	return &couponsMysqlRepository{
		db: d,
	}
}
