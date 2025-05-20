package main

import (
    "log"
    "os"
    "shop-ease-server/internal/routes"
)

func main() {

    /*if err := storage.InitPostgres(); err != nil {
        log.Fatal("PostGre init failed:", err)
    }*/

    router := routes.SetupRouter() // Initialisation du routeur

    port := os.Getenv("PORT")
    if port == "" {
        port = "8090"
    }

    if err := router.Run(":" + port); err != nil {
        log.Panicf("error: %s", err)
    }

}