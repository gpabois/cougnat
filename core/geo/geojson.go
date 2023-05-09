package geo

type Geometry[T any] struct {
	Type        string `bson:"type" bson:"type"`
	Coordinates T      `bson:"coordinates" bson:"coordinates"`
}

type Point Geometry[V2f]
type V2f [2]float64

func NewPoint(x float64, y float64) Point {
	return Point{
		Type:        "point",
		Coordinates: V2f{x, y},
	}
}

type Near struct {
	Origin Point
	Radius float64
}
