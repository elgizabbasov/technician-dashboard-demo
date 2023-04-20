package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/elgizabbasov/technician-dashboard-demo/mappers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetServices(c *gin.Context) {
	services, err := mappers.ServicesGet()

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	if services == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Service Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": services})
	}
}

func GetServiceById(c *gin.Context) {
	id := c.Query("id")

	service, err := mappers.ServiceGetByID(id)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if service.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Service Records Found For Given ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": service})
	}
}

func GetServiceByDeviceId(c *gin.Context) {
	deviceid := c.Query("deviceid")

	service, err := mappers.ServiceGetByDeviceID(deviceid)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if service.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Sensor Records Found For Given Device ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": service})
	}
}

func AddService(c *gin.Context) {
	var json models.Service

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new UUID for the device
	// TODO: add error handling
	uuid := uuid.New().String()

	// Set the UUID in the device struct
	json.ID = uuid

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	json.Created = currentTime
	json.Updated = currentTime

	success, err := mappers.ServiceAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Service added"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateService(c *gin.Context) {
	var json models.Service

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add error handling
	serviceID := c.Query("id")

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	json.Updated = currentTime

	success, err := mappers.ServiceUpdate(json, serviceID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Service updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteService(c *gin.Context) {
	// TODO: add error handling
	serviceID := c.Query("id")

	success, err := mappers.ServiceDelete(serviceID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
