package usecase

import "github.com/rlapz/bayarin_aja/model"

type CustomerUsecase interface {
	// the password must be hashed
	// return: customer id
	Login(usename, password string) (int64, error)

	// this action will invalidate or delete the whitelisted token
	Logout(id int64, tokenId int64) error
}

type TokenUsecase interface {
	GetIdByToken(token string) (int64, error)
	GetTokenById(id int64) (string, error)
	Verify(id int64, customerId int64) error
	AddOne(token *model.Token) error
	RemoveOne(id int64) error
}

type PaymentUsecase interface {
	Pay(payment *model.Payment, tokenId int64) error
}
