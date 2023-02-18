package main

import (
	"log"
	configure "site-backend/config"
	"site-backend/database"
	"site-backend/routes"
)

func main() {
	database.ConnectDB()
	defer database.CloseMongo()
	api := configure.ConfigureApi()
	routes.HandleRequests(api)
	err := api.Listen(":8080")
	if err != nil {
		log.Print(err.Error())
	}
}
