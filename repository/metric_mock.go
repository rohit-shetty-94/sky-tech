package repository

import (
	"context"
	"sky-tech/entity"
)

type MockMetricRepository struct {
	Data []entity.Metric
}

func (m *MockMetricRepository) GetMetrics(ctx context.Context, start, end int64) ([]entity.Metric, error) {
	return m.Data, nil
}
