package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/response"
	"github.com/guilhermealegre/slot-games-api/api/v1/http"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/helpers"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/rate_limiter"
	status "net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type SpinRateLimiterMiddleware struct {
	app         domain.IApp
	rateLimiter *rate_limiter.RateLimiter
}

func NewSpinRateLimiterMiddleware(app domain.IApp, rateLimiter *rate_limiter.RateLimiter) domain.IMiddleware {
	return &SpinRateLimiterMiddleware{
		app:         app,
		rateLimiter: rateLimiter,
	}
}

func (c *SpinRateLimiterMiddleware) RegisterMiddlewares() {
	http.SpinSlotMachine.AddMiddlewares(c)
}

func (c *SpinRateLimiterMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		c.RateLimit,
	}
}

func (c *SpinRateLimiterMiddleware) RateLimit(gCtx *gin.Context) {

	id, err := helpers.GetJwtUserID(gCtx)
	if err != nil {
		_, errResp := response.GetResponse(nil, nil, nil, internal.ErrorUserIDNotFound())
		gCtx.JSON(status.StatusUnauthorized, errResp)
		gCtx.Abort()
		return
	}

	if !c.rateLimiter.Allow(id) {
		_, errResp := response.GetResponse(nil, nil, nil, internal.ErrorRateLimitExceeded())
		gCtx.JSON(status.StatusTooManyRequests, errResp)
		gCtx.Abort()
		return
	}

	gCtx.Next()
}
