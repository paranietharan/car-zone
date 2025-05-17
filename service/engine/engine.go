package engine

import (
	"car-zone/models"
	"car-zone/store"
	"context"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (e *EngineService) GetEngineByID(ctx context.Context, id string) (models.Engine, error) {
	engine, err := e.store.EngineByID(ctx, id)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (e *EngineService) CreateEngine(ctx context.Context, engineReq models.EngineRequest) (models.Engine, error) {
	engine, err := e.store.CreateEngine(ctx, engineReq)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (e *EngineService) UpdateEngine(ctx context.Context, id string, engineReq models.EngineRequest) (models.Engine, error) {
	engine, err := e.store.UpdateEngine(ctx, id, engineReq)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (e *EngineService) EngineDelete(ctx context.Context, id string) (models.Engine, error) {

	engine, err := e.store.EngineDelete(ctx, id)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}
