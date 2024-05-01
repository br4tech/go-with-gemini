package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
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

func (h *SummaryHandler) Find(c echo.Context) error {
	id := c.Param("id")

	summaryID, err := strconv.Atoi(id)

	if err != nil {
		return HandlerResponse(c, http.StatusBadRequest, "Invalid Summary ID")
	}

	summary, err := h.summaryUseCase.Find(summaryID)
	if err != nil {
		return HandlerResponse(c, http.StatusNotFound, "Summary not found")
	}

	summaryJSON, err := json.Marshal(summary)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal Summary")
	}

	return HandlerResponse(c, http.StatusOK, string(summaryJSON))
}

func (h *SummaryHandler) CreateSummary(c echo.Context) error {
	reqBody := new(dto.SummarytDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	summary, err := h.summaryUseCase.CreateSummary(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to created Summary")
	}

	summaryJSON, err := json.Marshal(summary)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal Summary")
	}

	return HandlerResponse(c, http.StatusCreated, string(summaryJSON))
}
