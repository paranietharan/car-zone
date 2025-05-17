package store

import (
	"car-zone/models"
	"context"
)

type CarStoreInterface interface {
	GetCarByID(ctx context.Context, id string) (models.Car, error)
	CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq models.Car) (models.Car, error) // Changed parameter type from CarRequest to Car
	DeleteCar(ctx context.Context, id string) error
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
}

type EngineStoreInterface interface {
	EngineByID(ctx context.Context, id string) (models.Engine, error)
	CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq models.EngineRequest) (models.Engine, error)
	EngineDelete(ctx context.Context, id string) (models.Engine, error)
}
