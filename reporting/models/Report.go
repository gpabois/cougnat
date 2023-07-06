package reporting_models

import (
	"time"

	auth "github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/option"
	geo "github.com/gpabois/gostd/geojson"
	"github.com/jinzhu/copier"
)

type ReportID = int
type ReportTypeID = int

type ReportType struct {
	ID     int    `serde:"id"`
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

type NewReport struct {
	Owner    option.Option[auth.ActorID] `serde:"owner"`
	Location geo.Feature                 `serde:"location"`
	TypeID   ReportTypeID                `serde:"type_id"`
	Rate     int                         `serde:"rate"`
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
