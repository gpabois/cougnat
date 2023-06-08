package services

import (
	"time"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
)

//go:generate mockery
type IPollutionService interface {
	GetPollutionTiles(orgID string, sectionID []string, zoom int, begin time.Time, end time.Time) result.Result[models.PolTileCollection]
}
