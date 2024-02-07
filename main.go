package main

import (
	"belajar-gin/database"
	"belajar-gin/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
