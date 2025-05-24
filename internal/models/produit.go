package models

type AddProduitRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	Nom           string `json:"nom" binding:"required"`
	Categories    string `json:"categories" binding:"required"`
	Reduction     string `json:"reduction"`
	Prix          string `json:"prix" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type GetProduitRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	Nom           string `json:"nom" binding:"required"`
}

type Produit struct {
	LoginBoutique string `json:"login_boutique"`
	Nom           string `json:"nom"`
	Categories    string `json:"categories"`
	Reduction     string `json:"reduction"`
	Prix          string `json:"prix"`
	Description   string `json:"description"`
}

type UpdateProduitRequest struct {
	LoginBoutique string `json:"login_boutique" binding:"required"`
	Nom           string `json:"nom" binding:"required"`
	Categories    string `json:"categories,omitempty"`
	Reduction     string `json:"reduction,omitempty"`
	Prix          string `json:"prix,omitempty"`
	Description   string `json:"description,omitempty"`
}
