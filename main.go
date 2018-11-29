package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {

	// get config
	config := getConfig()

	// login
	student := login(config.CampusId, config.CampusPassword, config.SupinfoAPIKey)

	// get planning
	planning := getPlaning(*student, config.SupinfoAPIKey)

	// write .ics file
	createCalendar(&planning, config.OutputPath)
}

func getConfig() *AppConfig {
	var opts AppConfig
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	return &opts
}
