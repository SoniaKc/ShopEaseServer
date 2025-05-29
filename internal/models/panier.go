package models

type AddPanierRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Quantite      string `json:"quantite" binding:"required"`
}

type UpdatePanierRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Quantite      string `json:"quantite"`
}
