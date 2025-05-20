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

type GetClientRequest struct {
	Login string `json:"login" binding:"required"`
}

type Client struct {
	Login         string `json:"login"`
	Password      string `json:"password"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Email         string `json:"email"`
	DateNaissance string `json:"date_naissance"`
	Telephone     string `json:"telephone"`
}

type UpdateClientRequest struct {
	Login         string `json:"login" binding:"required"`
	Password      string `json:"password,omitempty"`
	Nom           string `json:"nom,omitempty"`
	Prenom        string `json:"prenom,omitempty"`
	Email         string `json:"email,omitempty"`
	DateNaissance string `json:"date_naissance,omitempty"`
	Telephone     string `json:"telephone,omitempty"`
}
