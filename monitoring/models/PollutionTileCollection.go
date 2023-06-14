package models

import (
	"github.com/gpabois/cougnat/core/iter"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
)

// Represents a raw list of pollution tiles
type PollutionTileCollection []PollutionTile

func (col PollutionTileCollection) Iter() iter.Iterator[PollutionTile] {
	return iter.IterSlice(&col)
}

func (col PollutionTileCollection) IterData() iter.Iterator[PollutionData] {
	return iter.Map(col.Iter(), func(tile PollutionTile) PollutionData {
		return tile.Data
	})
}

// Transform the collection into a pollution matrix
func (col PollutionTileCollection) IntoPollutionMatrix(bounds slippy_map.TileBounds, aggregation func(tiles PollutionTileCollection) PollutionTile) PollutionTileCollection {
	// Group tiles by TileIndex
	groups := iter.Group[PollutionTileCollection](col.Iter(), func(tile PollutionTile) slippy_map.TileIndex { return tile.ID.TileID })

	// Perform tile collection aggregation
	elements := iter.CollectToSlice[PollutionTileCollection](iter.Map(groups.Iter(), func(group iter.KV[slippy_map.TileIndex, PollutionTileCollection]) PollutionTile {
		return aggregation(group.GetSecond())
	}))

	// Create a sparse matrix
	return elements
}
