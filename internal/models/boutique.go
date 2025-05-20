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
}
