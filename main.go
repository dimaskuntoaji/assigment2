package main

import (
	db "restapi2/model"
	"restapi2/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	db.Init()
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file: ", e)
	}
	router := router.SetupRouter()
	port := os.Getenv("SERVER_PORT")
	if len(os.Args) > 1 {
		reqPort := os.Args[1]

		if reqPort != "" {
			port = reqPort
		}
	}
	if port == "" {
		port = "8080"
	}
	type Job interface {
		Run()
	}
	router.Run(":" + port)
}