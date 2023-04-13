package config

import (
	"log"
	"strconv"
	"time"

	"github.com/rlapz/bayarin_aja/utils"
)

type Api struct {
	Host string
	Port string
}

type Secret struct {
	Key       []byte
	ExpiresIn time.Duration
}

type App struct {
	Api        Api
	Secret     Secret
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
			Host: utils.GetEnvFrom("HTTP_SERVER_HOST"),
			Port: utils.GetEnvFrom("HTTP_SERVER_PORT"),
		},
		Secret: Secret{
			Key:       []byte(utils.GetEnvFrom("SECRET_KEY")),
			ExpiresIn: time.Duration(exp),
		},
		DbJSONPath: utils.GetEnvFrom("DB_JSON_PATH"),
	}

	return &ret
}
