package request

//swagger:parameters SpinSlotRequest
type SpinSlotRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		// Bet Amount
		// required: true
		BetAmount float64 `json:"bet_amount" validate:"required,gt=0"`
	}
}
