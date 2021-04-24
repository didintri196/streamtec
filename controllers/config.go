package controllers

import (
	"net/http"
	"streamtec/models"
	"streamtec/responses"

	"github.com/gin-gonic/gin"
)

type ConfigController struct{}

var configmodel = new(models.ConfigModel)

func (A *ConfigController) CreateUpdateHandler(c *gin.Context) {
	var body models.ConfigModel

	// App level validation
	bindErr := c.BindJSON(&body)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, &responses.ConfigDefault{
			Status:  "ERR",
			Message: bindErr.Error(),
		})
		return
	}
	// Inserting data
	insertErr := configmodel.CreateUpdate(body)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, &responses.ConfigDefault{
			Status:  "ERR",
			Message: insertErr.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.ConfigDefault{
			Status:  "SUCCESS",
			Message: "Success update data",
		})
		return
	}

}

func (A *ConfigController) ShowAllHandler(c *gin.Context) {
	data, err := configmodel.Read()
	// Check if resource exist
	if err != nil {
		c.JSON(http.StatusNotFound, &responses.ConfigDefault{
			Status:  "ERR",
			Message: err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, &responses.ConfigData{
			Status:  "SUCCESS",
			Message: "Berhasil menampilkan data",
			Data:    data,
		})
		return
	}
}
