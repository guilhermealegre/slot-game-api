package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) EmailExist(ctx ctxDomain.IContext, email string) (exist bool, err error) {
	args := r.Called(ctx, email)
	return args.Get(0).(bool), args.Error(1)
}

func (r *RepositoryMock) CreateAuthentication(ctx ctxDomain.IContext, tx session.ITx, authDetails *v1.CreateAuth) error {
	args := r.Called(ctx, tx, authDetails)
	return args.Error(0)
}

func (r *RepositoryMock) GetAuthDetailsByEmail(ctx ctxDomain.IContext, email string) (authDetails *v1.Auth, err error) {
	args := r.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.Auth), args.Error(1)
}
