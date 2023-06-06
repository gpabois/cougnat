package models

import (
	"time"

	auth "github.com/gpabois/cougnat/auth/models"
	geo "github.com/gpabois/cougnat/core/geojson"
	"github.com/gpabois/cougnat/core/option"
)

type ReportID = string

type ReportType struct {
	Name  string
	Type  string
	Label string
}

// A report
type Report struct {
	ID        option.Option[ReportID]
	Owner     option.Option[auth.ActorID]
	Location  geo.Feature
	Nature    string
	Rate      int
	CreatedAt time.Time
}

func (report Report) ObjectID() option.Option[auth.ObjectID] {
	return option.Map(report.ID, ReportObjectID)
}

func ReportObjectID(reportID ReportID) auth.ObjectID {
	return auth.NewObjectID("report", reportID)
}
