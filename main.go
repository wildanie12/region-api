package main

import (
	"github.com/joho/godotenv"
	"github.com/wildanie12/region-api/config"
	appHttp "github.com/wildanie12/region-api/internal/http"
	"github.com/wildanie12/region-api/internal/http/handler"

	_ "embed"
)

//go:embed .version
var version string

func main() {
	// config initiation
	_ = godotenv.Load()
	cfg := config.Get(version)

	// -- repository factory dependency injection. construct your repository here...

	// -- service factory dependency injection. construct your service here...

	// app http tranposrt initiation
	appHttp := appHttp.NewAppHTTP(cfg.HTTP.Port)
	appHttp.Provide(
		handler.NewMetricHandler(),
	)
	appHttp.Start()
}
