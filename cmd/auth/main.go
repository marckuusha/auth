package main

import (
	"context"
	"os"

	"github.com/marckuusha/auth/internal/controller"
	"github.com/marckuusha/auth/internal/db"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	dbstring := os.Getenv("PG_STRING")
	if dbstring == "" {
		logger.Fatalf("cannot get PG_STRING")
	}

	// postgres://habrpguser:pgpwd4habr@postgres:5432/habrdb?sslmode=disable

	dbconn, err := db.CreateDatabasePoolConnections(dbstring)
	if err != nil {
		logger.Fatalf("cannot connect to db: %s", err)
	}
	name := ""
	err = dbconn.QueryRow(context.Background(), `select name from users where id = $1`, 1).Scan(&name)
	if err != nil {
		logger.Fatalf("cannot get user: %s", err)
	}
	logger.Infof("name %s", name)

	serv := controller.NewServer(logger)

	serv.Run()
}
