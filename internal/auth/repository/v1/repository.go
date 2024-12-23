package v1

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal"
	"github.com/guilhermealegre/slot-games-api/internal/auth/domain/v1"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/database"
)

type Repository struct {
	app domain.IApp
}

func NewRepository(app domain.IApp) v1.IRepository {
	return &Repository{
		app: app,
	}
}

func (r *Repository) EmailExist(ctx ctxDomain.IContext, email string) (exist bool, err error) {
	_, err = r.app.Database().Read().
		Select("COUNT(*) > 0").
		From(database.AuthTableAuth).
		Where("LOWER(email) = LOWER(?)", email).
		LoadContext(ctx, &exist)

	if err != nil {
		return false, r.app.Logger().DBLog(err)
	}

	return exist, nil
}

func (r *Repository) CreateAuthentication(ctx ctxDomain.IContext, tx session.ITx, authDetails *v1.CreateAuth) (err error) {
	_, err = tx.InsertInto(database.AuthTableAuth.String()).
		Columns(
			"user_fk",
			"email",
			"password",
		).
		Values(
			authDetails.UserID,
			authDetails.Email,
			authDetails.Password,
		).
		ExecContext(ctx)

	if err != nil {
		return r.app.Logger().DBLog(err)
	}

	return nil
}

func (r *Repository) GetAuthDetailsByEmail(ctx ctxDomain.IContext, email string) (authDetails *v1.Auth, err error) {
	count, err := r.app.Database().Read().
		Select("user_fk as user_id",
			"email",
			"password").
		From(database.AuthTableAuth).
		Where("LOWER(email) = LOWER(?)", email).
		LoadContext(ctx, &authDetails)

	if err != nil {
		return nil, r.app.Logger().DBLog(err)
	}

	if count == 0 {
		return nil, internal.ErrorGeneric().Formats("Email not found")
	}

	return authDetails, nil
}
