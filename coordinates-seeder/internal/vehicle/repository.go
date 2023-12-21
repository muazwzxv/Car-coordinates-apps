package vehicle

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	TableVehicles = "vehicle"
)

type vehicle struct {
	name        string `db:"name"`
	vehicleType string `db:"type"`
	brand       string `db:"brand"`
	buildDate   string `db:"build_date"`

	lastLatitude  float64 `db:"last_latitude"`
	lastLongitude float64 `db:"last_longitude"`
}

type IVehicleRepository interface {
	GetAllVehicle(context.Context) ([]*VehicleDomain, error)
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

func (r *VehicleRepository) GetAllVehicle(ctx context.Context) ([]*VehicleDomain, error) {
	query := `
    SELECT 
      name, 
      type, 
      brand, 
      build_date, 
      last_latitude, 
      last_longitude
    FROM %s
    LIMIT 10`

	vehicles := make([]*vehicle, 0)
	rows, err := r.db.QueryxContext(ctx, fmt.Sprintf(query, TableVehicles))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		vehicle := &vehicle{}
		if err := rows.StructScan(vehicle); err != nil {
			return nil, errors.New("scanning error")
		}
    vehicles = append(vehicles, vehicle)
	}

	return convertToDomainVehicle(vehicles), nil
}

func convertToDomainVehicle(vehicles []*vehicle) []*VehicleDomain {
  domainVehicles := make([]*VehicleDomain, 0)

  for _, vehicle := range vehicles {
    domainVehicles = append(domainVehicles, &VehicleDomain{
      Name: vehicle.name,
      Type: vehicle.vehicleType,
      Brand: vehicle.brand,
      BuildDate: vehicle.buildDate,
      LastLatitude: vehicle.lastLatitude,
      LastLongitude: vehicle.lastLongitude,
    })
  }

  return domainVehicles
}
