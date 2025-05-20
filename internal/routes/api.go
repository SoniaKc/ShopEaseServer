package routes

import (
	"shop-ease-server/internal/handlers"
	"shop-ease-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	// Routes
	api := router.Group("/api")
	{
		api.HEAD("/health", handlers.HealthCheck)
		/*
					api.GET("/admin/users", middleware.AdminAuth(), handlers.GetAllUsers)
			        api.DELETE("/admin/del_table", middleware.AdminAuth(), handlers.DeleteTable)

			        api.GET("/user/forgetPswrd", handlers.ForgetPswrd)
		*/

		api.POST("/user/add", handlers.AddClient)
		api.POST("/boutique/add", handlers.AddBoutique)
		api.POST("/parametre/add", handlers.AddParametre)
		api.POST("/paiement/add", handlers.AddPaiement)
		api.POST("/adresse/add", handlers.AddAdresse)

		/*
		   api.DELETE("/user/delete", handlers.DeleteUser)

		   api.POST("/user/addPlantCollection", handlers.AddPlantCollection)
		   api.GET("/user/getPlantCollection", handlers.GetPlantCollection)
		   api.GET("/user/getAllPlantCollections", handlers.GetPlantCollections)
		   api.DELETE("/user/deletePlantCollection", handlers.DeletePlantCollection)
		   api.DELETE("/user/deletePlantCollections", handlers.DeletePlantCollections)

		   api.POST("/user/addPlant", handlers.AddPlant)
		   api.GET("/user/getPlant", handlers.GetPlant)
		   api.GET("/user/getAllPlants", handlers.GetPlants)
		   api.DELETE("/user/deletePlant", handlers.DeletePlant)
		   api.DELETE("/user/deletePlants", handlers.DeletePlants)*/

	}

	return router
}
