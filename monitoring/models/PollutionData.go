package models

import "github.com/gpabois/cougnat/core/iter"

type PollutionData map[string]PollutionIntensity

func PollutionDataSum(d1 PollutionData, d2 PollutionData) PollutionData {
	var res PollutionData

	res = iter.Reduce(iter.IterMap(&d1), func(acc PollutionData, entry iter.KV[string, PollutionIntensity]) PollutionData {
		return acc.Add(entry.GetFirst(), entry.GetSecond())
	}, res)

	res = iter.Reduce(iter.IterMap(&d2), func(acc PollutionData, entry iter.KV[string, PollutionIntensity]) PollutionData {
		return acc.Add(entry.GetFirst(), entry.GetSecond())
	}, res)

	return res
}

// Add intensity to the pollution data
func (data PollutionData) Add(typ string, intensity PollutionIntensity) PollutionData {
	inten := data[typ]
	inten.Count += intensity.Count
	inten.Weight += inten.Weight

	data[typ] = inten
	return data
}

func PollutionDataReduceSum(data PollutionData) PollutionData {
	return data.ReduceSum()
}

func (data PollutionData) Sum(it iter.Iterator[PollutionData]) PollutionData {
	return iter.Reduce(it, PollutionDataSum, data)
}

// Reduce the pollution data into a pollution data with a single pollution intensity
// stored as "$all"
func (data PollutionData) ReduceSum() PollutionData {
	intensity := iter.Reduce(
		iter.IterMap(&data),
		func(acc PollutionIntensity, kv iter.KV[string, PollutionIntensity]) PollutionIntensity {
			acc.Count += kv.GetSecond().Count
			acc.Weight += kv.GetSecond().Weight
			return acc
		},
		PollutionIntensity{},
	)

	return PollutionData{"$all": intensity}
}
