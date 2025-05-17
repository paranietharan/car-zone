package main

import (
	"car-zone/driver"
	"fmt"

	carService "car-zone/service/car"
	carStore "car-zone/store/car"

	engineService "car-zone/service/engine"
	engineStore "car-zone/store/engine"

	carHandler "car-zone/handler/car"
	engineHandler "car-zone/handler/engine"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	if db == nil {
		fmt.Println("Database connection is nil")
		return
	}

	// Initialize the stores
	carStore := carStore.NewStore(db)
	carService := carService.NewCarService(carStore)

	EngineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(EngineStore)

	// Initialize the handlers
	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

}
