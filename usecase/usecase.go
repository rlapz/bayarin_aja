package usecase

import (
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/utils"
)

type CustomerUsecase interface {
	// the password must be hashed
	// return generated token and expires_in
	Login(cust *model.Customer, conf *config.Secret) (utils.Token, error)

	// this action will invalidate or delete the whitelisted token
	Logout(id int64, tokenId int64) error

	GetActivities(id int64) ([]model.CustomerActivity, error)
}

type TokenUsecase interface {
	Verify(token string, customerId int64) (int64, error)
	AddOne(token *model.Token) error
	RemoveOne(id int64) error
}

type PaymentUsecase interface {
	Pay(payment *model.Payment) error
	GetAll() ([]model.Payment, error)
}

type MerchantUsecase interface {
	GetById(id int64) (*model.Merchant, error)
}

type ItemUsecase interface {
	GetById(id int64) (*model.Item, error)
}
