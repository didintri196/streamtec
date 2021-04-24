package controllers

import (
	"net/http"
	"streamtec/models"
	"streamtec/responses"

	"github.com/gin-gonic/gin"
)

type CameraController struct{}

var cameramodel = new(models.CameraModel)

func (A *CameraController) CreateUpdateHandler(c *gin.Context) {
	var body []models.CameraModel

	// App level validation
	bindErr := c.BindJSON(&body)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.CameraDefault{
			Status:  "ERR",
			Message: bindErr.Error(),
		})
		return
	}
	// Inserting data
	insertErr := cameramodel.CreateUpdate(body)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, &responses.CameraDefault{
			Status:  "ERR",
			Message: insertErr.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, &responses.CameraDefault{
			Status:  "SUCCESS",
			Message: "Success create data",
		})
		return
	}

}

func (A *CameraController) ShowAllHandler(c *gin.Context) {
	host, data, err := cameramodel.Read()
	// Check if resource exist
	if err != nil {
		c.JSON(http.StatusNotFound, &responses.CameraDefault{
			Status:  "ERR",
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.CameraData{
			Status:  "SUCCESS",
			Message: "Berhasil menampilkan data",
			Data:    data,
			Host:    host,
		})
		return
	}
}

func (A *CameraController) StartStreamHandler(c *gin.Context) {
	err := cameramodel.StartStream()
	// Check if resource exist
	if err != nil {
		c.JSON(http.StatusNotFound, &responses.CameraDefault{
			Status:  "ERR",
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.CameraDefault{
			Status:  "SUCCESS",
			Message: "Berhasil Restart Stream",
		})
		return
	}
}

func (A *CameraController) StartStream() {
	cameramodel.StartStream()
}
