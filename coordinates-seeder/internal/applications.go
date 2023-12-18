package main

import (
	"coordinates-seeder/internal/pkg/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// todo: load config
	_, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	// create mux
	_ = fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
	})

	// todo: create db connection
}
