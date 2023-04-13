package main

import (
	"log"

	"github.com/rlapz/bayarin_aja/delivery"
)

func main() {
	restApi := delivery.NewRestApiDelivery()
	if err := restApi.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
