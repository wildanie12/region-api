package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

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

	for {
		fmt.Print("Search keyword: ")
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		fmt.Println(color.ThisF(color.BLUE, "Searching %s...", input))
		data, err := rgRepo.JoinedSearch(context.TODO(), strings.TrimSpace(input), 50)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-20s| %-20s| %-32s| %-20s\n", "Desa", "Kecamatan", "Kabupaten", "Provinsi")
		fmt.Printf("--------------------------------------------------------------------------------\n")
		for _, row := range data {
			fmt.Printf("%-20s| %-20s| %-32s| %-20s\n", row.VillageName, row.District.DistrictName, row.District.Regency.RegencyName, row.District.Regency.Province.ProvinceName)
		}

		fmt.Println("exit? (y)")
		str := ""
		if _, _ = fmt.Scanln(&str); strings.ToLower(strings.TrimSpace(str)) == "y" {
			return
		}
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	return

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
