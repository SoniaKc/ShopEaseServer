package models

type AddAdresseRequest struct {
	Login      string `json:"login" binding:"required"`
	NomAdresse string `json:"nom_adresse" binding:"required"`
	Numero     string `json:"numero" binding:"required"`
	NomRue     string `json:"nom_rue" binding:"required"`
	CodePostal int    `json:"code_postal" binding:"required"`
	Ville      string `json:"ville" binding:"required"`
	Pays       string `json:"pays" binding:"required"`
}
