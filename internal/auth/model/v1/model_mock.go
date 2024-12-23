package v1

import (
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1 "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) Login(ctx ctxDomain.IContext, email, password string) (tokenPair *v1.TokenPair, err error) {
	args := m.Called(ctx, email, password)
	if args.Error(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*v1.TokenPair), args.Error(1)
}

func (m *ModelMock) Signup(ctx ctxDomain.IContext, userDetails *v1UserDomain.CreateUser, authDetails *v1.CreateAuth) error {
	args := m.Called(ctx, userDetails, authDetails)
	return args.Error(1)
}
