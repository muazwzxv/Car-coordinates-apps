package vehicle

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	TableVehicles = ""
)

type IVehicleRepository interface {
	GetAllVehicle() ([]Vehicle, error)
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

func (r *VehicleRepository) RegisterVehicleData(ctx context.Context, req *Vehicle) error {
	query := `
    INSERT INTO %s
      (name, type, brand, build_date)
    VALUE
      ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, fmt.Sprintf(query, TableVehicles),
		req.Name,
		req.Type,
		req.Brand,
		req.BuildDate,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *VehicleRepository) GetAllVehicle() ([]Vehicle, error) {
	return nil, nil
}
