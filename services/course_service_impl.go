package services

import "github.com/hananloser/test_course/model"


func NewCourseService() CourseService{
  return &CourseServiceImpl{}
}

type CourseServiceImpl struct {}

func (courseserviceimpl *CourseServiceImpl) List() []model.Course {
	panic("not implemented") // TODO: Implement
}

func (courseserviceimpl *CourseServiceImpl) FindById() model.Course {
	panic("not implemented") // TODO: Implement
}
