package app

import (
	"github.com/andrew-nino/messaggio/config"
	httpserver "github.com/andrew-nino/messaggio/internal/app/httpserver"
	postgres "github.com/andrew-nino/messaggio/internal/repository/postgres"
	service "github.com/andrew-nino/messaggio/internal/service/registry"
	"github.com/sirupsen/logrus"
)

type App struct {
	HTTPServer *httpserver.Server
}

func NewApplication(log *logrus.Logger, port string, cfgPg *config.PG) *App {

	repository := postgres.New(log, cfgPg)
	services := service.New(log, repository)
	server := httpserver.New(log, services, port)

	return &App{
		HTTPServer: server,
	}
}
