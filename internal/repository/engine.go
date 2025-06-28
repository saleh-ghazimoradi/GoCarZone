package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/saleh-ghazimoradi/GoCarZone/internal/domain"
	"time"
)

type EngineRepository interface {
	GetEngineById(ctx context.Context, id string) (*domain.Engine, error)
	CreateEngine(ctx context.Context, engine *domain.Engine) (*domain.Engine, error)
	UpdateEngine(ctx context.Context, id string, engine *domain.Engine) (*domain.Engine, error)
	DeleteEngine(ctx context.Context, id string) error
	WithTx(tx *sql.Tx) EngineRepository
	Close() error
}

type engineRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (e *engineRepository) GetEngineById(ctx context.Context, id string) (*domain.Engine, error) {
	engine := &domain.Engine{}
	query := `SELECT id, displacement, no_of_cylinders, car_range, created_at, updated_at FROM engine WHERE id = $1`

	err := e.exec(false).QueryRowContext(ctx, query, id).Scan(
		&engine.EngineId,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
		&engine.CreatedAt,
		&engine.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("engine not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query engine by id: %w", err)
	}
	return engine, nil
}

func (e *engineRepository) CreateEngine(ctx context.Context, engine *domain.Engine) (*domain.Engine, error) {
	engine.EngineId = uuid.New()
	engine.CreatedAt = time.Now()
	engine.UpdatedAt = time.Now()

	query := `INSERT INTO engine (id, displacement, no_of_cylinders, car_range, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := e.exec(true).ExecContext(
		ctx,
		query,
		engine.EngineId,
		engine.Displacement,
		engine.NoOfCylinders,
		engine.CarRange,
		engine.CreatedAt,
		engine.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create engine: %w", err)
	}

	return engine, nil
}

func (e *engineRepository) UpdateEngine(ctx context.Context, id string, engine *domain.Engine) (*domain.Engine, error) {
	engine.UpdatedAt = time.Now()

	query := `UPDATE engine 
              SET displacement = $1, no_of_cylinders = $2, car_range = $3, updated_at = $4 
              WHERE id = $5`

	result, err := e.exec(true).ExecContext(
		ctx,
		query,
		engine.Displacement,
		engine.NoOfCylinders,
		engine.CarRange,
		engine.UpdatedAt,
		id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update engine: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("engine not found")
	}

	return engine, nil
}

func (e *engineRepository) DeleteEngine(ctx context.Context, id string) error {
	query := `DELETE FROM engine WHERE id = $1`

	result, err := e.exec(true).ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete engine: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("engine not found")
	}

	return nil
}

func (e *engineRepository) WithTx(tx *sql.Tx) EngineRepository {
	return &engineRepository{
		dbWrite: e.dbWrite,
		dbRead:  e.dbRead,
		tx:      tx,
	}
}

func (e *engineRepository) Close() error {
	var errs []error
	if e.dbWrite != nil {
		if err := e.dbWrite.Close(); err != nil {
			errs = append(errs, errors.New("failed to close write database"+err.Error()))
		}
		e.dbWrite = nil
	}

	if e.dbRead != nil {
		if err := e.dbRead.Close(); err != nil {
			errs = append(errs, errors.New("failed to close read database"+err.Error()))
		}
		e.dbRead = nil
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func NewEngineRepository(dbWrite *sql.DB, dbRead *sql.DB) EngineRepository {
	return &engineRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
