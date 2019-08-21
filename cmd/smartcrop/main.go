package main

import (
	"log"

	"github.com/titpetric/smartcrop/service"
)

func main() {
	if err := service.Start(); err != nil {
		log.Fatal(err)
	}
}
