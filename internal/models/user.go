package models

type AddUserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
}