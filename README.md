# `bayarin_aja`

## Disclaimer: This project is aims to fulfill the interview assignment


## How to initialize
1. You should copy the `example.env` file to `.env`.
2. Edit the `.env` file, change the `SECRET_KEY`, please make sure you don't share this key.
3. Edit other env variable if necessary
4. Import `db.sql` database scheme. (PostgreSQL)


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

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
		 }
		 ```

   - Logout

	   `POST /v1/customer/logout`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
		 {
		 }
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
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
		 }
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
		 }
		 ```

   - Show payment activities (History)

	   `GET /v1/payment/activity`
	   - Request header
		 * `Authorization`: `Bearer xxxxxxxxxxxxx`
		 * `Content-Type`: `application/json`

	   - Request body
		 ```
		 {
		 }
		 ```

	   - Responses

		 `200 OK`

		 Object:
		 ```
		 {
		 }
		 ```
