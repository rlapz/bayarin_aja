package usecase

import (
	"github.com/rlapz/bayarin_aja/model"
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

func (self *token) GetIdByToken(token string) (int64, error) {
	return 0, nil
}

func (self *token) GetTokenById(id int64) (string, error) {
	return "", nil
}

func (self *token) Verify(token string, customerId int64) (int64, error) {
	return 0, nil
}

func (self *token) AddOne(token *model.Token) error {
	return nil
}

func (self *token) RemoveOne(id int64) error {
	return nil
}
