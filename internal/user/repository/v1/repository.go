package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/database"
	v1 "github.com/guilhermealegre/slot-games-api/internal/user/domain/v1"
)

type Repository struct {
	app domain.IApp
}

func NewRepository(app domain.IApp) v1.IRepository {
	return &Repository{
		app: app,
	}
}

func (r *Repository) CreateUser(ctx ctxDomain.IContext, tx session.ITx, userDetails *v1.CreateUser) (userID int, err error) {
	err = tx.InsertInto(database.UserTableUser.String()).
		Columns(
			"uuid",
			"first_name",
			"last_name",
			"avatar",
		).
		Values(
			userDetails.UserUUID,
			userDetails.FirstName,
			userDetails.LastName,
			userDetails.Avatar,
		).Returning("user_id").
		LoadContext(ctx, &userID)

	if err != nil {
		return 0, r.app.Logger().DBLog(err)
	}

	return userID, nil
}

func (r *Repository) GetUserDetails(ctx ctxDomain.IContext, userID int) (user *v1.User, err error) {
	_, err = r.app.Database().Read().
		Select(
			"u.user_id",
			"u.uuid",
			"u.first_name",
			"u.last_name",
			"u.avatar",
			"a.email",
			"w.wallet_id",
			"w.balance").
		From(database.UserTableUser.As("u")).
		Join(database.AuthTableAuth.As("a"), "u.user_id = a.user_fk").
		Join(database.UserTableWallet.As("w"), "u.user_id = w.user_fk").
		Where("user_id = ?", userID).
		LoadContext(ctx, &user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) UpdateWalletCredits(ctx ctxDomain.IContext, tx session.ITx, userID int, balance float64) (newBalance float64, err error) {
	var commit bool
	if tx == nil {
		tx, err = r.app.Database().Write().Begin()
		if err != nil {
			return 0, r.app.Logger().DBLog(err)
		}
		commit = true
		defer tx.RollbackUnlessCommitted()
	}

	err = tx.Update(database.UserTableWallet.String()).
		Set("balance", balance).
		Where("user_fk = ?", userID).
		Returning("balance").
		LoadContext(ctx, &newBalance)

	if err != nil {
		return 0, r.app.Logger().DBLog(err)
	}

	if commit {
		if err = tx.Commit(); err != nil {
			return 0, r.app.Logger().DBLog(err)
		}
	}

	return newBalance, nil
}

func (r *Repository) CreateWallet(ctx ctxDomain.IContext, tx session.ITx, userID int) error {
	_, err := tx.InsertInto(database.UserTableWallet.String()).
		Columns(
			"user_fk",
		).
		Values(
			userID,
		).
		ExecContext(ctx)

	if err != nil {
		return r.app.Logger().DBLog(err)
	}

	return nil
}
