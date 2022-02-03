package repository

import "github.com/hananloser/test_course/model"

// Course interface    
type Course interface {
  // Return List Of Course
	List() []model.Course
  // Create Course
	Create(model.Course) model.Course
  // Update Course
	Update(id int)
  // Find Course By Id
	FindById(id int)
 // Delete Course
	Delete(id int)
}
