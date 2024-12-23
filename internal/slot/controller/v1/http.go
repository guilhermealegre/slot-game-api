package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Routes "github.com/guilhermealegre/slot-games-api/api/v1/http"
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/request"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	v1Domain "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
)

type Controller struct {
	*domain.DefaultController
	model v1Domain.IModel
}

func NewController(app domain.IApp, model v1Domain.IModel) v1Domain.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	v1Routes.SpinSlotMachine.SetRoute(engine, c.SpinSlotMachine)
	v1Routes.SpinSlotHistory.SetRoute(engine, c.GetSpinSlotHistory)

}

/*
	 swagger:route POST /slot/spin slot SpinSlotRequest

	 Spin slot machine

		Produces:
		- application/json

		Security:
		  Bearer:

		Responses:
		  200: SwaggerSpinSlotMachineResponse
		  400: ErrorResponse
*/
func (c *Controller) SpinSlotMachine(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	var req request.SpinSlotRequest

	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	if err := c.App().Validator().Validate(ctx, req); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	userID, err := helpers.GetJwtUserID(gCtx)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	obj, err := c.model.SpinSlotMachine(ctx, req.Body.BetAmount, userID)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	c.Json(ctx, obj.FromDomainToApi(), err)
}

/*
	 swagger:route GET /slot/history slot slot

	 Get slot history

		Produces:
		- application/json

		Security:
	  		Bearer:

		Responses:
			200: SwaggerSpinSlotMachineHistoryResponse
	  		400: ErrorResponse
*/
func (c *Controller) GetSpinSlotHistory(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	userID, err := helpers.GetJwtUserID(gCtx)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	obj, err := c.model.GetSpinSlotHistory(ctx, userID)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	c.Json(ctx, obj.FromDomainToApi(), err)
}
