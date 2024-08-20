package models

type (
	AddCouponData struct {
		Title       string
		Description string
		Multiplier  float32
		Author      string
	}

	GetCouponData struct {
		ID          uint
		Title       string
		Description string
		Author      string
	}
)
