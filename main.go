package main

import (
	"log"
	"os"
	"shop-ease-server/internal/routes"
	"shop-ease-server/internal/storage"
)

func main() {

	if err := storage.InitPostgres(); err != nil {
		log.Fatal("PostGre init failed:", err)
	}

	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
