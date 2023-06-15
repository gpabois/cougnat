package services

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"

	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/ops"
	"github.com/gpabois/cougnat/core/result"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
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

		return models.PollutionTile{
			ID: tileID,
			Data: models.PollutionData{}.Sum(iter.Map(tiles.IterData(), func(data models.PollutionData) models.PollutionData {
				return data.ReduceSum()
			})),
		}
	}
}

// Paint the pollution onto an image depending on the $all's weight value
func DrawPollutionTile(img *image.RGBA, tile models.PollutionTile, tileBounds slippy_map.TileBounds, threshold int) *image.RGBA {
	dx := float64(img.Bounds().Dx()) / float64(tileBounds.DX())
	dy := float64(img.Bounds().Dy()) / float64(tileBounds.DY())

	absX := float64(tile.ID.TileID.X-tileBounds.MinX()) * dx
	absY := float64(tile.ID.TileID.Y-tileBounds.MinY()) * dy

	rect := image.Rect(
		int(math.Round(absX)),
		int(math.Round(absY)),
		int(math.Round(absX+dx)),
		int(math.Round(absY+dy)),
	)

	// Keep in range [0, 1]
	t := ops.Bounds(1-float64(tile.Data.Get("$all").Weight), 0.0, 1.0)

	// No values, we don't perform any operations
	if t == 0.0 {
		return img
	}

	// Go from Green to Red
	hue := 0.33 * t
	r, g, b := colorutil.HslToRgb(hue, 0.5*(1-t), 0.6)
	tileColor := color.RGBA{r, g, b, 255}
	draw.Draw(img, rect, &image.Uniform{tileColor}, image.ZP, draw.Src)
	return img
}

// Generate a tile image for a slippy-map
// TODO(1): Filter the tiles which intersects with the monitoring sections allowed to the requester's organisation
// TODO(2): Implements a cache system to avoid useless computation
func (svc *PollutionService) GetTile(ctx context.Context, args GetTileArgs) result.Result[[]byte] {
	// A tile image is alway 256x256 as per the slippy-map requirements
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))

	// Upscale the tile bounds to get a pollution resolution
	// With an upscale of 4, it generates a pollution tile which resolution is equal to half the image's resolution
	tileBounds := args.TileIndex.Upscale(4)

	// Determine the reduction to $all strategy (reduce_sum, reduce_avg, reduce_max, reduce_min, ...)
	// sum, avg, max, min,... are automatically suffixed with "reduce_" as it is required to generate the tile to use the intensity of $all only.
	// To reduce with only a subset of all report types, the requester must use the filter option.
	aggMode := fmt.Sprintf("reduce_%s", args.AggregationOperation.UnwrapOr(func() string { return "sum" }))
	aggFunc := GetAggregationFunc(aggMode)

	// We create a collection of pollution pixels
	polPixelsResult := result.Map(
		svc.pollutionRepo.GetPollutionTiles(tileBounds, args.TimeBounds),
		// Apply the reduction to get raw pollution matrix elements (row = X, col = Y)
		// Perform aggregation on the tiles which (X,Y) are equals
		func(tiles models.PollutionTileCollection) models.PollutionTileCollection {
			return tiles.IntoPollutionMatrix(tileBounds, aggFunc)
		},
	)

	// Failed to get a collection of tiles
	if polPixelsResult.HasFailed() {
		return result.Result[[]byte]{}.Failed(polPixelsResult.UnwrapError())
	}

	// Draw each tile onto the image
	for _, tile := range polPixelsResult.Expect() {
		DrawPollutionTile(img, tile, tileBounds, 250)
	}

	// Create a buffer to get the image content
	buffer := bytes.NewBuffer(make([]byte, 256*256))

	// Encode the image into a png and store its content in the buffer
	err := png.Encode(buffer, img)

	// Failed to encode the image
	if err == nil {
		return result.Result[[]byte]{}.Failed(err)
	}

	return result.Success(buffer.Bytes())
}
