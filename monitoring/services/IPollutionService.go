package services

import (
	"context"

	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
	"github.com/gpabois/cougnat/monitoring/models"
)

type GetPollutionMatrixArgs struct {
	OrganisationID       string
	SectionID            []string
	TileBounds           slippy_map.TileBounds
	TimeBounds           time_serie.TimeInterval
	AggregationOperation option.Option[string]
}

type GetTileArgs struct {
	OrganisationID       string
	SectionID            []string
	TileIndex            slippy_map.TileIndex
	TimeBounds           time_serie.TimeInterval
	AggregationOperation option.Option[string]
}

//go:generate mockery
type IPollutionService interface {
	GetPollutionMatrix(ctx context.Context, args GetPollutionMatrixArgs) result.Result[models.PollutionMatrix]
}
