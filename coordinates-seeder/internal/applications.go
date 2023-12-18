package main

import (
	"coordinates-seeder/internal/pkg/config"
	"coordinates-seeder/internal/pkg/db"
	"coordinates-seeder/internal/pkg/kafka"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	mux *fiber.App
	db  *sqlx.DB
	pub *kafka.KafkaPublisher
}

func setup() *Application {
	// setup config
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	// setup mux
	mux := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
	})

	// setup db
	db, err := db.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// setup publisher
	pub, err := kafka.NewAppPublisher(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &Application{
		mux: mux,
		db:  db,
		pub: pub,
	}
}

func (a *Application) GetMux() *fiber.App {
	return a.mux
}

func (a *Application) GetDB() *sqlx.DB {
	return a.db
}

func (a *Application) GetPublisher() *kafka.KafkaPublisher {
	return a.pub
}

func (a *Application) Run() {}
