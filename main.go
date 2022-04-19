package main

import (
	"file_uploader/db"
	"file_uploader/routes"
	"file_uploader/utils"
)

func main() {

	utils.ReadSettings()

	db.StartDbConnection()

	routes.RunAllRoutes()
}
