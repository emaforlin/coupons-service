package entities

type (
	InsertCouponDto struct {
		ID          uint    `gorm:"primaryKey;autoIncrement" json:"coupon_id"`
		Title       string  `gorm:"not null;unique" json:"title"`
		Description string  `gorm:"not null" json:"decription"`
		Multiplier  float32 `gorm:"not null;default:1" json:"multiplier"`
		Active      bool    `gorm:"not null;default:false"`
		Author      int     `gorm:"not null" json:"author"`
	}

	GetCouponDto struct {
		ID          uint
		Title       string
		Description string
		Author      int
	}

	Coupon struct {
		ID          uint    `json:"coupon_id"`
		Title       string  `json:"title"`
		Description string  `json:"decription"`
		Multiplier  float32 `json:"multiplier"`
		Author      int     `json:"author"`
		Active      bool
	}
)

func (InsertCouponDto) TableName() string {
	return "coupons"
}
func (GetCouponDto) TableName() string {
	return "coupons"
}
