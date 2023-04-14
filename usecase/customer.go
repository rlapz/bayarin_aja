package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type customer struct {
	repoCustomer repo.CustomerRepo
	repoToken    repo.TokenRepo
}

func NewCustomerUsecase(c repo.CustomerRepo, t repo.TokenRepo) CustomerUsecase {
	return &customer{
		c,
		t,
	}
}

func (self *customer) Login(usename, password string) (int64, error) {
	res, err := self.repoCustomer.SelectByUsernme(usename)
	return res.Id, err
}

func (self *customer) Logout(id int64, tokenId int64) error {
	return nil
}

func (self *customer) GetActivities(id int64) ([]model.CustomerActivity, error) {
	return []model.CustomerActivity{}, nil
}
