package v1

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type CreateAuth struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenJWTDetails struct {
	UserUUID string `json:"user_uuid"`
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
}
