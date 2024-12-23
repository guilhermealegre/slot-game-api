package v1

import (
	"github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/response"
)

func (t *TokenPair) FromDomainToAPI() *response.AuthResponse {
	if t == nil {
		return nil
	}

	return &response.AuthResponse{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
	}
}
