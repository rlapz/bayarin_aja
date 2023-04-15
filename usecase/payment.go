package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type payment struct {
	repoPayment     repo.PaymentRepo
	usecaseItem     ItemUsecase
	usecaseMerchant MerchantUsecase
}

func NewPaymentUsecase(p repo.PaymentRepo, i ItemUsecase, m MerchantUsecase) PaymentUsecase {
	return &payment{
		p,
		i,
		m,
	}
}

func (self *payment) Pay(payment *model.Payment) error {
	return self.repoPayment.InsertOne(payment)
}

func (self *payment) GetAll() ([]model.Payment, error) {
	return self.repoPayment.SelectAll()
}
