package psql

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type custAct struct {
	db *sql.DB
}

func NewPsqlCustomerActivityRepo(db *sql.DB) repo.CustomerActivityRepo {
	return &custAct{
		db,
	}
}

func (self *custAct) SelectAllByCustomerId(customerId int64) ([]model.CustomerActivity, error) {
	const q = `
		SELECT	 id
			,customer_id
			,description
			,created_at
		FROM t_customer_activity
		WHERE customer_id = $1
	`

	rows, err := self.db.Query(q, customerId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ret := make([]model.CustomerActivity, 0)
	for rows.Next() {
		c := model.CustomerActivity{}
		err := rows.Scan(&c.Id, &c.CustomerId, &c.Description, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		ret = append(ret, c)
	}

	if len(ret) == 0 {
		return nil, my_errors.ErrNoData
	}

	return ret, nil
}

func (self *custAct) InsertOne(act *model.CustomerActivity) error {
	const q = `
		INSERT INTO t_customer_activity(
			 customer_id
			,description
			,created_at
		)
		VALUES($1, $2, $3)
	`

	_, err := self.db.Exec(q, act.CustomerId, act.Description, act.CreatedAt)
	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return my_errors.ErrDuplicateEntry
		}

		return err
	}

	return nil
}
