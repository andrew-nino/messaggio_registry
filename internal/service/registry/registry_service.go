package service

import "github.com/andrew-nino/messaggio/internal/domain/models"

func (s *ApplicationServices) RegisterClient(client models.Client) (int, error) {

	id, err := s.clients.RegisterClientOnRepo(client)
	if err != nil {
		s.log.Error("Error registering client to BD: ", err)
		return 0, err
	}
	return id, nil
}

func (s *ApplicationServices) UpdateClient(client models.Client) error {

	err := s.clients.UpdateClientOnRepo(client)
	if err != nil {
		s.log.Error("Error updating client in BD: ", err)
		return err
	}
	return nil
}

func (s *ApplicationServices) DeleteClient(clientID int) error {

	err := s.clients.DeleteClientOnRepo(clientID)
	if err != nil {
		s.log.Error("Error deleting client from BD: ", err)
		return err
	}
	return nil
}
