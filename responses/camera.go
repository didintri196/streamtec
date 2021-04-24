package responses

import (
	"streamtec/models"
)

type CameraDefault struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type CameraData struct {
	Message string               `json:"message"`
	Status  string               `json:"status"`
	Host    string               `json:"host"`
	Data    []models.CameraModel `json:"data"`
}
