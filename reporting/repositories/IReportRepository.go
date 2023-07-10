package reporting_repositories

import (
	"github.com/gpabois/cougnat/core/transaction"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type IReportOps interface {
	// Create a report
	Create(report reporting_models.NewReport) result.Result[reporting_models.ReportID]
	// Get a report
	GetById(id reporting_models.ReportID) result.Result[option.Option[reporting_models.Report]]
	// Delete a report
	Delete(reportID reporting_models.ReportID) result.Result[bool]
}

//go:generate mockery
type IReportRepository interface {
	// Should be able to execute operations within a transaction
	transaction.ITransactionable[IReportRepositoryTx]
	// Operations
	IReportOps
}

// A transaction to the reports's repository
//
//go:generate mockery
type IReportRepositoryTx interface {
	transaction.ITransaction
	IReportOps
}
