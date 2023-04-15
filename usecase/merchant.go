package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type merchant struct {
	repoMerchant repo.MerchantRepo
}

func NewMerchantUsecase(r repo.MerchantRepo) MerchantUsecase {
	return &merchant{
		r,
	}
}

func (self *merchant) GetById(id int64) (*model.Merchant, error) {
	return self.repoMerchant.SelectById(id)
}
