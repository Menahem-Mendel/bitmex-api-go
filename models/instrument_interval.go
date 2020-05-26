package models

type InstrumentInterval struct {
	Intervals []string `json:"intervals"`
	Symbols   []string `json:"symbols"`
}
