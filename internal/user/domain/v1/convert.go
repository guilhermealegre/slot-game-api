package v1

import "github.com/guilhermealegre/slot-games-api/api/v1/http/envelope/response"

func (u *User) FromDomainToAPI() *response.ProfileResponse {
	if u == nil {
		return nil
	}
	return &response.ProfileResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Wallet:    u.Wallet.FromDomainToAPI(),
	}
}

func (w *Wallet) FromDomainToAPI() *response.WalletResponse {
	if w == nil {
		return nil
	}

	return &response.WalletResponse{
		Balance: w.Balance,
	}
}
