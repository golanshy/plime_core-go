package client_dto

type ClientCredentials struct {
	ClientId   string `json:"client_id"`
	ClientName string `json:"client_name"`
	AppName    string `json:"app_name"`
	AppDetails string `json:"app_details"`
	ClientSecret string `json:"client_secret"`
}

type ClientCredentialsRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}