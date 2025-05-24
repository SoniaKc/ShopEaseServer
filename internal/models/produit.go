package models

type AddProduitRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	Nom           string `json:"nom" binding:"required"`
	Categories    string `json:"categories" binding:"required"`
	Reduction     string `json:"reduction" binding:"required"`
	Prix          string `json:"prix" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type GetProduitRequest struct {
	Id int `json:"id" binding:"required"`
}

type Produit struct {
	Id            int    `json:"id"`
	LoginBoutique string `json:"login_boutique"`
	Nom           string `json:"nom"`
	Categories    string `json:"categories"`
	Reduction     string `json:"reduction"`
	Prix          string `json:"prix"`
	Description   string `json:"description"`
}

type UpdateProduitRequest struct {
	Id            int    `json:"id" binding:"required"`
	LoginBoutique string `json:"login_boutique,omitempty"`
	Nom           string `json:"nom,omitempty"`
	Categories    string `json:"categories,omitempty"`
	Reduction     string `json:"reduction,omitempty"`
	Prix          string `json:"prix,omitempty"`
	Description   string `json:"description,omitempty"`
}
