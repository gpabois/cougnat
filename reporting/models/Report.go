package reporting_models

import (
	"time"

	auth "github.com/gpabois/cougnat/auth/models"
	geo "github.com/gpabois/cougnat/core/geojson"
	"github.com/gpabois/cougnat/core/option"
	"github.com/jinzhu/copier"
)

type ReportID = string

type ReportType struct {
	Name   string `serde:"name"`
	Label  string `serde:"label"`
	Nature string `serde:"nature"`
}

// A report
type Report struct {
	ID         option.Option[ReportID]     `serde:"id"`
	Owner      option.Option[auth.ActorID] `serde:"owner"`
	Location   geo.Feature                 `serde:"location"`
	Type       ReportType                  `serde:"type"`
	Rate       int                         `serde:"rate"`
	ReportedAt time.Time                   `serde:"reported_at"`
}

func (report Report) From(attr any) Report {
	copier.Copy(&report, attr)
	return report
}

// Return the ObjectID of the Report
func (report Report) ObjectID() option.Option[auth.ObjectID] {
	return option.Map(report.ID, ReportObjectID)
}

// Transform the ReportID into an ObjectID
func ReportObjectID(reportID ReportID) auth.ObjectID {
	return auth.NewObjectID("report", reportID)
}
