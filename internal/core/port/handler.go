package port

import "github.com/labstack/echo"

type (
	IOpinionHandler interface {
		Find(c echo.Context) error
		CreateOpinion(c echo.Context) error
	}

	IProductHandler interface {
		Find(c echo.Context) error
		CreateOpinion(c echo.Context) error
	}

	ISummaryHandler interface {
		Find(c echo.Context) error
		CreateOpinion(c echo.Context) error
	}
)
