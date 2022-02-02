package repository

import "github.com/hananloser/test_course/model"

type Course interface {
	List() []model.User

	Create(model.User)

	Update(id int)

	FindById(id int)

	Delete(id int)
}
