package main

import ("github.com/gofiber/fiber/v2"
		"api-golang-docker/database"
)

func setupRoutes(app *fibre.App){
	app.Get("/", handlers.Home)
	

}