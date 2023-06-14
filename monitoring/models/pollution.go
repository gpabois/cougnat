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

func (col PollutionTileCollection) IterData() iter.Iterator[PollutionData] {
	return iter.Map(col.Iter(), func(tile PollutionTile) PollutionData {
		return tile.Data
	})
}

func (col PollutionTileCollection) Sum() PollutionData {
	return iter.Reduce(col.IterData(), PollutionDataSum, PollutionData{})
}

func (col PollutionTileCollection) ReduceSum() PollutionTile {
	tileID := col[0].ID
	intensity := iter.Reduce(col.IterData(), func(acc PollutionIntensity, data PollutionData) { return PollutionDataReduceSum(data) }, PollutionIntensity{})

	return PollutionTile{
		ID: tileID,
		Data: PollutionData{
			"$all": intensity,
		},
	}
}

func (col PollutionTileCollection) IntoPollutionMatrix(bounds slippy_map.TileBounds, aggregation func(tiles PollutionTileCollection) PollutionTile) PollutionMatrix {
	rowLength := bounds.DX()
	columnLength := bounds.DY()

	// Group tiles by TileIndex
	groups := iter.Group[PollutionTileCollection](col.Iter(), func(tile PollutionTile) slippy_map.TileIndex { return tile.ID.TileID })

	// Perform tile collection aggregation
	elements := iter.CollectToSlice[PollutionTileCollection](iter.Map(groups.Iter(), func(group iter.KV[slippy_map.TileIndex, PollutionTileCollection]) PollutionTile {
		return aggregation(group.GetSecond())
	}))

	// Create a sparse matrix
	return tensor.NewWSM(bounds.MinX(), rowLength, bounds.MinY(), columnLength, iter.IntoSliceIterable(elements),
		func(tile PollutionTile) (int, int) {
			return tile.ID.TileID.X, tile.ID.TileID.Y
		},
		func(tile PollutionTile) PollutionData { return tile.Data },
	)
}

func (col PollutionTileCollection) IntoTimeSerie(tileBounds slippy_map.TileBounds, sampling unit.Sampling, aggregation func(tiles PollutionTileCollection) PollutionTile) PollutionTimeSerie {
	// Group the tiles per time steps
	timeGroup := iter.Group[PollutionTileCollection](col.Iter(), func(tile PollutionTile) int { return tile.ID.TimeID.Step })

	// Transform PollutionTileCollection into a PollutionMatrix
	// At constant-time
	points := iter.Map(
		timeGroup.Iter(),
		func(kv iter.KV[int, PollutionTileCollection]) iter.KV[int, PollutionMatrix] {
			return iter.KV[int, PollutionMatrix]{
				Key:   kv.Key,
				Value: kv.Value.IntoPollutionMatrix(tileBounds, aggregation),
			}
		},
	)

	return time_serie.FromIter[PollutionMatrix](points, sampling)
}

type PollutionIntensity struct {
	Count  int `serde:"count"`
	Weight int `serde:"weight"`
}

type PollutionData map[string]PollutionIntensity

func PollutionDataSum(d1 PollutionData, d2 PollutionData) PollutionData {
	var res PollutionData

	res = iter.Reduce(iter.IterMap(&d1), func(acc PollutionData, entry iter.KV[string, PollutionIntensity]) PollutionData {
		return acc.AddIntensity(entry.GetFirst(), entry.GetSecond())
	}, res)

	res = iter.Reduce(iter.IterMap(&d2), func(acc PollutionData, entry iter.KV[string, PollutionIntensity]) PollutionData {
		return acc.AddIntensity(entry.GetFirst(), entry.GetSecond())
	}, res)

	return res
}

// Add intensity to the pollution data
func (data PollutionData) AddIntensity(typ string, intensity PollutionIntensity) PollutionData {
	inten := data[typ]
	inten.Count += intensity.Count
	inten.Weight += inten.Weight

	data[typ] = inten
	return data
}

func PollutionDataReduceSum(data PollutionData) PollutionIntensity {
	return data.ReduceSum()
}

// Reduce the pollution data into a single pollution intensity
func (data PollutionData) ReduceSum() PollutionIntensity {
	return iter.Reduce(
		iter.IterMap(&data),
		func(acc PollutionIntensity, kv iter.KV[string, PollutionIntensity]) PollutionIntensity {
			acc.Count += kv.GetSecond().Count
			acc.Weight += kv.GetSecond().Weight
			return acc
		},
		PollutionIntensity{},
	)
}

type PollutionTile struct {
	ID   TimeTileIndex
	Data PollutionData
}

type PollutionMatrix = tensor.WSM[PollutionData]
type PollutionTimeSerie = time_serie.TimeSerie[PollutionMatrix]
