package repo

import (
	"github.com/rlapz/bayarin_aja/model"
)

type CustomerRepo interface {
	SelectByUsernme(username string) (*model.Customer, error)
	SelectActivities(customerId int64) ([]model.CustomerActivity, error)
}

type TokenRepo interface {
	SelectIdByToken(token string) (int64, error)
	SelectTokenById(id int64) (string, error)
	InsertOne(token *model.Token) error
	DeleteOne(id int64) error
}

type PaymentRepo interface {
	InsertOne(payment *model.Payment) error
}
