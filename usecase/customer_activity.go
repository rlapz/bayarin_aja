package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type custAct struct {
	repoCustomerAct repo.CustomerActivityRepo
}

func NewCustomerActivityUsecase(c repo.CustomerActivityRepo) CustomerActivityUsecase {
	return &custAct{
		c,
	}
}

func (self *custAct) GetActivities(customerId int64) ([]model.CustomerActivity, error) {
	return self.repoCustomerAct.SelectAllByCustomerId(customerId)
}

func (self *custAct) AddOne(act *model.CustomerActivity) error {
	return self.repoCustomerAct.InsertOne(act)
}
