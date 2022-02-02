package services

import (
	"github.com/hananloser/test_course/model"
)

type UserService interface {

	Add(request *model.User) model.WebResponse

	List() model.WebResponse

	Delete() model.WebResponse
}
