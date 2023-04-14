package json_repo

import (
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/repo"
	"github.com/rlapz/bayarin_aja/repo/json_repo/numb_db"
)

type token struct {
	db *numb_db.NumbDB
}

func NewJSONTokenRepo(db *numb_db.NumbDB) repo.TokenRepo {
	return &token{
		db,
	}
}

func (self *token) InsertOne(token *model.Token) error {
	return nil
}

func (self *token) DeleteOne(id int64) error {
	return nil
}
