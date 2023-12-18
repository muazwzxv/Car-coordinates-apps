package truck

import "github.com/jmoiron/sqlx"


type ITruckRepository interface {

}

type TruckRepository struct {
	db *sqlx.DB
}

var _ ITruckRepository = (*TruckRepository)(nil)

func NewTruckRepository(db *sqlx.DB) *TruckRepository {
	return &TruckRepository{
		db: db,
	}
}

func (r *TruckRepository) GetAllTruck() {}
