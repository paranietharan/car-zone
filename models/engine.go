package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return fmt.Errorf("engine ID cannot be empty")
	}

	if engine.Displacement <= 0 {
		return fmt.Errorf("displacement must be greater than 0")
	}

	if engine.NoOfCylinders <= 0 {
		return fmt.Errorf("number of cylinders must be greater than 0")
	}

	if engine.CarRange <= 0 {
		return fmt.Errorf("car range must be greater than 0")
	}
	return nil
}
