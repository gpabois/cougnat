package services

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
)

type AggregatePollutionTilesArgs struct {
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
	AggregatePollutionMatrix(ctx context.Context, args AggregatePollutionTilesArgs) result.Result[models.PollutionMatrix]
}
