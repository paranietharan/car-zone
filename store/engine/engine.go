package engine

import (
	errorutils "car-zone/error"
	"car-zone/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	engineQuery := `SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1`
	err = tx.QueryRowContext(
		ctx,
		engineQuery,
		id,
	).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return engine, errorutils.ErrEngineNotFound
		}
		return engine, err
	}

	return engine, nil
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	engineId := uuid.New()
	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO engine (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)",
		engineId, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return engine, errorutils.ErrEngineCreateFailed
		}
		return engine, err
	}

	err = tx.Commit()
	if err != nil {
		return engine, err
	}
	engine = models.Engine{
		EngineID:      engineId,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	return engine, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineReq models.EngineRequest) (models.Engine, error) {
	var engine models.Engine

	engineId, err := uuid.Parse(id)
	if err != nil {
		return engine, errorutils.ErrEngineInvalidID
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(
		ctx,
		"UPDATE engine SET displacement = $1, no_of_cylinders = $2, car_range = $3 WHERE id = $4",
		engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, engineId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return engine, errorutils.ErrEngineUpdateFailed
		}
		return engine, err
	}
	err = tx.Commit()
	if err != nil {
		return engine, err
	}

	engine = models.Engine{
		EngineID:      engineId,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	engineId, err := uuid.Parse(id)
	if err != nil {
		return engine, errorutils.ErrEngineInvalidID
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	engineQuery := `SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id = $1`
	err = tx.QueryRowContext(
		ctx,
		engineQuery,
		engineId,
	).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return engine, errorutils.ErrEngineNotFound
		}
		return engine, err
	}

	deleteQuery := `DELETE FROM engine WHERE id = $1`
	_, err = tx.ExecContext(ctx, deleteQuery, engineId)
	if err != nil {
		return engine, errorutils.ErrEngineDeleteFailed
	}

	err = tx.Commit()
	if err != nil {
		return engine, err
	}

	return engine, nil
}
