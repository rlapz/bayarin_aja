# `bayarin_aja`

## Disclaimer: This project is aims to fulfill the interview assignment


## How to initialize
1. You should copy the `example.env` file to `.env`.
2. Edit the `.env` file, change the `SECRET_KEY`, please make sure you don't share this key.
3. Edit other env variable if necessary
4. Create a new database, give it a name: `bayarin_aja`
5. Import `db.sql` database scheme. (PostgreSQL)
6. Import `example_db_data.sql` database data. (PostgreSQL), there are data examples inside it:

   `m_customer`

	 ```
	 username  : rezki
	 password  : $2a$12$HEQe5GKPA2EMC95qhbPbgeNk25PyDwu8zSJ34SGGkNtkBg0uKNP1a
	 first_name: Rezki
	 sure_name : aaaaa
	 ```

	 ```
	 username  : doni
	 password  : $2a$12$QzT7GtR2lRhU1RDeYU1MSORLxadAycIBpKt7SDlQTPGyQ9ys2Pf2u
	 first_name: Doni
	 sure_name : bbbbb
	 ```

	 ----

	 `m_merchant`

	 ```
	 code: M0000001
	 ```

	 ```
	 code: M0000002
	 ```

	 ```
	 code: M0000003
	 ```


## How to run
There are two options:
1. `go run .`
2. `go build` then `./bayarin_aja`


## How to stop
1. You can just hit: `<Ctrl-C>` and wait for 3 seconds


## API Documentation
The API documentation below inspired by: https://doc.gopay.com

For the `login`, `logout`, and `payment` are automatically "recorded" as `activity` (History).

**1. Customer**
   - Login

	   `POST /v1/customer/login`
	   - Request header
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
		 {
		     "username": "[username]",
		     "password": "[hashed_password]"
		 }
		 ```

		 Example:

		 ```
		{
			  "username": "rezki",
				"password": "$2a$12$HEQe5GKPA2EMC95qhbPbgeNk25PyDwu8zSJ34SGGkNtkBg0uKNP1a"
		}
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
     {
         "status": "Success",
         "message": "login success",
         "data": {
             "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE4MzgyNDMsImlkIjoxfQ.qP3OTczoI4isWe1Q3w9GVmpiX30TD9WQMzg0zmBfls4",
             "expires_in": 1681838243 // <-- Unix epoch
         }
     }
		 ```

   - Logout

	   `POST /v1/customer/logout`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
		     "status": "Success",
		     "message": "logged out"
		 }
		 ```

   - Show Activity (History)

	   `GET /v1/customer/activity`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
     {
         "status": "Success",
         "data": [
             {
                 "id": 1,
                 "customer_id": 1,
                 "description": "login",
                 "created_at": "2023-04-18T22:52:51.7917Z"
             },
             {
                 "id": 2,
                 "customer_id": 1,
                 "description": "logout",
                 "created_at": "2023-04-18T22:56:29.239617Z"
             },
             {
                 "id": 3,
                 "customer_id": 1,
                 "description": "login",
                 "created_at": "2023-04-18T23:17:23.899902Z"
             },
             {
                 "id": 4,
                 "customer_id": 1,
                 "description": "login",
                 "created_at": "2023-04-18T23:39:29.001597Z"
             },
             {
                 "id": 5,
                 "customer_id": 1,
                 "description": "logout",
                 "created_at": "2023-04-18T23:39:44.438759Z"
             },
             {
                 "id": 6,
                 "customer_id": 1,
                 "description": "login",
                 "created_at": "2023-04-18T23:39:49.156432Z"
             }
         ]
     }
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
		 }
		 ```

**2. Payment**
   - Do payment

	   `POST /v1/payment/pay`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
     {
         "customer_id": 1,
         "merchant_code": "M0000001",
         "amount": 1234,
         "order_number": "111"
     }
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
     {
         "status": "Success",
         "data": {
             "id": 2,
             "customer_id": 1,
             "merchant_code": "M0000001",
             "amount": 1234,
             "order_number": "111",
             "order_description": "",
             "created_at": "0001-01-01T00:00:00Z"
         }
     }
		 ```

   - Show payment activities (History)

	   `GET /v1/payment/activity`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
     {
         "status": "Success",
         "data": [
             {
                 "id": 1,
                 "customer_id": 1,
                 "merchant_code": "M0000001",
                 "amount": 1234,
                 "order_number": "111",
                 "order_description": "",
                 "created_at": "0001-01-01T00:00:00Z"
             },
             {
                 "id": 2,
                 "customer_id": 1,
                 "merchant_code": "M0000001",
                 "amount": 1234,
                 "order_number": "111",
                 "order_description": "",
                 "created_at": "0001-01-01T00:00:00Z"
             }
         ]
     }
		 ```
