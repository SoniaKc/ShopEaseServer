package models

type AddPaiementRequest struct {
	Login            string `json:"login" binding:"required"`
	NomCarte         string `json:"nom_carte" binding:"required"`
	NomPersonneCarte string `json:"nom_personne_carte" binding:"required"`
	CVC              int    `json:"cvc" binding:"required"`
	DateExpiration   string `json:"date_expiration" binding:"required"`
}
