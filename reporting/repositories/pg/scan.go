package reporting_pg

import (
	auth_pg "github.com/gpabois/cougnat/auth/repositories/pg"
	"github.com/gpabois/cougnat/core/pg"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
)

// Scan the row to return a ReportType
// Expected order: ID, Name, Label, Nature
func ScanReportType(reportType *reporting_models.ReportType) pg.ScanCommands {
	return pg.Scan(&reportType.ID).
		Append(pg.Scan(&reportType.Name)).
		Append(pg.Scan(&reportType.Label)).
		Append(pg.Scan(&reportType.Nature))
}

// Scan the row to return a Report
// Expected order: ID, ActorID, ReportType, Rate, Location
func ScanReport(report *reporting_models.Report) pg.ScanCommands {
	return pg.Scan(&report.ID).
		Append(auth_pg.ScanOptionalActorID(&report.Owner)).
		Append(ScanReportType(&report.Type)).
		Append(pg.Scan(&report.Rate)).
		Append(pg.Scan(&report.Location))

}
