package handler

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Data interface{} `json:"data,omitempty"`
}

func HandlerResponse(c echo.Context, responseCode int, data interface{}) error {
	return c.JSON(responseCode, &BaseResponse{
		Data: data,
	})
}
