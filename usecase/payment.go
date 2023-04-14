package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type payment struct {
	repoPayment repo.PaymentRepo
}

func NewPaymentUsecase(p repo.PaymentRepo) PaymentUsecase {
	return &payment{
		p,
	}
}

func (self *payment) Pay(payment *model.Payment) error {
	return self.repoPayment.InsertOne(payment)
}

func (self *payment) GetAll() ([]model.Payment, error) {
	return self.repoPayment.SelectAll()
}
