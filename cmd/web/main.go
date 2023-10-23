package main

import (
	"flag"
	"fmt"
	"gg/pkg/app"
	"gg/pkg/config"
	"os"
)

const (
	defaultConfigPath = "conf/config.yml"
)

func main() {
	configPath := flag.String("config", defaultConfigPath, "config file")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Config loaded: %+v\n", config)

	_, err = app.ListenAndServe(config)
	if err != nil {
		fmt.Printf("Error starting web service: %s\n", err)
		os.Exit(1)
	}
}
