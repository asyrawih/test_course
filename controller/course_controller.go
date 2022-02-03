package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/test_course/model"
	"github.com/hananloser/test_course/services"
)

func NewCourseController(services *services.CourseService) *CourseController {
	return &CourseController{service: services}
}

type CourseController struct {
	service *services.CourseService
}

func(course *CourseController) GetService(ctx *fiber.Ctx) error { 
  return ctx.JSON(model.Course{
   Name: "GOLANG PRACTICE", 
   Price: 200, 
   Author: model.User{
      Name: "Gifa",
      Age: 21,
      Address: "Makassar",
      PhoneAddress: "085123121231512",
    }, 
  })
}
func(course *CourseController) GetCourse(ctx *fiber.Ctx) error { 
  return ctx.JSON(model.Course{
   Name: "TUTORIAL NODE JS", 
   Price: 200, 
   Author: model.User{
      Name: "Aldy",
      Age: 21,
      Address: "Makassar",
      PhoneAddress: "085123121231512",
    }, 
  })
}

// Create User   
func (course *CourseController) Create(ctx *fiber.Ctx) error {
  return ctx.JSON(model.WebResponse{
    Code: 200,
    Status: "Success Create User",
    Data: model.User{
      Name: "Hanan Asyrawi Rivai",
      Age: 21,
      Address: "Tomoni", 
      PhoneAddress: "085123121231512", 
    },
  })
}

// Login method    
func (course *CourseController) Login(ctx *fiber.Ctx) error {
	return ctx.JSON(model.Course{
		Name:   "Hanan",
		Author: model.User{Name: "Hanan"},
		Price:  200,
	})
}
