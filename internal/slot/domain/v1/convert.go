package v1

import (
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/response"
)

func (smh SpinSlotMachineHistory) FromDomainToApi() response.SpinSlotMachineHistoryResponse {
	if smh == nil {
		return nil
	}

	var history response.SpinSlotMachineHistoryResponse

	for _, ssm := range smh {
		history = append(history, *ssm.FromDomainToApi())
	}

	return history
}

func (ssm *SpinSlotMachine) FromDomainToApi() *response.SpinSlotMachineResponse {
	if ssm == nil {
		return nil
	}

	return &response.SpinSlotMachineResponse{
		SpinResult: ssm.SpinResult.FromDomainToApi(),
		Balance:    ssm.Balance,
		CreatedAt:  ssm.CreatedAt,
	}
}

func (sr *SpinResult) FromDomainToApi() *response.SpinResultResponse {

	if sr == nil {
		return nil
	}

	return &response.SpinResultResponse{
		Symbols:    sr.Symbols,
		BetAmount:  sr.BetAmount,
		Payout:     sr.Payout,
		Winning:    sr.Winning,
		ResultType: *sr.ResultType.FromDomainToApi(),
	}
}

func (srt *SpinResultType) FromDomainToApi() *response.SpinResultTypeResponse {
	if srt == nil {
		return nil
	}

	return &response.SpinResultTypeResponse{
		Id:   srt.Id,
		Key:  srt.Key,
		Name: srt.Name,
	}
}
