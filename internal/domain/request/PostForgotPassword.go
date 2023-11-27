package request

type ForgotPassword struct {
	Email string `json:"email" validate:"required"`
}
