package models

import (
	"time"

	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
	"github.com/gpabois/cougnat/core/unit"
)

type TimeTileIndex struct {
	TileID slippy_map.TileIndex `serde:"tile"`
	TimeID time_serie.TimeIndex `serde:"time"`
}

func (tti TimeTileIndex) ZoomOut(out int) TimeTileIndex {
	tti.TileID = tti.TileID.ZoomOut(out)
	return tti
}

func (tti TimeTileIndex) ZoomIn(in int) TimeTileIndex {
	tti.TileID = tti.TileID.ZoomIn(in)
	return tti
}

// Generate a TTI based on the tile coordinates, the datetime and the time sampling
func GetTimeTileIndex(tile slippy_map.TileIndex, date time.Time, sampling unit.Sampling) TimeTileIndex {
	step := unit.Sample(date, sampling)

	return TimeTileIndex{
		TileID: tile,
		TimeID: time_serie.TimeIndex{
			Step:     step,
			Sampling: sampling,
		},
	}
}
