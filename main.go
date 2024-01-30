package main

import (
	"flag"
	"fmt"
	"github.com/homagames/changewatch-exporter/command"
	"github.com/homagames/changewatch-exporter/logger"
	"github.com/homagames/changewatch-exporter/monitor"
	"os"
	"time"
	// See https://prometheus.io/docs/guides/go-application/
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// Defined loggers
	logger := logger.GetLogger()

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

	settings := v.AllSettings()
	interval := settings["interval"].(int)

	// Starting prometheus metrics server in background
	prom_port := settings["prometheus"].(map[string]interface{})["port"].(int)
	prom_path := settings["prometheus"].(map[string]interface{})["path"].(string)
	go func() {
		http.Handle(prom_path, promhttp.Handler())
		err := http.ListenAndServe(fmt.Sprintf(":%d", prom_port), nil)
		if err != nil {
			logger.Error("Error starting prometheus metrics server", err)
			// Exit with error
			os.Exit(1)
		}
	}()
	logger.Info("Prometheus metrics server started", "port", prom_port, "path", prom_path)

	// Get monitors from config file
	monitors := monitor.LoadMonitors(settings)

	logger.Info("Monitors", "monitors", monitors)

	for {
		// Loop through monitors
		for _, m := range monitors {
			logger.Info("Monitor", "monitor", m)
			err := m.Process()
			if err != nil {
				logger.Error("Error processing monitor", err)
			}
		}
		// Sleep for 60 seconds
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
