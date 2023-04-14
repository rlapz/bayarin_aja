package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
	"github.com/rlapz/bayarin_aja/repo/json_repo/numb_db"
)

type payment struct {
	db *numb_db.NumbDB
}

func NewJSONPaymentRepo(db *numb_db.NumbDB) repo.PaymentRepo {
	return &payment{
		db,
	}
}

func (self *payment) SelectAll() ([]model.Payment, error) {
	return []model.Payment{}, nil
}

func (self *payment) InsertOne(payment *model.Payment) error {
	return nil
}
