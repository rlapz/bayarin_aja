package config

import (
	"log"
	"strconv"
	"time"

	"github.com/rlapz/bayarin_aja/utils"
)

type Api struct {
	HostPort string
}

type Secret struct {
	Key       []byte
	ExpiresIn time.Duration
}

type Db struct {
	HostPort string
	Name     string
	Username string
	Password string
	SslMode  string
}

type App struct {
	Api        Api
	Secret     Secret
	Db         Db
	DbJSONPath string
}

func NewAppConfig() *App {
	exp, err := strconv.Atoi(utils.GetEnvFrom("TOKEN_EXPIRES_IN"))
	if err != nil {
		// use default expiration time
		exp = int(time.Hour) * 24

		log.Println(
			"invalid `TOKEN_EXPIRES_IN` value, use default value:",
			exp,
		)
	}

	var ret = App{
		Api: Api{
			HostPort: utils.GetEnvFrom("HTTP_SERVER_HOST_PORT"),
		},
		Secret: Secret{
			Key:       []byte(utils.GetEnvFrom("SECRET_KEY")),
			ExpiresIn: time.Duration(exp),
		},
		Db: Db{
			HostPort: utils.GetEnvFrom("DB_HOST_PORT"),
			Name:     utils.GetEnvFrom("DB_NAME"),
			Username: utils.GetEnvFrom("DB_USERNAME"),
			Password: utils.GetEnvFrom("DB_PASSWORD"),
			SslMode:  utils.GetEnvFrom("DB_SSL_MODE"),
		},
	}

	return &ret
}
