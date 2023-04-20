package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elgizabbasov/technician-dashboard-demo/mappers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetServiceData(c *gin.Context) {
	servicedata, err := mappers.ServiceDataGet()

	if err != nil {
		log.Fatal(err)
	}

	if servicedata == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ServiceData Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": servicedata})
	}
}

func GetServiceDataByDate(c *gin.Context) {
	//search the query for two strings to use as start and end date respectively (returns strings)
	startdatetime := c.Query("startdatetime")
	enddatetime := c.Query("enddatetime")

	servicedata, err := mappers.ServiceDataGetByDate(startdatetime, enddatetime)

	if err != nil {
		log.Fatal(err)
	}

	if servicedata == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ServiceData Records Found For Given Range"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": servicedata})
	}

}

func GetServiceDataByServiceId(c *gin.Context) {
	serviceid := c.Query("serviceid")

	servicedata, err := mappers.ServiceDataGetByServiceID(serviceid)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	if len(servicedata) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ServiceData Records Found For Given Service ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": servicedata})
	}
}

func AddServiceData(c *gin.Context) {
	var json models.ServiceData

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new UUID for the servicedata
	// TODO: add error handling
	uuid := uuid.New().String()

	// Set the UUID in the servicedata struct
	json.ID = uuid

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	json.Created = currentTime
	json.Updated = currentTime
	success, err := mappers.ServiceDataAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "ServiceData added"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

// Send the image to be uploaded to a location
func AddServiceDataImage(c *gin.Context) {
	// Set the folder to save the uploaded file
	localFolder := "./uploaded_form_images/"

	// Parse the multipart form data *******
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the file data from the form data
	fileHeaders, ok := c.Request.MultipartForm.File["image"]

	if !ok || len(fileHeaders) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found in the request"})
		return
	}

	image := fileHeaders[0]

	// Add the service data image to the fileserver
	if err := mappers.ServiceDataImageAdd(image, localFolder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// url format: "/uploaded_form_images/<image_name.png>""
	imageURL := fmt.Sprintf("%s%s", localFolder[1:], image.Filename)

	c.JSON(http.StatusOK, gin.H{"message": "ServiceData added", "url": imageURL})
}

func UpdateServiceData(c *gin.Context) {
	var json models.ServiceData

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add error handling
	servicedataID := c.Query("id")

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	json.Updated = currentTime

	success, err := mappers.ServiceDataUpdate(json, servicedataID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "ServiceData updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteServiceData(c *gin.Context) {
	// TODO: add error handling
	siteID := c.Query("id")

	success, err := mappers.ServiceDataDelete(siteID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "ServiceData deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
