package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/database"
	"github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	"github.com/lib/pq"
)

type Model struct {
	app              domain.IApp
	repo             v1.IRepository
	userRepo         v1UserDomain.IRepository
	symbolsGenerator helpers.IRandomSymbols
}

func NewModel(
	app domain.IApp,
	repo v1.IRepository,
	userRepo v1UserDomain.IRepository,
	symbolsGenerator helpers.IRandomSymbols) v1.IModel {
	return &Model{
		app:              app,
		repo:             repo,
		userRepo:         userRepo,
		symbolsGenerator: symbolsGenerator,
	}
}

func (m *Model) SpinSlotMachine(ctx ctxDomain.IContext, BetAmount float64, userID int) (spinSlotMachine *v1.SpinSlotMachine, err error) {
	spinSlotMachine = &v1.SpinSlotMachine{
		SpinResult: v1.SpinResult{
			Symbols: make(pq.Int32Array, 3),
		},
	}

	details, err := m.userRepo.GetUserDetails(ctx, userID)
	if err != nil {
		return nil, err
	}

	spinSlotMachine.Balance = details.Wallet.Balance - BetAmount
	if spinSlotMachine.Balance < 0 {
		return nil, internal.ErrInsufficientFunds()
	}

	spinSlotMachine.SpinResult.Symbols = m.symbolsGenerator.Generate()
	spinResult := m.calculatePayout(spinSlotMachine.SpinResult.Symbols, BetAmount)
	spinSlotMachine.SpinResult = *spinResult
	spinSlotMachine.Balance += spinResult.Payout

	resultType, err := m.repo.GetSpinResultType(ctx, spinResult.ResultType.Key)
	if err != nil {
		return nil, err
	}

	spinSlotMachine.SpinResult.ResultType = *resultType
	tx, err := m.app.Database().Write().Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		return nil, m.app.Logger().DBLog(err)
	}

	// Update user balance
	_, err = m.userRepo.UpdateWalletCredits(ctx, tx, userID, spinSlotMachine.Balance)
	if err != nil {
		return nil, err
	}

	// Save game history
	err = m.repo.SaveSpinSlotResult(ctx, tx, userID, spinSlotMachine)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, m.app.Logger().DBLog(err)
	}

	return spinSlotMachine, nil

}

func (m *Model) calculatePayout(symbols pq.Int32Array, betAmount float64) *v1.SpinResult {
	result := &v1.SpinResult{
		Symbols:   symbols,
		BetAmount: betAmount,
	}

	symbolCounts := make(map[int32]int32)
	for _, symbol := range symbols {
		symbolCounts[symbol]++
	}

	maxCount := int32(0)
	for _, count := range symbolCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	switch maxCount {
	case 3:
		result.Payout = betAmount * 10
		result.Winning = true
		result.ResultType.Key = database.SpinResultTypeX10Key
	case 2:
		result.Payout = betAmount * 2
		result.Winning = true
		result.ResultType.Key = database.SpinResultTypeX2Key
	default:
		result.Payout = 0
		result.Winning = false
		result.ResultType.Key = database.SpinResultTypeLossKey
	}

	return result
}

func (m *Model) GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (v1.SpinSlotMachineHistory, error) {
	return m.repo.GetSpinSlotHistory(ctx, userID)
}
