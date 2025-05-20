package repository

import (
	"context"
	"sky-tech/entity"

	"github.com/go-pg/pg/v10"
)

type MetricRepository interface {
	GetMetrics(ctx context.Context, start, end int64) ([]entity.Metric, error)
}

type metricRepo struct {
	db *pg.DB
}

func NewMetricRepository(db *pg.DB) MetricRepository {
	return &metricRepo{db: db}
}

func (r *metricRepo) GetMetrics(ctx context.Context, start, end int64) ([]entity.Metric, error) {
	var metrics []entity.Metric
	err := r.db.ModelContext(ctx, &metrics).
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Order("timestamp ASC").
		Select()
	return metrics, err
}
