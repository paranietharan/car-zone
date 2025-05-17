package car

import (
	"car-zone/models"
	"car-zone/store"
	"context"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (c *CarService) GetCarByID(ctx context.Context, id string) (models.Car, error) {
	car, err := c.store.GetCarByID(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c *CarService) CreateCar(ctx context.Context, carReq models.CarRequest) (models.Car, error) {
	car, err := c.store.CreateCar(ctx, carReq)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c *CarService) UpdateCar(ctx context.Context, id string, carReq models.Car) (models.Car, error) {
	car, err := c.store.UpdateCar(ctx, id, carReq)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c *CarService) DeleteCar(ctx context.Context, id string) error {
	err := c.store.DeleteCar(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CarService) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	cars, err := c.store.GetCarByBrand(ctx, brand, false)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
