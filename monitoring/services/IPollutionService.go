package services

import (
	"context"

	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
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
	// Generate a tile image for a slippy-map
	GetTile(ctx context.Context, args GetTileArgs) result.Result[[]byte]
}
