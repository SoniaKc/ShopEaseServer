package models

type AddFavorisRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
}
