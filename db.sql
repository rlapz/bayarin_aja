-- table m_customer
CREATE TABLE m_customer(
	id		BIGSERIAL PRIMARY KEY,
	username	VARCHAR(64) UNIQUE NOT NULL,
	password	VARCHAR(256) NOT NULL,
	first_name	VARCHAR(64) NOT NULL,
	sure_name	VARCHAR(64)
);

-- table t_token
CREATE TABLE t_token(
	id		BIGSERIAL PRIMARY KEY,
	customer_id	BIGINT REFERENCES m_customer(id) NOT NULL,
	token		TEXT NOT NULL
);

-- table m_merchant
CREATE TABLE m_merchant(
	id		BIGSERIAL PRIMARY KEY,
	code		CHAR(8) UNIQUE NOT NULL
);

-- table t_customer_activity
CREATE TABLE t_customer_activity(
	id		BIGSERIAL PRIMARY KEY,
	customer_id	BIGINT REFERENCES m_customer(id) NOT NULL,
	description	VARCHAR(32) NOT NULL,
	created_at	TIMESTAMP NOT NULL
);

-- table t_payment
CREATE TABLE t_payment(
	id			BIGSERIAL PRIMARY KEY,
	customer_id		BIGINT REFERENCES m_customer(id) NOT NULL,
	merchant_id		BIGINT REFERENCES m_merchant(id) NOT NULL,
	amount			BIGINT NOT NULL,
	order_number		VARCHAR(128) NOT NULL,
	order_description	VARCHAR(256),
	created_at		TIMESTAMP NOT NULL
);
