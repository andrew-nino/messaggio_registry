package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Registry interface {
	RegisterClient() error
}

type Handler struct {
	log      *logrus.Logger
	services Registry
}

func NewHandler(log *logrus.Logger, services Registry) *Handler {
	return &Handler{
		log:      log,
		services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/client")
	{
		auth.POST("/add", h.addClient)
	}

	return router
}
