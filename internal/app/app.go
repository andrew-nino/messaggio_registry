package app

import (
	"github.com/andrew-nino/messaggio/config"
	httpserver "github.com/andrew-nino/messaggio/internal/app/httpserver"
	postgres "github.com/andrew-nino/messaggio/internal/repository/postgres"
	producer "github.com/andrew-nino/messaggio/internal/service/producer"
	service "github.com/andrew-nino/messaggio/internal/service/registry"
	"github.com/sirupsen/logrus"
)

type App struct {
	HTTPServer *httpserver.Server
}

func NewApplication(log *logrus.Logger, port string, cfg *config.Config) *App {

	repository := postgres.New(log, &cfg.PG)
	sender := producer.New(log, &cfg.Kafka)
	services := service.New(log, sender, repository, repository)
	server := httpserver.New(log, services, port)

	return &App{
		HTTPServer: server,
	}
}
