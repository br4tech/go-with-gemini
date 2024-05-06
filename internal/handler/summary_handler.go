package handler

import (
	"net/http"
	"strconv"

	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/labstack/echo/v4"
)

type SummaryHandler struct {
	summaryUseCase port.ISummaryUseCase
}

func NewSummaryHandler(summaryUseCase port.ISummaryUseCase) *SummaryHandler {
	return &SummaryHandler{
		summaryUseCase: summaryUseCase,
	}
}

func (h *SummaryHandler) Positive(c echo.Context) error {
	product_id := c.Param("product_id")

	productID, err := strconv.Atoi(product_id)

	if err != nil {
		return HandlerResponse(c, http.StatusBadRequest, "Invalid Summary ID")
	}

	summary, err := h.summaryUseCase.Positive(productID)
	if err != nil {
		return HandlerResponse(c, http.StatusNotFound, "Summary not found")
	}

	return HandlerResponse(c, http.StatusOK, summary)
}
