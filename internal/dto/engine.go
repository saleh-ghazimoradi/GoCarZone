package dto

import (
	"errors"
	"github.com/google/uuid"
)

type Engine struct {
	EngineId      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
}

type UpdateEngine struct {
	Displacement  *int64 `json:"displacement"`
	NoOfCylinders *int64 `json:"no_of_cylinders"`
	CarRange      *int64 `json:"car_range"`
}

func (e *Engine) ValidateEngine() error {
	var errs []error
	if e.EngineId == uuid.Nil {
		errs = append(errs, errors.New("engine id cannot be empty"))
	}
	if e.Displacement <= 0 {
		errs = append(errs, errors.New("engine displacement must be greater than zero"))
	}
	if e.Displacement > 10000 {
		errs = append(errs, errors.New("engine displacement must not exceed 10000 cc"))
	}
	if e.NoOfCylinders <= 0 {
		errs = append(errs, errors.New("engine number of cylinders must be greater than zero"))
	}
	if e.NoOfCylinders > 16 {
		errs = append(errs, errors.New("engine number of cylinders must not exceed 16"))
	}
	if e.CarRange <= 0 {
		errs = append(errs, errors.New("engine car range must be greater than zero"))
	}
	if e.CarRange > 1000 {
		errs = append(errs, errors.New("engine car range must not exceed 1000 km"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (e *UpdateEngine) ValidateUpdateEngine() error {
	var errs []error
	if e.Displacement != nil {
		if *e.Displacement <= 0 {
			errs = append(errs, errors.New("engine displacement must be greater than zero"))
		}
		if *e.Displacement > 10000 {
			errs = append(errs, errors.New("engine displacement must not exceed 10000 cc"))
		}
	}
	if e.NoOfCylinders != nil {
		if *e.NoOfCylinders <= 0 {
			errs = append(errs, errors.New("engine number of cylinders must be greater than zero"))
		}
		if *e.NoOfCylinders > 16 {
			errs = append(errs, errors.New("engine number of cylinders must not exceed 16"))
		}
	}
	if e.CarRange != nil {
		if *e.CarRange <= 0 {
			errs = append(errs, errors.New("engine car range must be greater than zero"))
		}
		if *e.CarRange > 1000 {
			errs = append(errs, errors.New("engine car range must not exceed 1000 km"))
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
