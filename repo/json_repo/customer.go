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

func (self *customer) SelectById(id int64) (*model.Customer, error) {
	return &model.Customer{}, nil
}

func (self *customer) SelectByUsernameAndPassword(username,
	password string) (*model.Customer, error) {
	return &model.Customer{}, nil
}

func (self *customer) SelectActivities(customerId int64) ([]model.CustomerActivity, error) {
	return []model.CustomerActivity{}, nil
}

func (self *customer) InsertOneActivity(act *model.CustomerActivity) error {
	return nil
}
