package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type IController interface {
	domain.IController
	Profile(gCtx *gin.Context)
	DepositCredits(gCtx *gin.Context)
	WithdrawCredits(gCtx *gin.Context)
}

type IModel interface {
	GetProfile(ctx ctxDomain.IContext, userID int) (*User, error)
	DepositCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error)
	WithdrawCredits(ctx ctxDomain.IContext, userID int, credits float64) (float64, error)
}

type IRepository interface {
	CreateUser(ctx ctxDomain.IContext, tx session.ITx, userDetails *CreateUser) (int, error)
	CreateWallet(ctx ctxDomain.IContext, tx session.ITx, userID int) error
	GetUserDetails(ctx ctxDomain.IContext, userID int) (*User, error)
	UpdateWalletCredits(ctx ctxDomain.IContext, tx session.ITx, userID int, balance float64) (float64, error)
}
