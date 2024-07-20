package postgresdb

func (p *Postgres) AddClient(id string) error {
	p.log.Infof("Postgres lient id = %s is successful", id)
	return nil
}
