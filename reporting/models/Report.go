package models

import (
	"time"

	auth "github.com/gpabois/cougnat/auth/models"
	geo "github.com/gpabois/cougnat/core/geo"
)

type ReportID = string

// A report
type Report struct {
	ID        ReportID
	Owner     auth.ActorID
	Location  geo.Point
	Nature    string
	Rate      int
	CreatedAt time.Time
}

func (report Report) ObjectID() auth.ObjectID {
	return auth.NewObjectID("report", report.ID)
}

func ReportObjectID(reportID ReportID) auth.ObjectID {
	return auth.NewObjectID("report", reportID)
}
