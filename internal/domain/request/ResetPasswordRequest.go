package request

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
