package models

import (
	"github.com/gpabois/cougnat/core/tensor"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
)

type PollutionTile struct {
	ID   TimeTileIndex
	Data PollutionData
}

type PollutionMatrix = tensor.WSM[PollutionData]
type PollutionTimeSerie = time_serie.TimeSerie[PollutionMatrix]
