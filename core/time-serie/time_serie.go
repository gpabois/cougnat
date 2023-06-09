package time_serie

import (
	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/unit"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type TimeIndex struct {
	Step     int           `serde:"t`
	Sampling unit.Sampling `serde:"time_sampling`
}

type TimeSeriePoint[Value any] struct {
	Step  int
	Value Value
}

func FromKeyValue[Value any](kv iter.KV[int, Value]) TimeSeriePoint[Value] {
	return TimeSeriePoint[Value]{
		Step:  kv.Key,
		Value: kv.Value,
	}
}

type TimeSerie[Value any] struct {
	Sampling unit.Sampling
	Begin    int
	End      int
	Points   []TimeSeriePoint[Value]
}

func FromKeyValueIterator[Value any](iterator iter.Iterator[iter.KV[int, Value]], sampling unit.Sampling) TimeSerie[Value] {
	points := iter.CollectToSlice[[]TimeSeriePoint[Value]](iter.Map(iterator, FromKeyValue[Value]))
	// Order the points
	slices.SortFunc(points, func(a, b TimeSeriePoint[Value]) bool {
		return a.Step < b.Step
	})
	begin, end := points[0], points[len(points)-1]

	return TimeSerie[Value]{
		Sampling: sampling,
		Begin:    begin.Step,
		End:      end.Step,
		Points:   points,
	}
}
