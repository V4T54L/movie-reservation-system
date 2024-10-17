package store

import (
	"context"
	"log"

	"github.com/V4T54L/movie-reservation-system/internals/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName = "pgx"

type postgresStore struct {
	dbUri string
	db    *sqlx.DB
}

var store *postgresStore

func GetPostgresStore() *postgresStore {
	if store == nil {
		uri := config.GetConfig().DBUri
		if uri == "" {
			log.Panic("Database URI not provided")
		}
		store = &postgresStore{
			dbUri: config.GetConfig().DBUri,
		}
	}
	return store
}

func (s *postgresStore) connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.dbUri)
	if err != nil {
		return err
	}

	s.db = dbx
	return nil
}

func (s *postgresStore) close() error {
	return s.db.Close()
}
