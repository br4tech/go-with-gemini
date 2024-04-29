package handler

import (
	"net/http"

	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	"github.com/labstack/echo/v4"
)

type TextHandler struct {
	textUseCase port.ITextUseCase
}

func NewTextHandler(textUseCase port.ITextUseCase) *TextHandler {
	return &TextHandler{
		textUseCase: textUseCase,
	}
}

func (h *TextHandler) CreateContext(c echo.Context) error {
	reqBody := new(dto.TextNotificationDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}
}
