package reporting_repositories

import (
	"context"

	"github.com/gpabois/cougnat/core/pg"
	"github.com/gpabois/cougnat/core/transaction"
	tx "github.com/gpabois/cougnat/core/transaction"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	reporting_pg "github.com/gpabois/cougnat/reporting/repositories/pg"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/jackc/pgx/v5"
)

// A postgres-based repository for Reports
type PgReportRepository struct {
	conn pgx.Conn
}

func (repo *PgReportRepository) Begin(parent option.Option[tx.Transaction]) result.Result[IReportRepositoryTx] {
	return result.Map(
		pg.Begin(repo.conn, context.Background(), parent, option.Some("tx.pg")),
		func(pgTx tx.PgTransaction) IReportRepositoryTx {
			return newPgReportRepository(pgTx)
		},
	)
}

func (repo PgReportRepository) GetById(reportID reporting_models.ReportID) result.Result[option.Option[reporting_models.Report]] {
	return transaction.With(repo.Begin(option.None[tx.Transaction]()), func(tx IReportRepositoryTx) result.Result[option.Option[reporting_models.Report]] {
		return tx.GetById(reportID)
	})
}

func (repo PgReportRepository) Create(report reporting_models.NewReport) result.Result[reporting_models.ReportID] {
	return transaction.With(repo.Begin(option.None[tx.Transaction]()),
		func(tx IReportRepositoryTx) result.Result[reporting_models.ReportID] {
			return tx.Create(report)
		},
	)
}

// Transaction to a postgres-based repository for reports
type PgReportRepositoryTx struct {
	Inner transaction.PgTransaction
}

func newPgReportRepository(tx transaction.PgTransaction) PgReportRepositoryTx {
	return PgReportRepositoryTx{tx}
}

func (repo PgReportRepositoryTx) IsManaged() bool {
	return repo.Inner.IsManaged()
}

func (repo PgReportRepositoryTx) Commit() result.Result[bool] {
	return repo.Inner.Commit()
}

func (repo PgReportRepositoryTx) Rollback() result.Result[bool] {
	return repo.Inner.Rollback()
}

func (repo PgReportRepositoryTx) GetById(reportID reporting_models.ReportID) result.Result[option.Option[reporting_models.Report]] {
	withQuery := pg.WithQuery(pg.WithOne(func(rows pgx.Rows) result.Result[reporting_models.Report] {
		report := reporting_models.Report{}
		scanRes := reporting_pg.ScanReport(&report).Exec(rows)
		if scanRes.HasFailed() {
			return result.Result[reporting_models.Report]{}.Failed(scanRes.UnwrapError())
		}

		return result.Success(report)
	}))

	return withQuery(
		repo.Inner.Inner,
		context.Background(),
		reporting_pg.PG_REPORT_GET_BY_ID_QUERY,
		reportID,
	)
}

func (repo PgReportRepositoryTx) Create(report reporting_models.NewReport) result.Result[reporting_models.ReportID] {
	withQuery := pg.WithQuery(pg.WithExpectedOne(func(rows pgx.Rows) result.Result[reporting_models.ReportID] {
		var reportId reporting_models.ReportID
		err := rows.Scan(&reportId)
		if err != nil {
			return result.Result[reporting_models.ReportID]{}.Failed(err)
		}
		return result.Success(reportId)
	}))

	return withQuery(repo.Inner.Inner, context.Background(),
		reporting_pg.PG_INSERT_REPORT_QUERY,
		report.Owner.UnwrapOrZero().ID,
		report.Owner.UnwrapOrZero().Nature,
		report.TypeID,
		report.Rate,
		report.Location,
	)
}

func (repo PgReportRepositoryTx) Delete(reportID reporting_models.ReportID) result.Result[bool] {
	res := pg.Query(repo.Inner.Inner, context.Background(), reporting_pg.PG_DELETE_BY_ID_REPORT_QUERY, reportID)
	if res.HasFailed() {
		return result.Result[bool]{}.Failed(res.UnwrapError())
	}
	res.Expect().Close()
	return result.Success(true)
}
