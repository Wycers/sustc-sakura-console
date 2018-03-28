package model

type LoginRequest struct {
	Action string `json:"action"`
	Username string `json:"username"`
	Passwrod string `json:"password"`
}

type TransRequest struct {
	Action     string `json:"action"`
	JSESSIONID string `json:"JSESSIONID"`
	Week       int `json:"week"`
}

type DownloadRequest struct {
	Week     int  `json:"week"`
}