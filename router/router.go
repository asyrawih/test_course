package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hananloser/test_course/controller"
	"github.com/hananloser/test_course/services"
)

// ALl Handler Will Come in here
func SetupRoutes(app *fiber.App) {

  // Depedency Injections 
  services := services.NewCourseService()
  serviceControler := controller.NewCourseController(&services)
	// Welcome Routes
	api := app.Group("/api", logger.New())
	api.Get("/",controller.Welcome)
  api.Get("/service", serviceControler.GetCourse)
  api.Get("/course" , serviceControler.GetCourse)

}
