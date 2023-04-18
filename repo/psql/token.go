package psql

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type token struct {
	db *sql.DB
}

func NewPsqlTokenRepo(db *sql.DB) repo.TokenRepo {
	return &token{
		db,
	}
}

func (self *token) SelectByToken(token string) (*model.Token, error) {
	const q = `
		SELECT	 id
			,customer_id
			,token
		FROM t_token
		WHERE token = $1
	`

	ret := new(model.Token)
	row := self.db.QueryRow(q, token)
	err := row.Scan(&ret.Id, &ret.CustomerId, &ret.TokenString)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, my_errors.ErrNoData
		}

		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *token) InsertOne(token *model.Token) error {
	const q = `
		INSERT INTO t_token(
			 customer_id
			,token
		)
		VALUES($1, $2)
	`

	_, err := self.db.Exec(q, token.CustomerId, token.TokenString)
	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return my_errors.ErrDuplicateEntry
		}

		return err
	}

	return nil
}

func (self *token) DeleteOne(id int64) error {
	res, err := self.db.Exec("DELETE FROM t_token WHERE id = $1", id)
	if err != nil {
		return err
	}

	rw, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rw == 0 {
		return my_errors.ErrNoData
	}

	return nil
}
