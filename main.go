package main

import (
	"auth/routes"
	"log"
)

func main() {

	app := routes.New()
	log.Fatal(app.Listen(":3000"))

}
