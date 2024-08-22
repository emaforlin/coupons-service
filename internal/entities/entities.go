package entities

type (
	InsertCouponDto struct {
		ID            uint    `gorm:"primaryKey;autoIncrement"`
		Title         string  `gorm:"not null;unique"`
		Description   string  `gorm:"not null"`
		Multiplier    float32 `gorm:"not null;default:1"`
		State         string  `gorm:"not null;default:disabled"`
		BussinessName string  `gorm:"not null"`
	}

	GetCouponDto struct {
		ID            uint
		Title         string
		Description   string
		BussinessName string
	}

	Coupon struct {
		ID            uint    `json:"coupon_id"`
		Title         string  `json:"title"`
		Description   string  `json:"decription"`
		Multiplier    float32 `json:"multiplier"`
		State         string  `json:"state"`
		BussinessName string  `json:"bussiness_name"`
	}
)

func (InsertCouponDto) TableName() string {
	return "coupons"
}
func (GetCouponDto) TableName() string {
	return "coupons"
}
