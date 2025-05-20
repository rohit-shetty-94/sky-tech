package handler

import (
	"net/http"
	"net/http/httptest"
	"sky-tech/entity"
	"sky-tech/repository"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetMetrics(t *testing.T) {
	e := echo.New()

	repo := &repository.MockMetricRepository{
		Data: []entity.Metric{
			{Timestamp: 1716200000, CPULoad: 45.5, Concurrency: 1200},
		},
	}

	h := New(repo)

	req := httptest.NewRequest(http.MethodGet, "/metrics?start=1716000000&end=1716300000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.GetMetrics(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
