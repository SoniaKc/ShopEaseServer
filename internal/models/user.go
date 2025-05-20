package models

type AddClientRequest struct {
	Login          string `json:"login" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Nom            string `json:"nom" binding:"required"`
	Prenom         string `json:"prenom" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Date_naissance string `json:"date_naissance" binding:"required"`
	Telephone      string `json:"telephone"`
}
