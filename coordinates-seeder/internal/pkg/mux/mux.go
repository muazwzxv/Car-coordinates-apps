package mux

import (
	"context"
	"coordinates-seeder/internal/pkg/config"
  
  "github.com/gofiber/fiber/v2/middleware/logger"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	mux *fiber.App
	cfg config.Config
}

func NewFiberServerWithConfig(cfg config.Config) *FiberServer {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ReadTimeout:   60 * time.Second,
		WriteTimeout:  60 * time.Second,
	})

	app.Use(logger.New())

	return &FiberServer{
		mux: app,
		cfg: cfg,
	}
}

func (s FiberServer) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return s.mux.Listen(fmt.Sprintf("%s%s", s.cfg.ServerAddress, port))
}

func (s FiberServer) Stop(ctx context.Context) error {
	if err := s.mux.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	return nil
}

func (s FiberServer) GetMux() *fiber.App {
	return s.mux
}
