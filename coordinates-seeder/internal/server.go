package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db *sqlx.DB
  mux *fiber.App
}

func (s *Server) GetDB() *sqlx.DB {
  return s.db
}

func (s *Server) GetMux() *fiber.App {
  return s.mux
}
