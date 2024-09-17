package http

import "github.com/labstack/echo/v4"

// AppHTTP serve as HTTP transport for the application
type AppHTTP struct {
	e    *echo.Echo
	port string
}

// NewAppHTTP construts concrete App HTTP transport object.
func NewAppHTTP(port string) *AppHTTP {
	return &AppHTTP{
		e:    echo.New(),
		port: port,
	}
}

func (ah *AppHTTP) Start() {
	ah.e.HideBanner = true
	ah.e.Logger.Fatal(ah.e.Start(":" + ah.port))
}
