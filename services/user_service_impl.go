package services

import "github.com/hananloser/test_course/model"


func NewUserService() UserService {return &UserServiceImpl{}}

type UserServiceImpl struct {}

func (userserviceimpl *UserServiceImpl) Add(request *model.User) model.WebResponse {
	panic("not implemented") // TODO: Implement
}

func (userserviceimpl *UserServiceImpl) List() model.WebResponse {
	panic("not implemented") // TODO: Implement
}

func (userserviceimpl *UserServiceImpl) Delete() model.WebResponse {
	panic("not implemented") // TODO: Implement
}
