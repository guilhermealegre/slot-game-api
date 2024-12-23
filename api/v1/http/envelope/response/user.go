package response

import "github.com/guilhermealegre/go-clean-arch-core-lib/response"

// swagger:model SwaggerProfileResponse
type swaggerProfileResponse struct { //nolint:all
	response.Response
	Data ProfileResponse `json:"data"`
}

//swagger:model ProfileResponse
type ProfileResponse struct {
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Avatar    string          `json:"avatar"`
	Email     string          `json:"email"`
	Wallet    *WalletResponse `json:"wallet"`
}

//swagger:model WalletResponse
type WalletResponse struct {
	Balance float64 `json:"balance"`
}

// swagger:model SwaggerDepositCreditsResponse
type swaggerDepositCreditsResponse struct { //nolint:all
	response.Response
	Data DepositCreditsResponse `json:"data"`
}

//swagger:parameters DepositCreditsResponse
type DepositCreditsResponse struct {
	Balance float64 `json:"balance"`
}

// swagger:model SwaggerWithdrawCreditsResponse
type SwaggerWithdrawCreditsResponse struct { //nolint:all
	response.Response
	Data WithdrawCreditsResponse `json:"data"`
}

//swagger:parameters WithdrawCreditsResponse
type WithdrawCreditsResponse struct {
	Balance float64 `json:"balance"`
}
