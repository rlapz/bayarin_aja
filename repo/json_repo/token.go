package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type token struct {
	path string
}

func NewJSONTokenRepo(path string) repo.TokenRepo {
	return &token{
		path,
	}
}

func (self *token) SelectIdByToken(token string) (int64, error) {
	return 0, nil
}

func (self *token) SelectTokenById(id int64) (string, error) {
	return "", nil
}

func (self *token) InsertOne(token *model.Token) error {
	return nil
}

func (self *token) DeleteOne(id int64) error {
	return nil
}
