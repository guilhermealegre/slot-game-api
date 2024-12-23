package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) CreateUser(ctx ctxDomain.IContext, tx session.ITx, userDetails *v1.CreateUser) (userID int, err error) {
	args := r.Called(ctx, tx, userDetails)
	return args.Get(0).(int), args.Error(1)
}

func (r *RepositoryMock) GetUserDetails(ctx ctxDomain.IContext, userID int) (user *v1.User, err error) {
	args := r.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.User), args.Error(1)
}

func (r *RepositoryMock) UpdateWalletCredits(ctx ctxDomain.IContext, tx session.ITx, userID int, balance float64) (newBalance float64, err error) {
	args := r.Called(ctx, tx, userID, balance)
	return args.Get(0).(float64), args.Error(1)
}

func (r *RepositoryMock) CreateWallet(ctx ctxDomain.IContext, tx session.ITx, userID int) error {
	args := r.Called(ctx, tx, userID)
	return args.Error(0)
}
