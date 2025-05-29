package models

type AddPanierRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Quantite      string `json:"quantite" binding:"required"`
}

type Panier struct {
	LoginBoutique string `json:"login_boutique"`
	NomProduit    string `json:"nom_produit"`
	IdClient      string `json:"idClient"`
	Quantite      string `json:"quantite"`
}

type UpdatePanierRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Quantite      string `json:"quantite"`
}
