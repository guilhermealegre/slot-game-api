package v1

import "time"

const (
	Email          = "email"
	UserUUID       = "user_uuid"
	UserID         = "user_id"
	ExpirationTime = "exp"

	// token ttl
	AccessTokenTTL  = time.Minute * 15
	RefreshTokenTTL = time.Hour * 24 * 30
)
