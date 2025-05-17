package main

import (
	"car-zone/driver"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	carService "car-zone/service/car"
	carStore "car-zone/store/car"

	engineService "car-zone/service/engine"
	engineStore "car-zone/store/engine"

	carHandler "car-zone/handler/car"
	engineHandler "car-zone/handler/engine"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		fmt.Printf("Error executing schema file: %v\n", err)
		return
	}

	router.HandleFunc("/api/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/api/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/api/cars/{id}", carHandler.DeleteCar).Methods("DELETE")
	router.HandleFunc("/api/cars/brand/{brand}", carHandler.GetCarByBrand).Methods("GET")

	router.HandleFunc("/api/engines", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/api/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/api/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/api/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}

	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

func executeSchemaFile(db *sql.DB, schemaFile string) error {
	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}
