package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func CreateDatabasePoolConnections(dbstring string) (*pgxpool.Pool, error) {
	poolCfg, _ := pgxpool.ParseConfig(dbstring)
	poolCfg.ConnConfig.PreferSimpleProtocol = true //disable prepared queries for pgbouncer

	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "err init connection for new pool")
	}
	return pool, nil
}
