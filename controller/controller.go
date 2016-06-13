package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ngc224/phck/config"
	"github.com/ngc224/phck/model"
)

type Controller struct {
	Health model.Health
}

func NewController(c *config.Config) *Controller {
	return &Controller{
		Health: c.Health,
	}
}

func (cntr *Controller) HealthCheck(c echo.Context) error {
	code := HealthStatusCode(cntr.Health.IsHealth())
	return c.JSON(code, cntr.Health)
}

func HealthStatusCode(ok bool) int {
	if ok {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}
