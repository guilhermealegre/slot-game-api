package v1

import "time"

const (
	Email          = "email"
	UserUUID       = "user_uuid"
	ExpirationTime = "exp"

	// token ttl
	AccessTokenTTL  = time.Minute * 15
	RefreshTokenTTL = time.Hour * 24 * 30

	SecretKey = "secret_slot_games"
)
