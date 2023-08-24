package model

type ResponseLogin struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Token   string `json:"token"`
}
