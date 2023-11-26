package request

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}
