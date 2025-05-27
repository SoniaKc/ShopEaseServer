package models

type AddVenteRequest struct {
	IdTransaction string `json:"idTransaction" binding:"required"`
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Quantite      string `json:"reduction" binding:"required"`
	Total         string `json:"total" binding:"required"`
	Date_vente    string `json:"date_vente" binding:"required"`
	Statut        string `json:"statut" binding:"required"`
}
