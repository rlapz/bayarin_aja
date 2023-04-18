package psql

import (
	"database/sql"

	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type customer struct {
	db *sql.DB
}

func NewPsqlCustomerRepo(db *sql.DB) repo.CustomerRepo {
	return &customer{
		db,
	}
}

func (self *customer) SelectById(id int64) (*model.Customer, error) {
	return self.selectOne("WHERE id = $1", id)
}

func (self *customer) SelectByUsernameAndPassword(uname, passw string) (*model.Customer, error) {
	return self.selectOne("WHERE username = $1 AND password = $2", uname, passw)
}

// private
func (self *customer) selectOne(key string, val ...any) (*model.Customer, error) {
	var q = `
		SELECT	 id
			,username
			,password
			,first_name
			,sure_name
		FROM m_customer 
	`

	ret := new(model.Customer)
	row := self.db.QueryRow(q+key, val...)
	err := row.Scan(
		&ret.Id,
		&ret.Username,
		&ret.Password,
		&ret.FirstName,
		&ret.SureName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, my_errors.ErrNoData
		}

		return nil, err
	}

	return ret, nil
}
