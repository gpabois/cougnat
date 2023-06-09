package services

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
)

type GetPollutionTilesArgs struct {
	OrganisationID string
	SectionID      []string
	Zoom           int
	TimeWindow     struct {
		Begin int
		End   int
	}
}

//go:generate mockery
type IPollutionService interface {
	GetPollutionTiles(ctx context.Context, args GetPollutionTilesArgs) result.Result[models.PolTileCollection]
}
