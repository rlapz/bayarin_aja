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

func (self *payment) Pay(payment *model.Payment) (*model.Payment, error) {
	id, err := self.repoPayment.InsertOne(payment)
	if err != nil {
		return nil, err
	}

	return self.repoPayment.SelectById(id)
}

func (self *payment) GetAllByCustomerId(customerId int64) ([]model.Payment, error) {
	return self.repoPayment.SelectAllByCustomerId(customerId)
}
