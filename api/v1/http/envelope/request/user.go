package request

//swagger:parameters DepositCreditsRequest
type DepositCreditsRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		Credits float64 `json:"credits" validate:"required,gt=0"`
	}
}

//swagger:parameters WithdrawCreditsRequest
type WithdrawCreditsRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		Credits float64 `json:"credits" validate:"required,gt=0"`
	}
}
