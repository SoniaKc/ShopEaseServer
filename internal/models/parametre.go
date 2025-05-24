package models

type AddParametreRequest struct {
	Login         string `json:"login" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Langue        string `json:"langue" binding:"required"`
	Cookies       string `json:"cookies" binding:"required"`
	Notifications string `json:"notifications" binding:"required"`
}

type GetParametreRequest struct {
	Login string `json:"login" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

type Parametre struct {
	Login         string `json:"login"`
	Type          string `json:"type"`
	Langue        string `json:"langue"`
	Cookies       string `json:"cookies"`
	Notifications string `json:"notifications"`
}

type UpdateParametreRequest struct {
	Login         string `json:"login" binding:"required"`
	Type          string `json:"type"  binding:"required"`
	Langue        string `json:"langue,omitempty"`
	Cookies       string `json:"cookies,omitempty"`
	Notifications string `json:"notifications,omitempty"`
}
