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

		api.POST("/client/add", handlers.AddClient)
		api.GET("/client/get", handlers.GetClient)
		api.DELETE("/client/delete", handlers.DeleteClient)
		api.PUT("/client/update", handlers.UpdateClient)

		api.POST("/boutique/add", handlers.AddBoutique)
		api.GET("/boutique/get", handlers.GetBoutique)
		api.DELETE("/boutique/delete", handlers.DeleteBoutique)
		api.PUT("/boutique/update", handlers.UpdateBoutique)

		api.POST("/parametre/add", handlers.AddParametre)
		api.GET("/parametre/get", handlers.GetParametre)
		api.DELETE("/parametre/delete", handlers.DeleteParametre)
		api.PUT("/parametre/update", handlers.UpdateParametre)

		api.POST("/paiement/add", handlers.AddPaiement)
		api.GET("/paiement/get", handlers.GetPaiement)
		api.GET("/paiement/getAll", handlers.GetAllPaiement)
		api.DELETE("/paiement/delete", handlers.DeletePaiement)
		api.PUT("/paiement/update", handlers.UpdatePaiement)

		api.POST("/adresse/add", handlers.AddAdresse)
		api.GET("/adresse/get", handlers.GetAdresse)
		api.GET("/adresse/getAll", handlers.GetAllAdresse)
		api.DELETE("/adresse/delete", handlers.DeleteAdresse)
		api.PUT("/adresse/update", handlers.UpdateAdresse)

		api.POST("/produit/add", handlers.AddProduit)
		api.GET("/produit/get", handlers.GetProduit)
		api.GET("/produit/getProduitsRecherche", handlers.GetProduitsRecherche)
		api.GET("/produit/getAllProduits", handlers.GetAllProduits)
		api.GET("/produit/getAllByBoutique", handlers.GetAllProduitByBoutique)
		api.GET("/produit/getPopulaires", handlers.GetPopulaires)
		api.DELETE("/produit/delete", handlers.DeleteProduit)
		api.PUT("/produit/update", handlers.UpdateProduit)

		api.POST("/panier/add", handlers.AddPanier)
		api.GET("/panier/getQte", handlers.GetQteInPanier)
		api.GET("/panier/getAll", handlers.GetFullPanier)
		api.DELETE("/panier/delete", handlers.DeletePanier)
		api.PUT("/panier/update", handlers.UpdateQteInPanier)

		api.POST("/favoris/add", handlers.AddFavori)
		api.GET("/favoris/getAll", handlers.GetAllFavoris)
		api.DELETE("/favoris/delete", handlers.DeleteFavoris)

		api.POST("/vente/add", handlers.AddVente)
		api.GET("/vente/getByIdTransaction", handlers.GetAllTransaction)
		api.GET("/vente/getByClient", handlers.GetAllVentesClient)
		api.GET("/vente/getByBoutique", handlers.GetAllVentesBoutique)
		api.DELETE("/vente/deleteByIdTransaction", handlers.DeleteAllTransaction)
		api.PUT("/vente/updateStatut", handlers.UpdateTransactionStatut)

		api.POST("/commentaire/add", handlers.AddCommentaire)
		api.GET("/commentaire/getByProduit", handlers.GetAllComsProduit)
		api.GET("/commentaire/getByClient", handlers.GetAllComsClient)
		api.DELETE("/commentaire/delete", handlers.DeleteCommentaire)
		api.PUT("/commentaire/update", handlers.UpdateCommentaire)
	}

	return router
}
