package postgresdb

import (
	"fmt"
	"log"
	"time"

	"github.com/andrew-nino/messaggio/config"
	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	log          *logrus.Logger
	connAttempts int
	connTimeout  time.Duration
	db           *sqlx.DB
}

func New(log *logrus.Logger, cfg *config.PG) *Postgres {

	db, err := NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}

	return &Postgres{
		log:          log,
		connAttempts: cfg.ConnAttempts,
		connTimeout:  cfg.ConnTimeout,
		db:           db,
	}
}

// Causes the database to open and checks the connection. If the connection is established, returns a pointer to the database.
// Returns an error if the database has not opened or there is no connection.
func NewPostgresDB(cfg *config.PG) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	for cfg.ConnAttempts > 0 {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Postgres is trying to connect, attempts left: %d", cfg.ConnAttempts)
		time.Sleep(cfg.ConnTimeout)
		cfg.ConnAttempts--
	}

	return db, err
}
