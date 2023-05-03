package main

import (
	"github.com/calculi-corp/config"
	"github.com/calculi-corp/hippo-service/internal"
	"github.com/calculi-corp/log"
	"os"
)

const (
	ConfigLogLevel = "log.level"
)

func main() {
	c := initConfiguration()

	err := internal.Start(c)
	if err != nil {
		log.Fatal("error while starting server", err)
		os.Exit(1)
	}
}

func initConfiguration() *internal.Configuration {

	// Server config
	config.Config.DefineStringFlag(ConfigLogLevel, "info", "Log level, choose between debug, info, warn, error, fatal")

	// gRPC downstream services config
	config.Config.DefineStringFlag("potato-endpoint", "localhost:9090", "URL of the potato-service")

	err := config.Config.SetCliFlags()
	if log.CheckErrorf(err, "Unable to set flags") {
		os.Exit(1)
	}
	log.Infof("Flag values: %s", config.Config.FlagValues())

	res := &internal.Configuration{}
	res.LogLevel = config.Config.GetString(ConfigLogLevel)

	res.PotatoServiceEndpoint = config.Config.GetString("potato-endpoint")

	return res
}
