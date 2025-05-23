package service

import (
	"car-zone/models"
	"context"
)

type CarServiceInterface interface {
	GetCarByID(ctx context.Context, id string) (models.Car, error)
	CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq models.Car) (models.Car, error)
	DeleteCar(ctx context.Context, id string) error
	GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineByID(ctx context.Context, id string) (models.Engine, error)
	CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq models.EngineRequest) (models.Engine, error)
	EngineDelete(ctx context.Context, id string) (models.Engine, error)
}
