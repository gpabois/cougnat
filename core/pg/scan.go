package pg

import (
	"github.com/gpabois/gostd/geojson"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
	"github.com/jackc/pgx/v5"
)

type ScanCommand struct {
	scan func(row pgx.Rows) result.Result[bool]
}

func (cmd ScanCommand) Exec(row pgx.Rows) result.Result[bool] {
	return cmd.scan(row)
}

type ScanCommands []ScanCommand

func (self ScanCommands) Append(cmds []ScanCommand) ScanCommands {
	return append(self, cmds...)
}

func (cmds ScanCommands) Exec(row pgx.Rows) result.Result[bool] {
	for _, cmd := range cmds {
		res := cmd.Exec(row)
		if res.HasFailed() {
			return result.Result[bool]{}.Failed(res.UnwrapError())
		}
	}

	return result.Success(true)
}

func ScanStringOption(opt *option.Option[string]) ScanCommands {
	return []ScanCommand{{scan: func(row pgx.Rows) result.Result[bool] {
		var ptrId *string

		if err := row.Scan(&ptrId); err != nil {
			return result.Result[bool]{}.Failed(err)
		}

		if ptrId != nil {
			*opt = option.Some(*ptrId)
		} else {
			*opt = option.None[string]()
		}

		return result.Success(true)
	}}}
}

func ScanPrimaryType(dest any) ScanCommands {
	return []ScanCommand{{scan: func(row pgx.Rows) result.Result[bool] {
		if err := row.Scan(dest); err != nil {
			return result.Result[bool]{}.Failed(err)
		}
		return result.Success(true)
	}}}
}

func ScanGeoJsonFeature(dest any) ScanCommands {
	return []ScanCommand{{scan: func(row pgx.Rows) result.Result[bool] {
		var jsonGeometry string
		if err := row.Scan(&jsonGeometry); err != nil {
			return result.Result[bool]{}.Failed(err)
		}
		serde.Reflect_DeserializeInto[geojson.Geometry]([]byte(jsonGeometry), "application/json", dest)
		return result.Success(true)
	}}}
}
