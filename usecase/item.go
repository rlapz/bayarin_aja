package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type item struct {
	repoItem repo.ItemRepo
}

func NewItemUsecase(r repo.ItemRepo) ItemUsecase {
	return &item{
		r,
	}
}

func (self *item) GetById(id int64) (*model.Item, error) {
	return self.repoItem.SelectById(id)
}
