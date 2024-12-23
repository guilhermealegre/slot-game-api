package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	v1UserDomain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type IController interface {
	domain.IController
	Login(gCtx *gin.Context)
	Signup(gCtx *gin.Context)
}

type IModel interface {
	Login(ctx ctxDomain.IContext, email, password string) (*TokenPair, error)
	Signup(ctx ctxDomain.IContext, userDetails *v1UserDomain.CreateUser, authDetails *CreateAuth) error
}

type IRepository interface {
	EmailExist(ctx ctxDomain.IContext, email string) (bool, error)
	CreateAuthentication(ctx ctxDomain.IContext, tx session.ITx, authDetails *CreateAuth) error
	GetAuthDetailsByEmail(ctx ctxDomain.IContext, email string) (*Auth, error)
}
