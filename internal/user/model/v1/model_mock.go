package v1

import (
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) GetProfile(ctx ctxDomain.IContext, userID int) (*v1.User, error) {
	args := m.Called(ctx, userID)
	if args.Error(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.User), args.Error(1)
}

func (m *ModelMock) DepositCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error) {
	args := m.Called(ctx, userID, credits)
	return args.Get(0).(float64), args.Error(1)
}

func (m *ModelMock) WithdrawCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error) {
	args := m.Called(ctx, userID, credits)
	return args.Get(0).(float64), args.Error(1)
}
