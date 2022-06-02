package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-example/database"
	"github.com/golang-example/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}

/*
------------------POINTERS-----------------
var name *string = new(string)
*name = 5
fmt.Println(name, *name)

or

name := 5
fmt.Println(name, &name)  -> & is a pointer to the variable name
*/
