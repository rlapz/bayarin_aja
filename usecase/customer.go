package usecase

import (
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
	"github.com/rlapz/bayarin_aja/utils"
)

type customer struct {
	repoCustomer repo.CustomerRepo
	usecaseToken TokenUsecase
}

func NewCustomerUsecase(c repo.CustomerRepo, t TokenUsecase) CustomerUsecase {
	return &customer{
		c,
		t,
	}
}

func (self *customer) Login(cust *model.Customer, secret *config.Secret) (string, error) {
	//TODO: verify username and password
	res, err := self.repoCustomer.SelectByUsername(cust.Username)
	if err != nil {
		return "", err
	}

	token, err := utils.TokenGenerate(secret.Key, res.Id, secret.ExpiresIn)
	if err != nil {
		return "", err
	}

	tok := model.Token{
		CustomerId:  res.Id,
		TokenString: token,
	}

	err = self.usecaseToken.AddOne(&tok)
	if err != nil {
		return "", err
	}

	err = self.addActivity(res.Id, "login")
	if err != nil {
		return "", err
	}

	return token, err
}

func (self *customer) Logout(id int64, tokenId int64) error {
	res, err := self.repoCustomer.SelectById(id)
	if err != nil {
		return err
	}

	err = self.usecaseToken.RemoveOne(tokenId)
	if err != nil {
		return err
	}

	return self.addActivity(res.Id, "logout")
}

func (self *customer) GetByUsername(username string) (*model.Customer, error) {
	return &model.Customer{}, nil
}

func (self *customer) GetActivities(id int64) ([]model.CustomerActivity, error) {
	return []model.CustomerActivity{}, nil
}

func (self *customer) addActivity(id int64, desc string) error {
	return nil
}
