package models

// REMEMBER: Add required struct tags

type (
	ChangeCouponState struct {
		ID    uint
		State string
	}
	AddCouponData struct {
		Title         string
		Description   string
		Multiplier    float32
		BussinessName string
		State         string
	}

	GetCouponData struct {
		ID            uint
		Title         string
		Description   string
		BussinessName string
	}
)
