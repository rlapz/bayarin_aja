package json_repo

import (
	"sync"

	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
)

type token struct {
	path  string
	mutex sync.Mutex
}

func NewJSONTokenRepo(path string) repo.TokenRepo {
	return &token{
		path: path,
	}
}

func (self *token) SelectByToken(token string) (*model.Token, error) {
	return nil, nil
}

func (self *token) SelectByCustomerId(customerId int64) (*model.Token, error) {
	return nil, nil
}

func (self *token) InsertOne(token *model.Token) error {
	return nil
}

func (self *token) DeleteOne(id int64) error {
	return nil
}
