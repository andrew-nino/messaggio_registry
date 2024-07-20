package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/andrew-nino/messaggio/config"
	"github.com/andrew-nino/messaggio/internal/app"
)

func main() {

	cfg := config.NewConfig()

	log := SetLogrus(cfg.Log.Level)

	application := app.NewApplication(log, cfg.HTTP.Port, cfg)

	go application.HTTPServer.MustRun()

	log.Print("App " + cfg.App.Name + " version: " + cfg.App.Version + " Started")

	// Waiting signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}
