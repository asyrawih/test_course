package repository

import "github.com/hananloser/test_course/model"



func NewCourseRepository() Course{
  return &CourseRepositoryImpl{}
}

type CourseRepositoryImpl struct {}

func (courserepositoryimpl *CourseRepositoryImpl) List() []model.User {
	panic("not implemented") // TODO: Implement
}

func (courserepositoryimpl *CourseRepositoryImpl) Create(_ model.User) {
	panic("not implemented") // TODO: Implement
}

func (courserepositoryimpl *CourseRepositoryImpl) Update(id int) {
	panic("not implemented") // TODO: Implement
}

func (courserepositoryimpl *CourseRepositoryImpl) FindById(id int) {
	panic("not implemented") // TODO: Implement
}

func (courserepositoryimpl *CourseRepositoryImpl) Delete(id int) {
	panic("not implemented") // TODO: Implement
}


