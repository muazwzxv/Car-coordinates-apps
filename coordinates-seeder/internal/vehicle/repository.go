package vehicle

import "github.com/jmoiron/sqlx"


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

func (r *VehicleRepository) GetAllVehicle() ([]Vehicle, error) {
  return nil, nil
}
