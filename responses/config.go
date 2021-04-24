package responses

import (
	"streamtec/models"
)

type ConfigDefault struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ConfigData struct {
	Message string             `json:"message"`
	Status  string             `json:"status"`
	Data    models.ConfigModel `json:"data"`
}
