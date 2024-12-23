package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Routes "github.com/guilhermealegre/slot-games-api/api/v1/http"
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/request"
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/response"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	v1Domain "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
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
	v1Routes.Profile.SetRoute(engine, c.Profile)
	v1Routes.DepositCredits.SetRoute(engine, c.DepositCredits)
	v1Routes.WithdrawCredits.SetRoute(engine, c.WithdrawCredits)

}

/*
	 swagger:route GET /profile user user

	 Ger user profile

		Produces:
		- application/json

		Security:
	  		Bearer:

		Responses:
		  200: SwaggerProfileResponse
		  400: ErrorResponse
*/
func (c *Controller) Profile(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	userID, err := helpers.GetJwtUserID(gCtx)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	obj, err := c.model.GetProfile(ctx, userID)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	c.Json(ctx, obj.FromDomainToAPI(), err)
}

/*
	 swagger:route POST /wallet/deposit user DepositCreditsRequest

	 Deposit credits

		Produces:
		- application/json

		Security:
	  		Bearer:

		Responses:
		  200: SwaggerDepositCreditsResponse
		  400: ErrorResponse
*/
func (c *Controller) DepositCredits(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	req := request.DepositCreditsRequest{}
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

	balance, err := c.model.DepositCredits(ctx, userID, req.Body.Credits)
	c.Json(ctx, response.DepositCreditsResponse{Balance: balance}, err)
}

/*
	 swagger:route POST /wallet/withdraw user WithdrawCreditsRequest

	 Withdraw credits

		Produces:
		- application/json

		Security:
	  		Bearer:

		Responses:
		  200: SwaggerWithdrawCreditsResponse
		  400: ErrorResponse
*/
func (c *Controller) WithdrawCredits(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	req := request.WithdrawCreditsRequest{}
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

	balance, err := c.model.WithdrawCredits(ctx, userID, req.Body.Credits)
	c.Json(ctx, response.WithdrawCreditsResponse{Balance: balance}, err)
}
