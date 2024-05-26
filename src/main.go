package main

import (
	"log"
	"vocabula/common"
)

func main() {
	log.Println("Loading server configuration...")
	err := common.LoadServerConfig()
	if err != nil {
		log.Println("Error loading server configuration: " + err.Error())
		return
	}
	log.Println("Loaded server configuration")

	log.Println("Initializing database connecton...")
	err = common.InitDbConnection()
	if err != nil {
		log.Println("Error initializing database connection: " + err.Error())
		return
	}
	log.Println("Initialized database connecton")

	log.Println("Starting server...")
	err = ServerMain()
	if err != nil {
		log.Println("Fatal error running server: " + err.Error())
		return
	}

	log.Println("Deinitializing database connection...")
	common.DeinitDbConnection()
}
