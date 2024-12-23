package v1

type User struct {
	UserID    int    `json:"user_id"`
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Wallet    Wallet `json:"wallet"`
}

type Wallet struct {
	WalletID int     `json:"wallet_id"`
	Balance  float64 `json:"balance"`
}

type CreateUser struct {
	UserUUID  string `json:"user_uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}
