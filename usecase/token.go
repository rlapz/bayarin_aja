package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type token struct {
	repoToken repo.TokenRepo
}

func NewTokenUsecase(t repo.TokenRepo) TokenUsecase {
	return &token{
		t,
	}
}

func (self *token) Verify(token string, customerId int64) (int64, error) {
	res, err := self.repoToken.SelectByToken(token)
	if err != nil {
		return -1, err
	}

	if res.CustomerId != customerId {
		return -1, my_errors.ErrNoData
	}

	return res.Id, nil
}

func (self *token) AddOne(token *model.Token) error {
	return self.repoToken.InsertOne(token)
}

func (self *token) RemoveOne(id int64) error {
	return self.repoToken.DeleteOne(id)
}
