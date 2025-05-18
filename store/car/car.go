package store

import (
	"car-zone/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) GetCarByID(ctx context.Context, id string) (models.Car, error) {
	var car models.Car

	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, 
	c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range FROM car c
	LEFT JOIN engine e ON c.engine_id = e.id WHERE c.id = $1`

	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&car.ID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Engine.EngineID,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngineID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCylinders,
		&car.Engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}

	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string

	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, 
	c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range FROM car c
	LEFT JOIN engine e ON c.engine_id = e.id WHERE c.brand = $1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car c WHERE brand = $1`
	}

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if isEngine {
			err := rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
				&car.Engine.EngineID,
				&car.Engine.Displacement,
				&car.Engine.NoOfCylinders,
				&car.Engine.CarRange,
			)

			if err != nil {
				return nil, err
			}
		} else {
			err := rows.Scan(
				&car.ID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
			)

			if err != nil {
				return nil, err
			}
		}

		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var engineID uuid.UUID

	fmt.Println("Received engine_id:", carReq.Engine.EngineID)
	err := s.db.QueryRowContext(
		ctx,
		"SELECT id FROM engine WHERE id = $1",
		carReq.Engine.EngineID,
	).Scan(&engineID)

	if err != nil {
		if err == sql.ErrNoRows {
			return createdCar, errors.New("engine_id does not exists is the engine table")
		}
		return createdCar, err
	}

	carID := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	newCar := models.Car{
		ID:        carID,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine,
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		err = tx.Commit()
	}()

	query := `INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = tx.ExecContext(
		ctx,
		query,
		newCar.ID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngineID,
		newCar.Price,
		newCar.CreatedAt,
		newCar.UpdatedAt,
	)

	if err != nil {
		return createdCar, err
	}

	if err = tx.Commit(); err != nil {
		return createdCar, err
	}

	createdCar = newCar
	return createdCar, nil
}

func (s Store) UpdateCar(ctx context.Context, id string, carUpdate models.Car) (models.Car, error) {
	var updatedCar models.Car

	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM car WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return updatedCar, err
	}

	if !exists {
		return updatedCar, errors.New("car with specified ID does not exist")
	}

	if carUpdate.Engine.EngineID != uuid.Nil {
		var engineExists bool
		err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM engine WHERE id = $1)",
			carUpdate.Engine.EngineID).Scan(&engineExists)
		if err != nil {
			return updatedCar, err
		}

		if !engineExists {
			return updatedCar, errors.New("specified engine_id does not exist in the engine table")
		}
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	carUpdate.UpdatedAt = time.Now()

	query := `
        UPDATE car
        SET name = $1, year = $2, brand = $3, fuel_type = $4, 
            engine_id = $5, price = $6, updated_at = $7
        WHERE id = $8
        RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at
    `

	err = tx.QueryRowContext(
		ctx,
		query,
		carUpdate.Name,
		carUpdate.Year,
		carUpdate.Brand,
		carUpdate.FuelType,
		carUpdate.Engine.EngineID,
		carUpdate.Price,
		carUpdate.UpdatedAt,
		id,
	).Scan(
		&updatedCar.ID,
		&updatedCar.Name,
		&updatedCar.Year,
		&updatedCar.Brand,
		&updatedCar.FuelType,
		&updatedCar.Engine.EngineID,
		&updatedCar.Price,
		&updatedCar.CreatedAt,
		&updatedCar.UpdatedAt,
	)

	if err != nil {
		return updatedCar, err
	}

	if err = tx.Commit(); err != nil {
		return updatedCar, err
	}

	if updatedCar.Engine.EngineID != uuid.Nil {
		engineQuery := `
            SELECT id, displacement, no_of_cylinders, car_range 
            FROM engine WHERE id = $1
        `

		err = s.db.QueryRowContext(
			ctx,
			engineQuery,
			updatedCar.Engine.EngineID,
		).Scan(
			&updatedCar.Engine.EngineID,
			&updatedCar.Engine.Displacement,
			&updatedCar.Engine.NoOfCylinders,
			&updatedCar.Engine.CarRange,
		)

		if err != nil {
			return updatedCar, err
		}
	}

	return updatedCar, nil
}

func (s Store) DeleteCar(ctx context.Context, id string) error {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM car WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("car with specified ID does not exist")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM car WHERE id = $1", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
