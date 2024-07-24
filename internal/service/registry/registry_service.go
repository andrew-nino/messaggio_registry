package service

import "github.com/andrew-nino/messaggio/internal/domain/models"

func (s *ApplicationServices) RegisterClient(client models.Client) (int, error) {

	id, err := s.clients.RegisterClientOnRepo(client)
	if err != nil {
		s.log.Error("Error registering client to BD: ", err)
		return 0, err
	}

	go func() {
		err = s.sender.SendToBroker(id, client)
		if err != nil {
			s.log.Error("Error sending data to broker: ", err)
		}
	}()

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

func (s *ApplicationServices) GetClient(clientID int) (models.Client, error) {

	client, err := s.clients.GetClientFromRepo(clientID)
    if err!= nil {
        s.log.Errorf("Error getting client %d from BD: %v", clientID, err)
        return models.Client{}, err
    }
    return client, nil
}

func (s *ApplicationServices) DeleteClient(clientID int) error {

	err := s.clients.DeleteClientOnRepo(clientID)
	if err != nil {
		s.log.Error("Error deleting client from BD: ", err)
		return err
	}
	return nil
}

func(s *ApplicationServices) GetStatistic() (models.Statistic, error){
	statistic, err := s.clients.GetStatisticOnRepo()
    if err!= nil {
        s.log.Error("Error getting statistic from BD: ", err)
        return models.Statistic{}, err
    }
    return statistic, nil
}
