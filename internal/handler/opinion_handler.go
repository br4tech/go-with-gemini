package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	"github.com/labstack/echo/v4"
)

type OpinionHandler struct {
	opinionUseCase port.IOpinionUseCase
}

func NewOpinionHandler(opinionUseCase port.IOpinionUseCase) *OpinionHandler {
	return &OpinionHandler{
		opinionUseCase: opinionUseCase,
	}
}

func (h *OpinionHandler) Find(c echo.Context) error {
	id := c.Param("id")

	opinionID, err := strconv.Atoi(id)

	if err != nil {
		return HandlerResponse(c, http.StatusBadRequest, "Invalid opinion ID")
	}

	opinion, err := h.opinionUseCase.Find(opinionID)
	if err != nil {
		return HandlerResponse(c, http.StatusNotFound, "Opinion not found")
	}

	opinionJSON, err := json.Marshal(opinion)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal opinion")
	}

	return HandlerResponse(c, http.StatusOK, string(opinionJSON))
}

func (h *OpinionHandler) CreateOpinion(c echo.Context) error {
	reqBody := new(dto.OpinionDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	opinion, err := h.opinionUseCase.CreateOpinion(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to created opinion")
	}

	opinionJSON, err := json.Marshal(opinion)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal opinion")
	}

	return HandlerResponse(c, http.StatusCreated, string(opinionJSON))
}
