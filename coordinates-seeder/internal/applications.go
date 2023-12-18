package app

import (
	"context"
	"coordinates-seeder/internal/pkg/config"
	"coordinates-seeder/internal/pkg/db"
	"coordinates-seeder/internal/pkg/kafka"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	mux *fiber.App
	db  *sqlx.DB
	pub *kafka.KafkaPublisher
}

func Setup() *Application {
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

func (a *Application) Run() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := a.GetMux().Start(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Stop(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
