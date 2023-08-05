package main

import (
	"auth/database"
	"auth/router"
	"log"
)

func main() {
	app := router.New()
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	log.Fatal(app.Listen(":3000"))
}
