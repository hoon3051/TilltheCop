package main

import (
	"os"
	"log"

	"github.com/hoon3051/TilltheCop/initializer"
	"github.com/hoon3051/TilltheCop/router"
)

func init(){
	initializer.LoadEnv()
	initializer.InitDB()
}

func main() {
	router := router.SetupRouter()

	port := os.Getenv("PORT")
	log.Println("Server is running on port " + port)
	router.Run(":" + port) 
}