package schema

type RegisterBody struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type LoginSchema struct {
	Phone string `json:"phone"`
}

type VerifyOTPSchema struct {
	Phone string `json:"phone"`
	Otp   string `json:"otp"`
}
