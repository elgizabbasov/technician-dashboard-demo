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

func GetSensorData(c *gin.Context) {
	sensordata, err := mappers.SensorDataGet()

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	if sensordata == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No SensorData Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensordata})
	}
}

func GetSensorDataByDate(c *gin.Context) {
	//search the query for two strings to use as start and end date respectively (returns strings)
	startdatetime := c.Query("startdatetime")
	enddatetime := c.Query("enddatetime")

	sensordata, err := mappers.SensorDataGetByDate(startdatetime, enddatetime)

	if err != nil {
		log.Fatal(err)
	}

	if sensordata == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No SensorData Records Found For Given Range"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sensordata})
	}

}

// Sensor data GET call for latest 100 rows
func GetSensorDataLatest(c *gin.Context) {
	sensorID := c.Query("sensorid")
	rows := c.Query("rows")
	servicedata, err := mappers.SensorDataGetLatest(rows, sensorID)

	if err != nil {
		log.Fatal(err)
	}

	if servicedata == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Latest SensorData Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": servicedata})
	}
}

func AddSensorData(c *gin.Context) {
	var json models.SensorData

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
	// json.Updated = currentTime

	success, err := mappers.SensorDataAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "SensorData added"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateSensorData() {
}

func DeleteSensorData(c *gin.Context) {
	// TODO: add error handling
	siteID := c.Query("id")

	success, err := mappers.SensorDataDelete(siteID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "SensorData deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
