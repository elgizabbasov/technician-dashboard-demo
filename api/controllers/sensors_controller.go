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

func GetSensors(c *gin.Context) {
	sensors, err := mappers.SensorsGet()

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	if sensors == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Sensor Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensors})
	}
}

func GetSensorById(c *gin.Context) {
	id := c.Query("id")

	sensor, err := mappers.SensorGetByID(id)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if sensor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Sensor Records Found For Given ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensor})
	}
}

func GetSensorByDeviceId(c *gin.Context) {
	deviceid := c.Query("deviceid")

	sensor, err := mappers.SensorGetByDeviceID(deviceid)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if len(sensor) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Sensor Records Found For Given Device ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensor})
	}
}

func GetSensorByName(c *gin.Context) {
	name := c.Query("name")

	sensor, err := mappers.SensorGetByName(name)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if sensor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Sensor Records Found For Given Name", "exists": false})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensor, "exists": true})
	}
}

func AddSensor(c *gin.Context) {
	var json models.Sensor

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

	success, err := mappers.SensorAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Sensor added", "id": uuid})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateSensor(c *gin.Context) {
	var json models.Sensor

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add error handling
	sensorID := c.Query("id")

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	json.Updated = currentTime

	success, err := mappers.SensorUpdate(json, sensorID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Sensor updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteSensor(c *gin.Context) {
	// TODO: add error handling
	sensorID := c.Query("id")

	success, err := mappers.SensorDelete(sensorID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Sensor deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
