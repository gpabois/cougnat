package transaction

import "github.com/gpabois/gostd/result"

type ITransactionable[Tx ITransaction] interface {
	Begin() result.Result[Tx]
}

type ITransaction interface {
	Commit() result.Result[bool]
	Rollback() result.Result[bool]
}

func BeginTransaction[Tx ITransaction, TxAble ITransactionable[Tx]](txAble TxAble, ctx func(tx Tx) result.Result[bool]) (res result.Result[bool]) {
	// Get a transaction
	txRes := txAble.Begin()
	if txRes.HasFailed() {
		return result.Result[bool]{}.Failed(txRes.UnwrapError())
	}
	tx := txRes.Expect()

	// Rollback on panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	res = ctx(tx)
	if res.HasFailed() {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return res
}
