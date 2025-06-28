package service

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GoCarZone/internal/domain"
	"github.com/saleh-ghazimoradi/GoCarZone/internal/dto"
	"github.com/saleh-ghazimoradi/GoCarZone/internal/repository"
)

type EngineService interface {
	GetEngineById(ctx context.Context, id string) (*domain.Engine, error)
	CreateEngine(ctx context.Context, input *dto.Engine) (*domain.Engine, error)
	UpdateEngine(ctx context.Context, id string, input *dto.UpdateEngine) (*domain.Engine, error)
	DeleteEngine(ctx context.Context, id string) error
}

type engineService struct {
	engineRepository repository.EngineRepository
}

func (e *engineService) GetEngineById(ctx context.Context, id string) (*domain.Engine, error) {
	return e.engineRepository.GetEngineById(ctx, id)
}

func (e *engineService) CreateEngine(ctx context.Context, input *dto.Engine) (*domain.Engine, error) {
	if err := input.ValidateEngine(); err != nil {
		return nil, fmt.Errorf("validate engine input: %w", err)
	}

	return e.engineRepository.CreateEngine(ctx, &domain.Engine{
		Displacement:  input.Displacement,
		NoOfCylinders: input.NoOfCylinders,
		CarRange:      input.CarRange,
	})
}

func (e *engineService) UpdateEngine(ctx context.Context, id string, input *dto.UpdateEngine) (*domain.Engine, error) {
	if err := input.ValidateUpdateEngine(); err != nil {
		return nil, fmt.Errorf("invalid update engine data: %w", err)
	}

	existingEngine, err := e.engineRepository.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Displacement != nil {
		existingEngine.Displacement = *input.Displacement
	}
	if input.NoOfCylinders != nil {
		existingEngine.NoOfCylinders = *input.NoOfCylinders
	}
	if input.CarRange != nil {
		existingEngine.CarRange = *input.CarRange
	}

	return e.engineRepository.UpdateEngine(ctx, id, existingEngine)
}

func (e *engineService) DeleteEngine(ctx context.Context, id string) error {
	return e.engineRepository.DeleteEngine(ctx, id)
}

func NewEngineService(engineRepository repository.EngineRepository) EngineService {
	return &engineService{
		engineRepository: engineRepository,
	}
}
