package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase port.IProductUseCase
}

func NewProductHandler(productUseCase port.IProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (h *ProductHandler) Find(c echo.Context) error {
	id := c.Param("id")

	productID, err := strconv.Atoi(id)

	if err != nil {
		return HandlerResponse(c, http.StatusBadRequest, "Invalid Product ID")
	}

	product, err := h.productUseCase.Find(productID)
	if err != nil {
		return HandlerResponse(c, http.StatusNotFound, "Product not found")
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal Product")
	}

	return HandlerResponse(c, http.StatusOK, string(productJSON))
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	reqBody := new(dto.ProductDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	product, err := h.productUseCase.CreateProduct(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to created Product")
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to marshal Product")
	}

	return HandlerResponse(c, http.StatusCreated, string(productJSON))
}
