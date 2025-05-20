package models

type AddParametreRequest struct {
	Login         string `json:"login" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Langue        string `json:"langue" binding:"required"`
	Cookies       string `json:"cookies" binding:"required"`
	Notifications string `json:"notifications" binding:"required"`
}
