package engine

import (
	"car-zone/models"
	"context"
	"database/sql"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{
		db: db,
	}
}

func (e EngineStore) EngineByID(ctx context.Context, id string) (models.Engine, error) {
	//
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error) {
	//
}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {
	//
}
