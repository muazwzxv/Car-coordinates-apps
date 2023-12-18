package vehicle

import "github.com/jmoiron/sqlx"


type IVehicleRepository interface {

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

func (r *VehicleRepository) GetAllTruck() {}
