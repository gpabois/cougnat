package models

import (
	"github.com/gpabois/cougnat/core/geojson"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type SectorMonitoring struct {
	ID    string
	label string
	area  geojson.Geometry
	emits []reportingModels.ReportType
}

type OrganisationMonitoring struct {
	OrganisationID string
	Sectors        []SectorMonitoring
}
