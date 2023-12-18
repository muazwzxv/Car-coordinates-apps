package main

import "github.com/gofiber/fiber/v2"


func main() {

  // todo: load config 
  
  // create mux
  app := fiber.New(fiber.Config{
    Prefork: true,
    StrictRouting: true,
  })

  // todo: create db connection

}
