package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/CleanO2tech/clean-technician-web/mappers"
	"github.com/CleanO2tech/clean-technician-web/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDevices(c *gin.Context) {
	devices, err := mappers.DevicesGet()

	checkErr(err)

	if devices == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Device Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": devices})
	}
}

func GetDeviceByID(c *gin.Context) {
	id := c.Query("id")

	device, err := mappers.DeviceGetByID(id)

	checkErr(err)

	// if the name is blank we can assume nothing is found
	if device.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Device Records Found For Given ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": device})
	}
}

func GetDeviceByName(c *gin.Context) {
	name := c.Query("name")

	device, err := mappers.DeviceGetByName(name)

	checkErr(err)

	// if the name is blank we can assume nothing is found
	if device.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Device Records Found For Given Name", "exists": false})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": device, "exists": true})
	}
}

func AddDevice(c *gin.Context) {
	var json models.Device

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

	success, err := mappers.DeviceAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Device added", "id": uuid})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateDevice(c *gin.Context) {
	var json models.Device

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add error handling
	deviceID := c.Query("id")

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	json.Updated = currentTime

	success, err := mappers.DeviceUpdate(json, deviceID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Device updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteDevice(c *gin.Context) {
	// TODO: add error handling
	deviceID := c.Query("id")

	success, err := mappers.DeviceDelete(deviceID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
