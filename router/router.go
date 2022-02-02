package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hananloser/test_course/controller"
)

// ALl Handler Will Come in here
func SetupRoutes(app *fiber.App) {
	// Welcome Routes
	api := app.Group("/api", logger.New())
	api.Get("/",controller.Welcome)

}
