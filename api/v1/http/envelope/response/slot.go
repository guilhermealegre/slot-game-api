package response

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/response"
	"time"
)

// swagger:model SwaggerSpinSlotMachineResponse
type swaggerSpinSlotMachineResponse struct { //nolint:all
	response.Response
	Data SpinSlotMachineResponse `json:"data"`
}

//swagger:model ProfileResponse
type SpinSlotMachineResponse struct {
	SpinResult *SpinResultResponse `json:"spin_result"`
	Balance    float64             `json:"balance"`
	CreatedAt  *time.Time          `json:"created_at"`
}

//swagger:model SpinResultResponse
type SpinResultResponse struct {
	Symbols    []int32                `json:"symbols"`
	BetAmount  float64                `json:"bet_amount"`
	Payout     float64                `json:"payout"`
	Winning    bool                   `json:"winning"`
	ResultType SpinResultTypeResponse `json:"result_type"`
}

//swagger:model SpinResultTypeResponse
type SpinResultTypeResponse struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// swagger:model SwaggerSpinSlotMachineHistoryResponse
type swaggerSpinSlotMachineHistoryResponse struct { //nolint:all
	response.Response
	Data SpinSlotMachineHistoryResponse `json:"data"`
}

type SpinSlotMachineHistoryResponse []SpinSlotMachineResponse
