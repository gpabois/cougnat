package pg

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type FieldBinding struct {
	id   exp.IdentifierExpression
	dest any
}

func (fb FieldBinding) Exec(q Query) Query {
	q.sqlQuery = q.sqlQuery.SelectAppend(fb.id)
	q.commands = q.commands.Append(ScanPrimaryType(fb.dest))
	return q
}

// Query builder
type Query struct {
	sqlQuery *goqu.SelectDataset
	commands ScanCommands
}
