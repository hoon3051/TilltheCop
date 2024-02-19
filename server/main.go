package main

import (
	"log"
	"os"

	"github.com/hoon3051/TilltheCop/server/router"
	"github.com/hoon3051/TilltheCop/server/initializer"
)

func init() {
	initializer.LoadEnv()
	initializer.InitDB()
}

func main() {
	router := router.SetupRouter()

	port := os.Getenv("PORT")
	log.Println("Server is running on port " + port)
	router.Run(":" + port)
}
