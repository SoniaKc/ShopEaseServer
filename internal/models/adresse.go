package models

type AddAdresseRequest struct {
	Login      string `json:"login" binding:"required"`
	NomAdresse string `json:"nom_adresse" binding:"required"`
	Numero     string `json:"numero" binding:"required"`
	NomRue     string `json:"nom_rue" binding:"required"`
	CodePostal string `json:"code_postal" binding:"required"`
	Ville      string `json:"ville" binding:"required"`
	Pays       string `json:"pays" binding:"required"`
}

type GetAdresseRequest struct {
	Login      string `json:"login" binding:"required"`
	NomAdresse string `json:"nom_adresse" binding:"required"`
}

type Adresse struct {
	Login      string `json:"login"`
	NomAdresse string `json:"nom_adresse"`
	Numero     string `json:"numero"`
	NomRue     string `json:"nom_rue"`
	CodePostal string `json:"code_postal"`
	Ville      string `json:"ville"`
	Pays       string `json:"pays"`
}

type UpdateAdresseRequest struct {
	Login      string `json:"login" binding:"required"`
	NomAdresse string `json:"nom_adresse" binding:"required"`
	Numero     string `json:"numero,omitempty"`
	NomRue     string `json:"nom_rue,omitempty"`
	CodePostal string `json:"code_postal,omitempty"`
	Ville      string `json:"ville,omitempty"`
	Pays       string `json:"pays,omitempty"`
}
