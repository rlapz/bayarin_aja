package repo

import (
	"github.com/rlapz/bayarin_aja/model"
)

type CustomerRepo interface {
	SelectById(id int64) (*model.Customer, error)
	SelectByUsernameAndPassword(uname, passw string) (*model.Customer, error)
}

type CustomerActivityRepo interface {
	SelectAllByCustomerId(customerId int64) ([]model.CustomerActivity, error)
	InsertOne(act *model.CustomerActivity) error
}

type TokenRepo interface {
	SelectByToken(token string) (*model.Token, error)
	InsertOne(token *model.Token) error
	DeleteOne(id int64) error
}

type PaymentRepo interface {
	SelectAllByCustomerId(customerId int64) ([]model.Payment, error)
	SelectById(id int64) (*model.Payment, error)

	// insert one record and return commited record id
	InsertOne(payment *model.Payment) (int64, error)
}
