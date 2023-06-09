package slippy_map

import (
	"math"

	"github.com/gpabois/cougnat/core/geojson"
	"github.com/gpabois/cougnat/core/result"
)

type TileIndex struct {
	X int
	Y int
	Z int
}

func (tile TileIndex) ZoomOut(out int) TileIndex {
	return TileIndex{
		X: int(float64(tile.X) * math.Pow(-2.0, float64(out))),
		Y: int(float64(tile.Y) * math.Pow(-2.0, float64(out))),
		Z: tile.Z - out,
	}
}

func (tile TileIndex) ZoomIn(out int) TileIndex {
	return TileIndex{
		X: int(float64(tile.X) * math.Pow(2.0, float64(out))),
		Y: int(float64(tile.Y) * math.Pow(2.0, float64(out))),
		Z: tile.Z - out,
	}
}

func TileIndexFromLatLng(lat float64, lng float64, zoom int) TileIndex {
	n := math.Exp2(float64(zoom))
	x := int(math.Floor((lng + 180.0) / 360.0 * n))

	if float64(x) >= n {
		x = int(n - 1)
	}

	y := int(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * n))

	return TileIndex{
		X: x,
		Y: y,
		Z: zoom,
	}
}

func (t TileIndex) IntoLatLng() (float64, float64) {
	n := math.Pi - 2.0*math.Pi*float64(t.Y)/math.Exp2(float64(t.Z))
	lat := 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	lng := float64(t.X)/math.Exp2(float64(t.Z))*360.0 - 180.0
	return lat, lng
}

// Returns Tile Index from Point[lng, lat]
func TileIndexFromGeometry(geometry geojson.Geometry, zoom int) result.Result[TileIndex] {
	lng := geometry.Coordinates[0]
	lat := geometry.Coordinates[1]

	return result.Success(TileIndexFromLatLng(lat, lng, zoom))
}
