package service

import (
	"github.com/andrew-nino/messaggio/internal/domain/models"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Clients interface {
	RegisterClientOnRepo(models.Client) (int, error)
	UpdateClientOnRepo(models.Client) error
	GetClientFromRepo(id int) (models.Client, error)
	DeleteClientOnRepo(id int) error
	GetStatisticOnRepo() (models.Statistic, error)
}

type ApprovalService interface {
	SetApproval(models.Answer) error
}

type Sender interface {
	SendToBroker(int, models.Client) error
}

type ApplicationServices struct {
	log      *logrus.Logger
	clients  Clients
	approval ApprovalService
	sender   Sender
}

func New(log *logrus.Logger, sender Sender, clients Clients, approval ApprovalService) *ApplicationServices {
	return &ApplicationServices{
		log:      log,
		clients:  clients,
		approval: approval,
		sender:   sender,
	}
}
