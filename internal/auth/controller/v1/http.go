package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Routes "github.com/guilhermealegre/slot-games-api/api/v1/http"
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/request"
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/response"
	v1Domain "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	v1 "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
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
	v1Routes.Register.SetRoute(engine, c.Signup)
	v1Routes.Login.SetRoute(engine, c.Login)

}

/*
swagger:route POST /p/login auth LoginRequest

	 Login user

		Produces:
		- application/json

		Responses:
		  200: SwaggerScannerLoginResponse
		  400: ErrorResponse
*/
func (c *Controller) Login(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	var req request.LoginRequest

	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	if err := c.App().Validator().Validate(ctx, req); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	obj, err := c.model.Login(ctx, req.Body.Email, req.Body.Password)
	if err != nil {
		c.Json(ctx, nil, err)
		return
	}

	c.Json(ctx, obj.FromDomainToAPI(), err)
}

/*
	 swagger:route POST /p/register auth SignupRequest

	 Signup  user

		Produces:
		- application/json

		Responses:
		  200: SuccessResponse
		  400: ErrorResponse
*/
func (c *Controller) Signup(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	req := request.SignupRequest{}
	if err := ctx.ShouldBindJSON(&req.Body); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	if err := c.App().Validator().Validate(ctx, req); err != nil {
		c.Json(ctx, nil, err)
		return
	}

	err := c.model.Signup(ctx,
		&v1.CreateUser{
			FirstName: req.Body.FirstName,
			LastName:  req.Body.LastName,
			Avatar:    req.Body.Avatar,
		},
		&v1Domain.CreateAuth{
			Email:    req.Body.Email,
			Password: req.Body.Password,
		})

	c.Json(ctx, response.SuccessResponse{Success: err == nil}, err)
}
