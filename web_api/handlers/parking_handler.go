package handlers

import (
	extracterServices "challenge/extractor/services"
	"challenge/models"
	"challenge/web_api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GetParking(c *gin.Context) {
	err := extracterServices.ExtractParkingData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "Error",
			Data:    err.Error(),
			Message: "Error: unable to extract information!",
		})
		return
	}

	fmt.Printf(os.Getenv("APP_ENV"))

	req := models.ParkingRequest{}

	err = c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "Error",
			Data:    err.Error(),
			Message: "Error: unmarshaling request!",
		})
		return
	}

	response, err := services.GetParking(req.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Data:    err.Error(),
			Message: "Error: unable to get parking!",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
