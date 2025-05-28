package models

type AddCommentaireRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Note          string `json:"note" binding:"required"`
	Commentaire   string `json:"commentaire"`
}

type UpdateCommentaireRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	NomProduit    string `json:"nom_produit" binding:"required"`
	IdClient      string `json:"idClient" binding:"required"`
	Note          string `json:"note,omitempty"`
	Commentaire   string `json:"commentaire,omitempty"`
}
