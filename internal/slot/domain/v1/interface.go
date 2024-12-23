package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
)

type IController interface {
	domain.IController
	SpinSlotMachine(gCtx *gin.Context)
	GetSpinSlotHistory(gCtx *gin.Context)
}

type IModel interface {
	SpinSlotMachine(ctx ctxDomain.IContext, BetAmount float64, userID int) (spinSlotMachine *SpinSlotMachine, err error)
	GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (SpinSlotMachineHistory, error)
}

type IRepository interface {
	SaveSpinSlotResult(ctx ctxDomain.IContext, tx session.ITx, userID int, result *SpinSlotMachine) error
	GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (SpinSlotMachineHistory, error)
	GetSpinResultType(ctx ctxDomain.IContext, key string) (*SpinResultType, error)
}
