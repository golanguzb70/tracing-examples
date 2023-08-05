package storage

import (
	"github.com/golanguzb70/tracing-examples/rest-api-database/config"
	"github.com/golanguzb70/tracing-examples/rest-api-database/pkg/db"
	"github.com/golanguzb70/tracing-examples/rest-api-database/pkg/logger"
	"github.com/golanguzb70/tracing-examples/rest-api-database/storage/postgres"
)

type StorageI interface {
	Postgres() postgres.PostgresI
}

type StoragePg struct {
	postgres postgres.PostgresI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg config.Config) StorageI {
	return &StoragePg{
		postgres: postgres.New(db, log, cfg),
	}
}

func (s *StoragePg) Postgres() postgres.PostgresI {
	return s.postgres
}
