package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/slot-games-api/internal"
	v1DomainAuth "github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
)

func GetJwtUserID(ctx *gin.Context) (int, error) {
	jwtUserID := ctx.GetFloat64(v1DomainAuth.UserID)

	userID := int(jwtUserID)

	if userID <= 0 {
		return 0, internal.ErrorUserIDNotFound()
	}

	return userID, nil
}
