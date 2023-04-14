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
	//TODO: verify username and password
	res, err := self.repoCustomer.SelectByUsernme(usename)
	if err != nil {
		return -1, err
	}

	err = self.addActivity(res.Id, "login")
	if err != nil {
		return -1, err
	}

	return res.Id, err
}

func (self *customer) Logout(id int64, tokenId int64) error {
	res, err := self.repoCustomer.SelectById(id)
	if err != nil {
		return err
	}

	err = self.repoToken.DeleteOne(tokenId)
	if err != nil {
		return err
	}

	return self.addActivity(res.Id, "logout")
}

func (self *customer) GetActivities(id int64) ([]model.CustomerActivity, error) {
	return []model.CustomerActivity{}, nil
}

func (self *customer) addActivity(id int64, desc string) error {
	return nil
}
