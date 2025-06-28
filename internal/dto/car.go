package dto

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	MinValidYear = 1886
)

type Car struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func (c *Car) validateName() error {
	if c.Name == "" {
		return errors.New("car name cannot be empty")
	}
	return nil
}

func (c *Car) validateYear() error {
	if c.Year == "" {
		return errors.New("car year cannot be empty")
	}
	yearInt, err := strconv.Atoi(c.Year)
	if err != nil {
		return errors.New("car year must be a valid number")
	}
	currentYear := time.Now().Year()
	if yearInt < MinValidYear || yearInt > currentYear {
		return errors.New("car year must be between 1886 and the current year")
	}
	return nil
}

func (c *Car) validateBrand() error {
	if c.Brand == "" {
		return errors.New("car brand cannot be empty")
	}
	return nil
}

func (c *Car) validateFuelType() error {
	validFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validFuelTypes {
		if strings.EqualFold(c.FuelType, validType) {
			return nil
		}
	}
	return errors.New("car fuel type must be one of: Petrol, Diesel, Electric, Hybrid")
}

func (c *Car) validatePrice() error {
	if c.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

func (c *Car) ValidateCar() error {
	var errs []error
	if err := c.validateName(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateYear(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateBrand(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validateFuelType(); err != nil {
		errs = append(errs, err)
	}
	if err := c.validatePrice(); err != nil {
		errs = append(errs, err)
	}
	if err := c.Engine.validateEngine(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
