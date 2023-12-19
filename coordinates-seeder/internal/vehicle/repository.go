package vehicle

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	TableVehicles = "vehicle"
)

type IVehicleRepository interface {
	GetAllVehicle() ([]Vehicle, error)
  RegisterVehicleData(context.Context, *RegisterVehicleRequest) error
}

type VehicleRepository struct {
	db *sqlx.DB
}

var _ IVehicleRepository = (*VehicleRepository)(nil)

func NewVehicleRepository(db *sqlx.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (r *VehicleRepository) RegisterVehicleData(ctx context.Context, req *RegisterVehicleRequest) error {
	query := `
    INSERT INTO %s
      (name, type, brand, build_date, last_longitude, last_latitude)
    VALUES
      ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, fmt.Sprintf(query, TableVehicles),
		req.Name,
		req.Type,
		req.Brand,
		req.BuildDate,
    0, // hack, set as 0 first
    0, // hack, set as 0 first
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *VehicleRepository) GetAllVehicle() ([]Vehicle, error) {
	return nil, nil
}
