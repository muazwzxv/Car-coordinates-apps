package vehicle

import "github.com/jmoiron/sqlx"

type TruckApp struct {
	topic      string
	repository IVehicleRepository
}

func NewVehicleApp(topicName string, db *sqlx.DB) *TruckApp {
	return &TruckApp{
		topic:      topicName,
		repository: NewVehicleRepository(db),
	}
}

func (t *TruckApp) GetAllTruck() {}

func (t *TruckApp) Publish() {}
