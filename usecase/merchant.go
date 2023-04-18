package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type merchant struct {
	repoMerch repo.MerchantRepo
}

func NewMerchantUsecase(m repo.MerchantRepo) MerchantUsecase {
	return &merchant{
		m,
	}
}

func (self *merchant) GetById(id int64) (*model.Merchant, error) {
	return self.repoMerch.SelectById(id)
}

func (self *merchant) GetByCode(code string) (*model.Merchant, error) {
	return self.repoMerch.SelectByCode(code)
}
