package models

type PollutionIntensity struct {
	Count  int `serde:"count"`
	Weight int `serde:"weight"`
}
