package main

import (
	"assignment-2/database"
	"assignment-2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
