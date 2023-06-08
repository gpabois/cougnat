package models

import (
	"time"

	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	"github.com/gpabois/cougnat/core/unit"
)

type TimeTileIndex struct {
	X            int           `serde:"x"`
	Y            int           `serde:"y"`
	Z            int           `serde:"z"`
	T            int           `serde:"t`
	TimeSampling unit.Sampling `serde:"time_sampling`
}

func (tti TimeTileIndex) ZoomOut(out int) TimeTileIndex {
	return TimeTileIndex{
		X: tti.X, Y: tti.Y, Z: tti.Z - out, T: tti.T, TimeSampling: tti.TimeSampling,
	}
}

func GetTimeTileIndex(tile slippy_map.TileIndex, date time.Time, sampling unit.Sampling) TimeTileIndex {
	period := unit.Sample(date, sampling)

	return TimeTileIndex{
		X:            tile.X,
		Y:            tile.Y,
		Z:            tile.Z,
		T:            period,
		TimeSampling: sampling,
	}
}

type PolTileCollection = []PolTile

type PolTimeSlice []PolTimeSerieEntry

type PolTimeSerieEntry struct {
	T      int `serde:"t"'`
	Matrix []PolTile
}

type PolData struct {
	count  int
	weight int
}

type PolTile struct {
	ID      TimeTileIndex
	Cluster TimeTileIndex
	data    map[string]PolData
}
