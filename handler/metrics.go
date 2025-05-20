package handler

import (
	"net/http"
	"sky-tech/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repo repository.MetricRepository
}

func New(repo repository.MetricRepository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetMetrics(c echo.Context) error {
	start, _ := strconv.ParseInt(c.QueryParam("start"), 10, 64)
	end, _ := strconv.ParseInt(c.QueryParam("end"), 10, 64)

	metrics, err := h.Repo.GetMetrics(c.Request().Context(), start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, metrics)
}
