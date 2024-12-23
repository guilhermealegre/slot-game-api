package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type Model struct {
	app        domain.IApp
	repository v1.IRepository
}

func NewModel(
	app domain.IApp,
	repository v1.IRepository) v1.IModel {
	return &Model{
		app:        app,
		repository: repository,
	}
}

func (m *Model) GetProfile(ctx ctxDomain.IContext, userID int) (*v1.User, error) {
	return m.repository.GetUserDetails(ctx, userID)
}

func (m *Model) DepositCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error) {
	userDetails, err := m.GetProfile(ctx, userID)
	if err != nil {
		return 0, err
	}

	userDetails.Wallet.Balance += credits
	return m.repository.UpdateWalletCredits(ctx, nil, userID, userDetails.Wallet.Balance)
}

func (m *Model) WithdrawCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error) {
	userDetails, err := m.GetProfile(ctx, userID)
	if err != nil {
		return 0, err
	}

	if userDetails.Wallet.Balance-credits < 0 {
		return 0, internal.ErrInsufficientFunds()
	}

	userDetails.Wallet.Balance -= credits
	return m.repository.UpdateWalletCredits(ctx, nil, userID, userDetails.Wallet.Balance)
}
