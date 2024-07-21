package v1

import (
	"github.com/andrew-nino/messaggio/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Registry interface {
	RegisterClient(models.Client) (int, error)
	UpdateClient(models.Client) error
	DeleteClient(id int) error
	GetStatistic() (models.Statistic, error)
}

type Approval interface {
	Approve(models.Answer) error
}

type Handler struct {
	log      *logrus.Logger
	services Registry
	approval Approval
}

func NewHandler(log *logrus.Logger, services Registry, approval Approval) *Handler {
	return &Handler{
		log:      log,
		services: services,
		approval: approval,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/client")
	{
		auth.POST("/add", h.addClient)
		auth.PUT("/update", h.updateClient)
		auth.DELETE("/delete/:id", h.deleteClient)
		auth.GET("/statistic", h.getStatistic)
	}

	approval := router.Group("/approval")
	{
		approval.POST("/", h.updateStatus)
	}

	return router
}
