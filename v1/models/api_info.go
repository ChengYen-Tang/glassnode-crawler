package models

type APIInfo struct {
	APIUrl     string `json:"api_url"`
	SaveFolder string `json:"save_folder"`
	FileName   string `json:"file_name"`
}
