package models

type AddVenteRequest struct {
	IdTransaction string `json:"idTransaction" binding:"required"`
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	NomAdresse    string `json:"nom_adresse" binding:"required"`
	NomPaiement   string `json:"nom_paiement" binding:"required"`
	Quantite      string `json:"quantite" binding:"required"`
	Total         string `json:"total" binding:"required"`
	Date_vente    string `json:"date_vente" binding:"required"`
	Statut        string `json:"statut" binding:"required"`
}
