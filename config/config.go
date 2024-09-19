package config

import "os"

type AppConfig struct {
	AppVersion string
	HTTP       struct {
		Port string
	}
	MySQL struct {
		ConnectionString string
	}
}

var appConfig *AppConfig

// Get current config for the application.
func Get(version ...string) *AppConfig {
	if appConfig != nil {
		return appConfig
	}
	ac := initConfig()

	// version
	if len(version) > 0 {
		ac.AppVersion = version[0]
	}
	appConfig = ac
	return appConfig
}

// init config when application start.
func initConfig() *AppConfig {
	ac := AppConfig{}
	ac.HTTP.Port = os.Getenv("HTTP_PORT")
	return &ac
}
