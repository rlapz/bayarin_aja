package repo

import (
	"github.com/rlapz/bayarin_aja/model"
)

type CustomerRepo interface {
	SelectById(id int64) (*model.Customer, error)
	SelectByUsernameAndPassword(username, password string) (*model.Customer, error)
	SelectActivities(customerId int64) ([]model.CustomerActivity, error)
	InsertOneActivity(act *model.CustomerActivity) error
}

type TokenRepo interface {
	SelectByToken(token string) (*model.Token, error)
	InsertOne(token *model.Token) error
	DeleteOne(id int64) error
}

type PaymentRepo interface {
	SelectAll() ([]model.Payment, error)
	InsertOne(payment *model.Payment) error
}

type MerchantRepo interface {
	SelectById(id int64) (*model.Merchant, error)
}

type ItemRepo interface {
	SelectById(id int64) (*model.Item, error)
}
