package models

type AddBoutiqueRequest struct {
	Login               string `json:"login" binding:"required"`
	Password            string `json:"password" binding:"required"`
	Nom                 string `json:"nom" binding:"required"`
	Email               string `json:"email" binding:"required"`
	Telephone           string `json:"telephone"`
	Siret               string `json:"siret" binding:"required"`
	Forme_juridique     string `json:"forme_juridique" binding:"required"`
	Siege_social        string `json:"siege_social"`
	Pays_enregistrement string `json:"pays_enregistrement"`
	Iban                string `json:"iban"`
	Image               []byte `json:"image"`
}

type GetBoutiqueRequest struct {
	Login string `json:"login" binding:"required"`
}

type Boutique struct {
	Login               string `json:"login"`
	Password            string `json:"password"`
	Nom                 string `json:"nom"`
	Email               string `json:"email"`
	Telephone           string `json:"telephone"`
	Siret               string `json:"siret"`
	Forme_juridique     string `json:"forme_juridique"`
	Siege_social        string `json:"siege_social"`
	Pays_enregistrement string `json:"pays_enregistrement"`
	Iban                string `json:"iban"`
	Image               []byte `json:"image"`
}

type UpdateBoutiqueRequest struct {
	Login               string `json:"login" binding:"required"`
	Password            string `json:"password,omitempty"`
	Nom                 string `json:"nom,omitempty"`
	Email               string `json:"email,omitempty"`
	Telephone           string `json:"telephone,omitempty"`
	Siret               string `json:"siret,omitempty"`
	Forme_juridique     string `json:"forme_juridique,omitempty"`
	Siege_social        string `json:"siege_social,omitempty"`
	Pays_enregistrement string `json:"pays_enregistrement,omitempty"`
	Iban                string `json:"iban,omitempty"`
	Image               []byte `json:"image,omitempty"`
}
