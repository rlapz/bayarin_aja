package config

import "time"

type Api struct {
	Host string
	Port string
}

type Secret struct {
	Key       []byte
	ExpiresIn time.Duration
}

type App struct {
	Api    Api
	Secret Secret
}

func NewAppConfig() *Api {
	return nil
}
