package service

import (
	"github.com/sirupsen/logrus"
)

type UserSaver interface {
	AddClient(id string) error
}

type ApplicationServices struct {
	log       *logrus.Logger
	userSaver UserSaver
}

func New(log *logrus.Logger, userSaver UserSaver) *ApplicationServices {
	return &ApplicationServices{
		log:       log,
		userSaver: userSaver,
	}
}

func (s *ApplicationServices) RegisterClient() error {
	s.userSaver.AddClient("32")
	s.log.Info("Service RegisterClient is successful")
	return nil
}
