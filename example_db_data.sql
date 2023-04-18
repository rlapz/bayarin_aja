-- m_customer
-- rezki: plain password: 123
INSERT INTO m_customer(username, password, first_name, sure_name)
VALUES( 'rezki'
	,'$2a$12$HEQe5GKPA2EMC95qhbPbgeNk25PyDwu8zSJ34SGGkNtkBg0uKNP1a'
	,'Rezki'
	,'aaaaa'
);

-- doni: plain password: aaaa
INSERT INTO m_customer(username, password, first_name, sure_name)
VALUES( 'doni'
	,'$2a$12$QzT7GtR2lRhU1RDeYU1MSORLxadAycIBpKt7SDlQTPGyQ9ys2Pf2u'
	,'Doni'
	,'bbbbb'
);


-- m_merchant
INSERT INTO m_merchant(code)
VALUES('M0000001');

INSERT INTO m_merchant(code)
VALUES('M0000002');

INSERT INTO m_merchant(code)
VALUES('M0000003');
