package app

import (
	"context"
	"coordinates-seeder/internal/pkg/config"
	"coordinates-seeder/internal/pkg/db"
	"coordinates-seeder/internal/pkg/kafka"
	"coordinates-seeder/internal/pkg/mux"
	"coordinates-seeder/internal/vehicle"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	server *mux.FiberServer
	db     *sqlx.DB
	pub    *kafka.KafkaPublisher
	cfg    *config.Config
}

func Setup() *Application {
	// setup config
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}

	// setup mux
	svr := mux.NewFiberServerWithConfig(*cfg)

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

	app := &Application{
		server: svr,
		db:     db,
		pub:    pub,
	}

  app.setupRoutes()

	return app
}

func (a *Application) getConfig() *config.Config {
	return a.cfg
}

func (a *Application) getServer() *mux.FiberServer {
	return a.server
}

func (a *Application) getDB() *sqlx.DB {
	return a.db
}

func (a *Application) getPublisher() *kafka.KafkaPublisher {
	return a.pub
}

func (a *Application) setupRoutes() {
	server := a.getServer()
  cfg := a.getConfig()

	mux := server.GetMux()

	// setup application module
  vehicleApp := vehicle.NewVehicleApp(cfg.Topic, a.db, a.getPublisher().Publisher)

  v1 := mux.Group("/api/v1")
  {
    v1.Post("/vehicle", vehicleApp.RegisterVehicle)
  }
}

func (a *Application) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := a.getServer().Start(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// 5 second teardown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Stop(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
