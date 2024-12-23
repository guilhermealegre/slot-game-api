package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) SaveSpinSlotResult(ctx ctxDomain.IContext, tx session.ITx, userID int, result *v1.SpinSlotMachine) error {
	args := r.Called(ctx, tx, userID, result)
	return args.Error(0)
}

func (r *RepositoryMock) GetSpinResultType(ctx ctxDomain.IContext, key string) (*v1.SpinResultType, error) {
	args := r.Called(ctx, key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.SpinResultType), args.Error(1)
}

func (r *RepositoryMock) GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (history v1.SpinSlotMachineHistory, err error) {
	args := r.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(v1.SpinSlotMachineHistory), args.Error(1)
}
