package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
	"github.com/rlapz/bayarin_aja/repo/json_repo/numb_db"
)

type customer struct {
	db *numb_db.NumbDB
}

func NewJSONCustomerRepo(db *numb_db.NumbDB) repo.CustomerRepo {
	return &customer{
		db,
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
