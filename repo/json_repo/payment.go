package json_repo

import (
	"sync"

	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type payment struct {
	path  string
	mutex sync.Mutex
}

func NewJSONPaymentRepo(path string) repo.PaymentRepo {
	return &payment{
		path: path,
	}
}

func (self *payment) SelectAll() ([]model.Payment, error) {
	return []model.Payment{}, nil
}

func (self *payment) InsertOne(payment *model.Payment) error {
	return nil
}
