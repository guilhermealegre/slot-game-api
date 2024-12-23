package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilhermealegre/go-clean-arch-core-lib/response"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	"github.com/guilhermealegre/slot-games-api/api/v1/http"
	status "net/http"
	"strings"
)

type AuthenticateMiddleware struct {
	app domain.IApp
}

func NewAuthenticateMiddleware(app domain.IApp) domain.IMiddleware {
	return &AuthenticateMiddleware{
		app: app,
	}
}

func (c *AuthenticateMiddleware) RegisterMiddlewares() {
	http.GroupV1.AddMiddleware(c)
}

func (c *AuthenticateMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		ValidateToken(c.app.Http().Config().JwtSecret),
	}
}

func ValidateToken(jwtSecret string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			_, errResp := response.GetResponse(nil, nil, nil, response.ErrorUnauthorized.Formats("Authorization header is required"))
			c.JSON(status.StatusUnauthorized, errResp)
			c.Abort()
			return
		}

		// Check Bearer scheme
		if !strings.HasPrefix(authHeader, "Bearer ") {
			_, errResp := response.GetResponse(nil, nil, nil, response.ErrorUnauthorized.Formats("Invalid Authorization header format"))
			c.JSON(status.StatusUnauthorized, errResp)
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			_, errResp := response.GetResponse(nil, nil, nil, response.ErrorUnauthorized.Formats(err.Error()))
			c.JSON(status.StatusUnauthorized, errResp)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("user_uuid", claims["user_uuid"])
			c.Set("email", claims["email"])
			c.Set("user_id", claims["user_id"])
		} else {
			_, errResp := response.GetResponse(nil, nil, nil, response.ErrorUnauthorized.Formats("Invalid token claims"))
			c.JSON(status.StatusUnauthorized, errResp)
			c.Abort()
			return
		}
	}
}
