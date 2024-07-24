package postgresdb

import (
	"fmt"
	"strings"

	"github.com/andrew-nino/messaggio/internal/domain/models"
)

func (p *Postgres) RegisterClientOnRepo(client models.Client) (int, error) {

	var clientID int
	query := fmt.Sprintf(`INSERT INTO %s (surname, name, patronymic, email) values ($1, $2, $3, $4) RETURNING id`, clientsTable)
	rowClient := p.db.QueryRow(query, client.Surname, client.Name, client.Patronymic, client.Email)
	err := rowClient.Scan(&clientID)
	if err != nil {
		p.log.Debugf("repository.RegisterClientOnRepo - rowClient.Scan : %v", err)
		return 0, err
	}
	return clientID, nil
}
func (p *Postgres) UpdateClientOnRepo(client models.Client) error {

	var checkID int

	setValues := make([]string, 0)

	if client.ID == 0 {
		return fmt.Errorf("client id must be provided")
	}
	if client.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname='%s'", client.Surname))
	}
	if client.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name='%s'", client.Name))
	}
	if client.Patronymic != "" {
		setValues = append(setValues, fmt.Sprintf("patronymic='%s'", client.Patronymic))
	}
	if client.Email != "" {
		setValues = append(setValues, fmt.Sprintf("email='%s'", client.Email))
	}
	setValues = append(setValues, "update_at=now()")

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s SET %s  WHERE id = $1 RETURNING id`, clientsTable, setQuery)

	row := p.db.QueryRow(query, client.ID)
	err := row.Scan(&checkID)
	if err != nil {
		p.log.Debugf("repository.UpdateClientOnRepo - row.Scan : %v", err)
		return err
	}

	return nil
}

func (p *Postgres) GetClientFromRepo(id int) (models.Client, error) {

	var client models.Client
    query := fmt.Sprintf(`SELECT id, surname, name, patronymic, email, approval FROM %s WHERE id = $1`, clientsTable)
    err := p.db.Get(&client, query, id)
    if err!= nil {
        p.log.Debugf("repository.GetClientFromRepo - Get : %v", err)
        return client, err
    }
    return client, nil
}

func (p *Postgres) DeleteClientOnRepo(id int) error {

	var checkID int
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 RETURNING id`, clientsTable)
	row := p.db.QueryRow(query, id)
	err := row.Scan(&checkID)
	if err != nil {
		p.log.Debugf("repository.DeleteClient - row.Scan : %v", err)
		return err
	}
	return nil
}

func (p *Postgres) GetStatisticOnRepo() (models.Statistic, error) {

	var statistic models.Statistic

    query := fmt.Sprintf(`SELECT count(*) as total_clients,
								 count(case when approval = 1 then 1 end) as approved,
								 count(case when approval = -1 then 1 end) as unapproved,
								 count(case when approval = 0 then 1 end) as waiting 
								 FROM %s`, clientsTable)

	row := p.db.QueryRow(query)
	err := row.Scan(&statistic.TotalClients, &statistic.Approved, &statistic.Unapproved, &statistic.Waiting)
	if err!= nil {
        p.log.Debugf("repository.GetStatisticOnRepo - row.Scan : %v", err)
        return statistic, err
    }
	return statistic, nil
}
