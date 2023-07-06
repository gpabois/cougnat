package reporting_repositories

import (
	"database/sql/driver"

	"github.com/doug-martin/goqu/v9"
	auth_models "github.com/gpabois/cougnat/auth/models"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

type PgReportRepository struct {
	conn  driver.Conn
	table string
}

func (repo *PgReportRepository) Create(report reporting_models.NewReport) result.Result[reporting_models.ReportID] {
	owner := report.Owner.UnwrapOr(func() auth_models.ActorID { auth_models.AnonymousID(option.Some("anonymous")) })

	locationRes := serde.Serialize(report.Location.Geometry, "application/json")

	if locationRes.HasFailed() {
		return result.Result[reporting_models.ReportID]{}.Failed(locationRes.UnwrapError())
	}

	ds := goqu.Insert(repo.table).
		Cols("owner__id", "owner__nature", "type__id", "rate", "location").
		Vals(goqu.Vals{
			owner.ID,
			owner.Nature,
			report.TypeID,
			report.Rate,
			goqu.Func("ST_GeomFromGeoJSON", locationRes),
		})

	query, _, err := ds.ToSQL()

	if err != nil {
		return result.Result[reporting_models.ReportID]{}.Failed(err)
	}

	goqu.New("postgres", repo.conn)
}
