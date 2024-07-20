package service

import (
	"github.com/andrew-nino/messaggio/internal/domain/models"
	"github.com/sirupsen/logrus"
)

type Clients interface {
	RegisterClientOnRepo(models.Client) (int, error)
	UpdateClientOnRepo(models.Client) error
	DeleteClientOnRepo(id int) error
}

type ApprovalService interface {
	SetApproval(models.Answer) error
}

type ApplicationServices struct {
	log      *logrus.Logger
	clients  Clients
	approval ApprovalService
}

func New(log *logrus.Logger, clients Clients, approval ApprovalService) *ApplicationServices {
	return &ApplicationServices{
		log:      log,
		clients:  clients,
		approval: approval,
	}
}
