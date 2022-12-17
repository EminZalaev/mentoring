package main

import (
	"log"

	"mentoring/cmd/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("error run app", err)
	}
}
