package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/wildanie12/region-api/config"
	appHttp "github.com/wildanie12/region-api/internal/http"
	"github.com/wildanie12/region-api/internal/http/handler"
	"github.com/wildanie12/region-api/internal/repository"
	"github.com/wildanie12/region-api/internal/service"
	"github.com/wildanie12/region-api/lib"
	"github.com/wildanie12/region-api/util/color"

	_ "embed"
)

//go:embed .version
var version string

func main() {
	// config initiation
	_ = godotenv.Load()
	cfg := config.Get(version)

	// lib initiation
	di := lib.NewDatabaseInstance()
	fmt.Print(color.This(color.YELLOW, "‣ Connecting to mysql database... "))
	di.ConnectMySQL()
	fmt.Println(color.This(color.GREEN, "[✔ Connected]"))

	// -- repository factory dependency injection. construct your repository here...
	rgRepo := repository.NewRegion(di.MySQL)

	// -- service factory dependency injection. construct your service here...
	rgSrv := service.NewRegion(rgRepo)

	// app http tranposrt initiation
	appHttp := appHttp.NewAppHTTP(cfg.HTTP.Port)
	appHttp.Provide(
		handler.NewMetricHandler(),
		handler.NewRegion(rgSrv),
	)
	appHttp.Start()
}
