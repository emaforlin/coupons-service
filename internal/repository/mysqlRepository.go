package repository

import (
	"errors"

	"github.com/emaforlin/coupons-service/internal/database"
	"github.com/emaforlin/coupons-service/internal/entities"
	"github.com/emaforlin/coupons-service/pkg/models"
	"gorm.io/gorm"
)

type couponsMysqlRepository struct {
	db database.Database
}

// SelectCoupon implements CouponsRepository.
func (c *couponsMysqlRepository) SelectCoupon(in *models.GetCouponData) (*entities.Coupon, error) {
	panic("unimplemented")
}

// InsertCoupon implements CouponsRepository.
func (c *couponsMysqlRepository) InsertCoupon(in *models.AddCouponData) error {
	res := c.db.GetDb().Create(in)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("coupon already exists")
		}
	}
	return nil
}

func NewCouponsRepository(d database.Database) CouponsRepository {
	return &couponsMysqlRepository{
		db: d,
	}
}
