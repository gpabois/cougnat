package reporting_repositories

import (
	"context"

	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	pgx "github.com/jackc/pgx/v5"
)

type PgRepositories struct {
	db pgx.Conn
	ns option.Option[string]
}

func (pg *PgRepositories) Begin() result.Result[*TxPgRepositories] {
	var tx pgx.Tx
	tx, err := pg.db.Begin(context.Background())
	if err != nil {
		return result.Result[*TxPgRepositories]{}.Failed(err)
	}
	return result.Success(&TxPgRepositories{tx, pg.ns})
}

type TxPgRepositories struct {
	tx pgx.Tx
	ns option.Option[string]
}

func (repoTx *TxPgRepositories) Reports() IReportRepository {
	return TxPgReportRepository{tx, repoTx.ns}
}

func (repoTx *TxPgRepositories) Rollback() result.Result[bool] {
	repoTx.tx.Rollback(context.Background())
	return result.Success(true)
}

func (repoTx *TxPgRepositories) Commit() result.Result[bool] {
	if err := repoTx.tx.Commit(context.Background()); err != nil {
		return result.Result[bool]{}.Failed(err)
	}
	return result.Success(true)
}
