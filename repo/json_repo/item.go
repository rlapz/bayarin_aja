package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type item struct {
	path string
}

func NewItemRepo(path string) repo.ItemRepo {
	return &item{
		path,
	}
}

func (self *item) SelectById(id int64) (*model.Item, error) {
	return &model.Item{}, nil
}
