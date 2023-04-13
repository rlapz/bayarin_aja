package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type customer struct {
	path string
}

func NewJSONCustomerRepo(path string) repo.CustomerRepo {
	return &customer{
		path,
	}
}

func (self *customer) SelectByUsernme(username string) (*model.Customer, error) {
	return &model.Customer{}, nil
}
