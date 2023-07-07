package reporting_repositories

import (
	tx "github.com/gpabois/cougnat/core/transaction"
)

type IRepositories interface {
	tx.ITransactionable[IRepositoriesTx]
}

// Transaction for all repository-related operations
type IRepositoriesTx interface {
	tx.ITransaction
	Reports() IReportRepository
}
