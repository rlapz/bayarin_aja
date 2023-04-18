package psql

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/repo"
)

type payment struct {
	db *sql.DB
}

func NewPsqlPaymentRepo(db *sql.DB) repo.PaymentRepo {
	return &payment{
		db,
	}
}

func (self *payment) SelectAllByCustomerId(customerId int64) ([]model.Payment, error) {
	const q = `
		SELECT	 id
			,customer_id
			,merchant_id
			,amount
			,order_number
			,order_description
			,created_at
		FROM t_payment
		WHERE customer_id = $1
	`

	rows, err := self.db.Query(q, customerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]model.Payment, 0)
	for rows.Next() {
		p := model.Payment{}
		err = rows.Scan(
			&p.Id,
			&p.CustomerId,
			&p.MerchantId,
			&p.Amount,
			&p.OrderNumber,
			&p.OrderDescription,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		ret = append(ret, p)
	}

	if len(ret) == 0 {
		return nil, my_errors.ErrNoData
	}

	return ret, nil
}

func (self *payment) SelectById(id int64) (*model.Payment, error) {
	const q = `
		SELECT	 id
			,customer_id
			,merchant_id
			,amount
			,order_number
			,order_description
			,created_at
		FROM t_payment
		WHERE id = $1
	`

	ret := new(model.Payment)
	row := self.db.QueryRow(q, id)
	err := row.Scan(
		&ret.Id,
		&ret.CustomerId,
		&ret.MerchantId,
		&ret.Amount,
		&ret.OrderNumber,
		&ret.OrderDescription,
		&ret.CreatedAt,
	)

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

// insert one record and return commited record id
func (self *payment) InsertOne(payment *model.Payment) (int64, error) {
	const q = `
		INSERT INTO t_payment(
			 customer_id
			,merchant_id
			,amount
			,order_number
			,order_description
			,created_at
		)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	ret := int64(0)
	row := self.db.QueryRow(
		q,
		payment.CustomerId,
		payment.MerchantId,
		payment.Amount,
		payment.OrderNumber,
		payment.OrderDescription,
		payment.CreatedAt,
	)
	err := row.Scan(&ret)
	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return -1, my_errors.ErrDuplicateEntry
		}

		return -1, err
	}

	if err = row.Err(); err != nil {
		return -1, err
	}

	return ret, nil
}
