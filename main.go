package main

import (
	"log"
	"twiteer/config"
	"twiteer/utils"
)

func main() {
	utils.LoadEnv()
	db := config.ConnectPostgres()
	err := utils.AutoMigration(db)
	if err != nil {
		log.Println(err)
	}

}
