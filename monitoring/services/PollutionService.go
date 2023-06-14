package services

import (
	"bytes"
	"context"
	"image"
	"image/png"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
)

type PollutionService struct {
	monitoringRepo repositories.IMonitoringRepository
	pollutionRepo  repositories.IPollutionRepository
}

type AggregationFunc func(tiles models.PollutionTileCollection) models.PollutionTile

func GetAggregationFunc(mode string) AggregationFunc {
	return func(tiles models.PollutionTileCollection) models.PollutionTile {
		tileID := tiles[0].ID

		tileIter := 
		
		intensity := iter.Reduce(
			iter.Map(iter.IterSlice(&tiles), 
				func(tile models.PollutionTile) models.PollutionData {
					return tile.Data
				},
			),
			func(acc models.PollutionIntensity, data models.PollutionData) models.PollutionIntensity {

			},
			models.PollutionIntensity{},
		)

		return models.PollutionTile{
			ID: tileID,
			Data: map[string] {
				"$all": intensity
			}
		}
	}
}

func (svc *PollutionService) GetTile(ctx context.Context, args GetTileArgs) result.Result[[]byte] {
	tile := image.NewRGBA(image.Rect(0, 0, 256, 256))
	tileBounds := args.TileIndex.Upscale(4)

	aggFunc := GetAggregationFunc(args.AggregationOperation.UnwrapOr(func() string { return "reduce_sum" }))
	
	matrixResult := result.Map(
		svc.pollutionRepo.GetPollutionTiles(tileBounds, args.TimeBounds),
		func(tiles models.PollutionTileCollection) models.PollutionMatrix {
			return tiles.IntoPollutionMatrix(tileBounds, aggFunc)
		},
	)

	if matrixResult.HasFailed() {
		return result.Result[[]byte]{}.Failed(matrixResult.UnwrapError())
	}

	matrix := matrixResult.Expect()

	dx := 256.0 / float64(tileBounds.DX())
	dy := 256.0 / float64(tileBounds.DY())

	buffer := bytes.NewBuffer(make([]byte, 256*256))
	err := png.Encode(buffer, tile)

	if err == nil {
		return result.Result[[]byte]{}.Failed(err)
	}

	return result.Success(buffer.Bytes())
}
