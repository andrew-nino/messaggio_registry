package postgresdb

import (
	"fmt"

	"github.com/andrew-nino/messaggio/internal/domain/models"
)

func (p *Postgres) SetApproval(answer models.Answer) error {

	var checkID int

	query := fmt.Sprintf(`UPDATE %s SET approval=$1  WHERE id = $2 RETURNING id`, clientsTable)
	row := p.db.QueryRow(query, answer.Approve, answer.ID)
	err := row.Scan(&checkID)
	if err != nil {
		p.log.Debugf("repository.SetApproval - row.Scan : %v", err)
		return err
	}

	return nil
}
