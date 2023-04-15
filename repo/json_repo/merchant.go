package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type merchant struct {
	path string
}

func NewMerchantRepo(path string) repo.MerchantRepo {
	return &merchant{
		path,
	}
}

func (self *merchant) SelectById(id int64) (*model.Merchant, error) {
	return &model.Merchant{}, nil
}
