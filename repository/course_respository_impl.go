package repository

import "github.com/hananloser/test_course/model"


func NewCourseRepository() Course{
  return &CourseRepositoryImpl{}
}

type CourseRepositoryImpl struct {}
// Return List Of Course
func (courserepositoryimpl *CourseRepositoryImpl) List() []model.Course {
	panic("not implemented") // TODO: Implement
}

// Create Course
// @Todo this must pointer model.USer
func (courserepositoryimpl *CourseRepositoryImpl) Create(_ model.Course) model.Course {
	panic("not implemented") // TODO: Implement
}

// Update Course
func (courserepositoryimpl *CourseRepositoryImpl) Update(id int) {
	panic("not implemented") // TODO: Implement
}

// Find Course By Id
func (courserepositoryimpl *CourseRepositoryImpl) FindById(id int) {
	panic("not implemented") // TODO: Implement
}

// Delete Course
func (courserepositoryimpl *CourseRepositoryImpl) Delete(id int) {
	panic("not implemented") // TODO: Implement
}


