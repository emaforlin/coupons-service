package repository

import (
	"errors"

	"github.com/emaforlin/coupons-service/internal/database"
	"github.com/emaforlin/coupons-service/internal/entities"
	"gorm.io/gorm"
)

type couponsMysqlRepository struct {
	db database.Database
}

// SelectCoupon implements CouponsRepository.
func (c *couponsMysqlRepository) SelectCoupon(in *entities.GetCouponDto) (*entities.Coupon, error) {
	found := new(entities.InsertCouponDto)
	err := c.db.GetDb().Model(entities.GetCouponDto{}).First(&found, in).Error
	if err != nil {
		return nil, errors.New("coupon not found")
	}
	return &entities.Coupon{
		ID:          found.ID,
		Title:       found.Title,
		Description: found.Description,
		Author:      found.Author,
		Multiplier:  found.Multiplier,
		Active:      found.Active,
	}, nil
}

// InsertCoupon implements CouponsRepository.
func (c *couponsMysqlRepository) InsertCoupon(in *entities.InsertCouponDto) error {
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
