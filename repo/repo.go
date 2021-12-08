package repo

import (
	"github.com/ISS-Dating/service-analyzer/model"
)

type Interface interface {
	ReadUserStatsById(id uint64) (model.Stats, error)
	UpdateUserStats(stats model.Stats) error
}
