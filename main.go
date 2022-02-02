package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/test_course/router"
)


func main(){
  app := fiber.New()

  router.SetupRoutes(app)

  log.Fatal(app.Listen(":8000"))
  
}
