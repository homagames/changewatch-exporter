package main

import (
    "log/slog"
	"os"
	"flag"
	"github.com/homagames/hg-infra-change-exporter/command"
	"github.com/homagames/hg-infra-change-exporter/monitor"
)


func main() {
    // Defined loggers
	// TODO: configure loglevel: https://pkg.go.dev/log/slog#hdr-Levels
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Read config file from args
	configFile := flag.String("config", "config.json", "path to config file")
	// Parse args
	flag.Parse()

	v, err := command.LoadConfig(*configFile)

	if err != nil {
		logger.Error("Error reading config file", err)
		// Exit with error
		os.Exit(1)
	}

	logger.Info("Reading config file")
	logger.Info("Config file", "config", v.AllSettings())

	// Get monitors from config file
	monitors := monitor.LoadMonitors(v.AllSettings())

	logger.Info("Monitors", "monitors", monitors)
}