package truck

import "github.com/jmoiron/sqlx"

type TruckApp struct {
	topic string
	db    *sqlx.DB
}

func New(topicName string, db *sqlx.DB) *TruckApp {
	return &TruckApp{
		topic: topicName,
		db:    db,
	}
}

func (t *TruckApp) Publish() {
}
