package rest

import (
	"encoding/json"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
	"github.com/pkg/errors"
)

type StatsService struct {
	request
}

type StatsSnapshot []models.Stats

func (t StatsService) Get() (StatsSnapshot, error) {
	var out StatsSnapshot

	bs, err := t.get(stats)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", stats)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

type StatsHistorySnapshot []models.StatsHistory

func (t StatsService) GetHistory() (StatsHistorySnapshot, error) {
	var out StatsHistorySnapshot

	bs, err := t.get(statsHistory)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", statsHistory)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}

type StatsUSDSnapshot []*models.StatsUSD

func (t StatsService) GetUSD() (StatsUSDSnapshot, error) {
	var out StatsUSDSnapshot

	bs, err := t.get(statsHistoryUSD)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get %s", statsHistoryUSD)
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal json")
	}

	return out, nil
}
