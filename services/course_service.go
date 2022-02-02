package services

import "github.com/hananloser/test_course/model"


type CourseService interface {
  
  List() []model.Course

  FindById() model.Course

}


