package port

import "github.com/labstack/echo/v4"

type (
	IOpinionHandler interface {
		Find(c echo.Context) error
		FindByProductId(c echo.Context) error
		CreateOpinion(c echo.Context) error
	}

	IProductHandler interface {
		Find(c echo.Context) error
		CreateProduct(c echo.Context) error
	}

	ISummaryHandler interface {
		Positive(c echo.Context) error
	}
)
