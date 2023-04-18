package psql

import (
	"database/sql"

	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type merchant struct {
	db *sql.DB
}

func NewMerchantRepo(db *sql.DB) repo.MerchantRepo {
	return &merchant{
		db,
	}
}

func (self *merchant) SelectById(id int64) (*model.Merchant, error) {
	return self.selectOne("id", id)
}

func (self *merchant) SelectByCode(code string) (*model.Merchant, error) {
	return self.selectOne("code", code)
}

func (self *merchant) selectOne(key string, val any) (*model.Merchant, error) {
	var q = `
		SELECT	 id
			,code
		FROM m_merchant
		WHERE 
	`

	q += (key + " = $1")

	ret := new(model.Merchant)
	row := self.db.QueryRow(q, val)
	err := row.Scan(&ret.Id, &ret.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, my_errors.ErrNoData
		}

		return nil, err
	}

	return ret, nil
}
