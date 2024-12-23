package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) SaveSpinSlotResult(ctx ctxDomain.IContext, tx session.ITx, userID int, result *v1.SpinSlotMachine) error {
	args := m.Called(ctx, tx, userID, result)
	return args.Error(0)
}

func (m *ModelMock) GetSpinResultType(ctx ctxDomain.IContext, key string) (resultType *v1.SpinResultType, err error) {
	args := m.Called(ctx, key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.SpinResultType), args.Error(1)
}

func (m *ModelMock) GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (history v1.SpinSlotMachineHistory, err error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(v1.SpinSlotMachineHistory), args.Error(1)
}
