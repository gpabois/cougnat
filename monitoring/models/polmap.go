package models

import (
	"math"
	"time"

	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
)

type TimeTileIndex struct {
	X int
	Y int
	Z int
	T int
}

func GetTimeTileIndex(tile slippy_map.TileIndex, date time.Time, sampling_in_seconds int) TimeTileIndex {
	t := int(math.Round(float64(date.UnixMilli()) / float64(1000) / float64(sampling_in_seconds)))

	return TimeTileIndex{
		X: tile.X,
		Y: tile.Y,
		Z: tile.Z,
		T: t,
	}
}

type PolData struct {
	count int
}

type PolTile struct {
	ID   TimeTileIndex
	data map[string]PolData
}
