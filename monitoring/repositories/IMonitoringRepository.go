package repositories

import (
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
)

//go:generate mockery
type IMonitoringRepository interface {
	GetOrganisationMonitoring(organisationID string) result.Result[option.Option[models.OrganisationMonitoring]]
}
