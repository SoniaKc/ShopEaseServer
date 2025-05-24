package models

type AddPaiementRequest struct {
	Login            string `json:"login" binding:"required"`
	NomCarte         string `json:"nom_carte" binding:"required"`
	NomPersonneCarte string `json:"nom_personne_carte" binding:"required"`
	Numero           string `json:"numero" binding:"required"`
	CVC              string `json:"cvc" binding:"required"`
	DateExpiration   string `json:"date_expiration" binding:"required"`
}

type GetPaiementRequest struct {
	Login    string `json:"login" binding:"required"`
	NomCarte string `json:"nom_carte" binding:"required"`
}

type Paiement struct {
	Login            string `json:"login"`
	NomCarte         string `json:"nom_carte"`
	NomPersonneCarte string `json:"nom_personne_carte"`
	Numero           string `json:"numero"`
	CVC              string `json:"cvc"`
	DateExpiration   string `json:"date_expiration"`
}

type UpdatePaiementRequest struct {
	Login            string `json:"login" binding:"required"`
	NomCarte         string `json:"nom_carte" binding:"required"`
	NomPersonneCarte string `json:"nom_personne_carte,omitempty"`
	Numero           string `json:"numero,omitempty"`
	CVC              string `json:"cvc,omitempty"`
	DateExpiration   string `json:"date_expiration,omitempty"`
}
