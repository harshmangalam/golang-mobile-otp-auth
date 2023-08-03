package main

import (
	"auth/database"
	"auth/routes"
	"log"
)

func main() {

	app := routes.New()
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	log.Fatal(app.Listen(":3000"))

}
