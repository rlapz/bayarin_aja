package usecase

import (
	"errors"
	"time"

	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
	"github.com/rlapz/bayarin_aja/utils"
)

type customer struct {
	repoCustomer    repo.CustomerRepo
	usecaseActivity CustomerActivityUsecase
	usecaseToken    TokenUsecase
}

func NewCustomerUsecase(c repo.CustomerRepo, ca CustomerActivityUsecase,
	t TokenUsecase) CustomerUsecase {
	return &customer{
		c,
		ca,
		t,
	}
}

func (self *customer) Login(cust *model.Customer, secret *config.Secret) (utils.Token, error) {
	res, err := self.repoCustomer.SelectByUsernameAndPassword(
		cust.Username,
		cust.Password,
	)
	if err != nil {
		if errors.Is(err, my_errors.ErrNoData) {
			return utils.Token{}, my_errors.ErrUnauthorize
		}

		return utils.Token{}, err
	}

	ret, err := utils.TokenGenerate(secret.Key, res.Id, secret.ExpiresIn)
	if err != nil {
		return ret, err
	}

	tok := model.Token{
		CustomerId:  res.Id,
		TokenString: ret.TokenString,
	}

	err = self.usecaseToken.AddOne(&tok)
	if err != nil {
		return ret, err
	}

	err = self.usecaseActivity.AddOne(&model.CustomerActivity{
		CustomerId:  cust.Id,
		Description: "login",
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return ret, err
	}

	return ret, err
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

	return self.usecaseActivity.AddOne(&model.CustomerActivity{
		CustomerId:  res.Id,
		Description: "logout",
		CreatedAt:   time.Now(),
	})
}
