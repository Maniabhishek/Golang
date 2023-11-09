package apis

import "github.com/labstack/echo"

type healthCheckApi struct {
}

func NewHealthCheckAPI() *healthCheckApi {
	return &healthCheckApi{}
}

func (h *healthCheckApi) HealthAPI(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{
		"message": "ok",
	})
}
