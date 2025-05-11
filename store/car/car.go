package store

import (
	"car-zone/models"
	"context"
	"database/sql"
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
	//
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	//
}

func (s Store) CreateCar(ctx context.Context, car models.Car) (models.Car, error) {
	//
}

func (s Store) UpdateCar(ctx context.Context, id string, car models.Car) (models.Car, error) {
	//
}

func (s Store) DeleteCar(ctx context.Context, id string) error {
	//
}
