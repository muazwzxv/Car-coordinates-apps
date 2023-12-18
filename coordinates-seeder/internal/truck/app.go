package truck

import "github.com/jmoiron/sqlx"

type TruckApp struct {
	topic      string
	repository ITruckRepository
}

func NewTruckApp(topicName string, db *sqlx.DB) *TruckApp {
	return &TruckApp{
		topic:      topicName,
		repository: NewTruckRepository(db),
	}
}

func (t *TruckApp) GetAllTruck() {}

func (t *TruckApp) Publish() {}
