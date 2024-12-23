package v1

import (
	"github.com/gocraft/dbr/v2"
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/session"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	ctxDomain "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"github.com/guilhermealegre/slot-games-api/internal/infrastructure/database"
	v1 "github.com/guilhermealegre/slot-games-api/internal/slot/domain/v1"
	"github.com/lib/pq"
)

type Repository struct {
	app domain.IApp
}

func NewRepository(app domain.IApp) v1.IRepository {
	return &Repository{
		app: app,
	}
}

func (r *Repository) SaveSpinSlotResult(ctx ctxDomain.IContext, tx session.ITx, userID int, result *v1.SpinSlotMachine) error {
	var pqSymbols pq.Int32Array
	for _, symbols := range result.SpinResult.Symbols {
		pqSymbols = append(pqSymbols, int32(symbols))
	}

	_, err := tx.InsertInto(database.SlotTableSpin.String()).
		Columns(
			"user_fk",
			"bet",
			"payout",
			"symbols",
			"winning",
			"spin_result_type_fk",
			"user_balance",
		).
		Values(
			userID,
			result.SpinResult.BetAmount,
			result.SpinResult.Payout,
			pqSymbols,
			result.SpinResult.Winning,
			dbr.Select("spin_result_type_id").From(database.SlotTableSpinResultType.String()).Where("key = ?", result.SpinResult.ResultType.Key),
			result.Balance,
		).
		ExecContext(ctx)

	if err != nil {
		return r.app.Logger().DBLog(err)
	}
	return nil
}

func (r *Repository) GetSpinSlotHistory(ctx ctxDomain.IContext, userID int) (history v1.SpinSlotMachineHistory, err error) {

	_, err = r.app.Database().Read().
		Select(
			"sr.bet as bet_amount",
			"sr.payout as payout",
			"sr.symbols",
			"sr.winning",
			"sr.user_balance as balance",
			"srt.spin_result_type_id as spin_result_type_id",
			"srt.key as spin_result_type_key",
			"srt.name as spin_result_type_name",
			"sr.created_at",
		).
		From(database.SlotTableSpin.As("sr")).
		Join(database.SlotTableSpinResultType.As("srt"), "sr.spin_result_type_fk = srt.spin_result_type_id").
		Where("sr.user_fk = ?", userID).
		OrderDesc("sr.created_at").
		LoadContext(ctx, &history)

	if err != nil {
		return nil, r.app.Logger().DBLog(err)
	}

	return history, nil
}

func (r *Repository) GetSpinResultType(ctx ctxDomain.IContext, key string) (resultType *v1.SpinResultType, err error) {
	_, err = r.app.Database().Read().
		Select(
			"spin_result_type_id",
			"key as spin_result_type_key",
			"name as spin_result_type_name").
		From(database.SlotTableSpinResultType).
		Where("key = ?", key).
		LoadContext(ctx, &resultType)

	if err != nil {
		return nil, r.app.Logger().DBLog(err)
	}

	return resultType, nil
}
