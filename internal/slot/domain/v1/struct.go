package v1

import (
	"github.com/lib/pq"
	"time"
)

type SpinSlotMachine struct {
	SpinResult SpinResult `json:"spin_result"`
	Balance    float64    `json:"balance"`
	CreatedAt  *time.Time `json:"created_at"`
}

type SpinResult struct {
	Symbols    pq.Int32Array  `json:"symbols"`
	BetAmount  float64        `json:"bet_amount"`
	Payout     float64        `json:"payout"`
	Winning    bool           `json:"winning"`
	ResultType SpinResultType `json:"result_type"`
}

type SpinResultType struct {
	Id   int    `json:"id" db:"spin_result_type_id"`
	Key  string `json:"key" db:"spin_result_type_key"`
	Name string `json:"name" db:"spin_result_type_name"`
}

type SpinSlotMachineHistory []SpinSlotMachine
