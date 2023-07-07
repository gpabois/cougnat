package reporting_pg

import (
	"github.com/doug-martin/goqu/v9"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

func GetByIdQuery(ns option.Option[string], reportID reporting_models.ReportID) result.Result[string] {
	reports := goqu.T(pg.PgFullTableName("reports", ns))
	reportTypes := goqu.T(pg.PgFullTableName("reports_types", ns))
	joinKey := goqu.On(reportTypes.Col("id").Eq(reports.Col("type__id")))

	query, _, err := goqu.From(reports).
		Join(reportTypes, joinKey).
		Select(
			reports.Col("id"),
			reports.Col("owner__id"),
			reports.Col("owner__nature"),
			reportTypes.Col("id"),
			reportTypes.Col("name"),
			reportTypes.Col("label"),
			reportTypes.Col("nature"),
			reports.Col("rate"),
			goqu.Func("ST_AsGeoJSON", "reports.location"),
		).
		Where(reports.Col("id").Eq(reportID)).
		ToSQL()

	if err != nil {
		return result.Result[string]{}.Failed(err)
	}

	return result.Success(query)
}

func CreateReportQuery(table string, report reporting_models.NewReport) result.Result[string] {
	owner := report.Owner.UnwrapOr(func() auth_models.ActorID { return auth_models.AnonymousID(option.Some("anonymous")) })
	locationRes := serde.Serialize(report.Location.Geometry, "application/json")

	if locationRes.HasFailed() {
		return result.Result[string]{}.Failed(locationRes.UnwrapError())
	}

	ds := goqu.Insert(table).
		Cols("owner__id", "owner__nature", "type__id", "rate", "location").
		Vals(goqu.Vals{
			owner.ID,
			owner.Nature,
			report.TypeID,
			report.Rate,
			goqu.Func("ST_GeomFromGeoJSON", locationRes),
		}).Returning("id")

	query, _, err := ds.ToSQL()

	if err != nil {
		return result.Result[string]{}.Failed(err)
	}

	return result.Success(query)
}
