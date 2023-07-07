package reporting_repositories

import (
	"context"

	reporting_models "github.com/gpabois/cougnat/reporting/models"
	reporting_pg "github.com/gpabois/cougnat/reporting/repositories/pg"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/jackc/pgx/v5"
)

// Tx Report Repository
type TxPgReportRepository struct {
	tx pgx.Tx
	ns option.Option[string]
}

func NewPgReportRepository(tx pgx.Tx, ns option.Option[string]) TxPgReportRepository {
	return TxPgReportRepository{tx, ns}
}

func (repo TxPgReportRepository) GetById(reportID reporting_models.ReportID) result.Result[option.Option[reporting_models.Report]] {
	// Generate the query
	var queryRes result.Result[string]
	if queryRes = reporting_pg.GetByIdQuery(repo.ns, reportID); queryRes.HasFailed() {
		return result.Result[option.Option[reporting_models.Report]]{}.Failed(queryRes.UnwrapError())
	}
	query := queryRes.Expect()

	// Execute the query
	rows, err := repo.tx.Query(context.Background(), query)

	if err != nil {
		return result.Result[option.Option[reporting_models.Report]]{}.Failed(err)
	}

	defer rows.Close()
	if exist := rows.Next(); !exist {
		return result.Success(option.None[reporting_models.Report]())
	}

	report := reporting_models.Report{}
	res := reporting_pg.ScanReport(&report).Exec(rows)

	if res.HasFailed() {
		return result.Result[option.Option[reporting_models.Report]]{}.Failed(res.UnwrapError())
	}
	// Return the report
	return result.Success(option.Some(report))
}
