package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	var migrationPath, username, password, host, port, db_name, ssl_mode string

	flag.StringVar(&migrationPath, "migration-path", "", "Path to migration")
	flag.StringVar(&username, "username", "postgres", "User postgres")
	flag.StringVar(&password, "password", "", "Password for connecting to DB")
	flag.StringVar(&host, "host", "localhost", "Host for connecting to DB")
	flag.StringVar(&port, "port", "5432", "Port for connecting to DB")
	flag.StringVar(&db_name, "db_name", "postgres", "Name DB")
	flag.StringVar(&ssl_mode, "ssl_mode", "disable", "SSL Mode")
	flag.Parse()

	if migrationPath == "" {
		panic("migration-path is required")
	}
	if password == "" {
		panic("password is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		"postgres://"+username+":"+password+"@"+host+":"+port+"/"+db_name+"?sslmode="+ssl_mode+"")

	if err != nil {
		log.Fatalf("failed to migrate db: %s", err.Error())
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Print("no migrations to apply")
		} else {
			log.Fatalf("failed to apply migrations: %s", err.Error())
		}
	}
}
