package pg

import (
	"fmt"

	"github.com/gpabois/gostd/option"
)

func FullTableName(table string, ns option.Option[string]) string {
	if ns.IsNone() {
		return table
	}

	return fmt.Sprintf("%s_%s", ns.Expect(), table)
}
