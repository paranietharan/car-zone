package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateCarRequest(carRequest CarRequest) error {
	if err := validateName(carRequest.Name); err != nil {
		return err
	}

	if err := validateYear(carRequest.Year); err != nil {
		return err
	}

	if err := validateBrand(carRequest.Brand); err != nil {
		return err
	}

	if err := validateFuelType(carRequest.FuelType); err != nil {
		return err
	}

	if err := validatePrice(carRequest.Price); err != nil {
		return err
	}

	if err := validateEngine(carRequest.Engine); err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return fmt.Errorf("year cannot be empty")
	}

	_, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("year must be a number")
	}

	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return fmt.Errorf("year must be between 1886 and %d", currentYear)
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return fmt.Errorf("brand cannot be empty")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	if fuelType == "" {
		return fmt.Errorf("fuel type cannot be empty")
	}

	fuelTypes := []string{"petrol", "diesel", "electric", "hybrid"}
	for _, validFuelType := range fuelTypes {
		// Convert both to lowercase for case-insensitive comparison
		// and trim any leading/trailing whitespace
		fuelType = strings.TrimSpace(strings.ToLower(fuelType))
		if fuelType == validFuelType {
			return nil
		}
	}
	return fmt.Errorf("invalid fuel type: %s", fuelType)
}

func validatePrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	return nil
}
