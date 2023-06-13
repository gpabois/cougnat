package models

import (
	"time"

	"github.com/gpabois/cougnat/core/iter"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	"github.com/gpabois/cougnat/core/tensor"
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

// Represents a random list of pollution tiles
type PollutionTileCollection []PollutionTile

func (col PollutionTileCollection) Iter() iter.Iterator[PollutionTile] {
	return iter.IterSlice(&col)
}

func (col PollutionTileCollection) IntoPollutionMatrix(upperRight slippy_map.TileIndex, lowerRight slippy_map.TileIndex) PollutionMatrix {
	rowLength := lowerRight.X - upperRight.X
	columnLength := lowerRight.Y - upperRight.Y

	return tensor.NewWSM(
		upperRight.X,
		rowLength,
		upperRight.Y,
		columnLength,
		col,
		func(tile PollutionTile) (int, int) {
			return tile.ID.TileID.X, tile.ID.TileID.Y
		},
		func(tile PollutionTile) PollutionData { return tile.Data },
	)
}

func (col PollutionTileCollection) IntoTimeSerie(upperRight slippy_map.TileIndex, lowerRight slippy_map.TileIndex, sampling unit.Sampling) PollutionTimeSerie {
	// Group the tiles per time steps
	timeGroup := iter.Group[PollutionTileCollection](col.Iter(), func(tile PollutionTile) int { return tile.ID.TimeID.Step })

	// Transform PollutionTileCollection into a PollutionMatrix
	// At constant-time
	points := iter.Map(
		timeGroup.Iter(),
		func(kv iter.KV[int, PollutionTileCollection]) iter.KV[int, PollutionMatrix] {
			return iter.KV[int, PollutionMatrix]{
				Key:   kv.Key,
				Value: kv.Value.IntoPollutionMatrix(upperRight, lowerRight),
			}
		},
	)

	return time_serie.FromIter[PollutionMatrix](points, sampling)
}

type PollutionData map[string]struct {
	count  int
	weight int
}

type PollutionTile struct {
	ID      TimeTileIndex
	Cluster TimeTileIndex
	Data    PollutionData
}

type PollutionMatrix = tensor.WSM[PollutionData]
type PollutionTimeSerie = time_serie.TimeSerie[PollutionMatrix]
