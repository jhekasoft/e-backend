package internal

import (
	"e-backend/internal/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BuildTime is time when executable was built
var BuildTime string = "unknown"
var Version string = "0.0.0"

type HTTPApp struct {
	Config models.Config
	Echo   *echo.Echo
}

func NewHTTPApp(config models.Config) HTTPApp {
	return HTTPApp{Config: config}
}

func (a *HTTPApp) Run() {
	a.Echo = echo.New()
	a.Echo.HideBanner = true
	a.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	a.Echo.Logger.Fatal(a.Echo.Start(fmt.Sprintf(":%d", a.Config.HTTP.Port)))
}
