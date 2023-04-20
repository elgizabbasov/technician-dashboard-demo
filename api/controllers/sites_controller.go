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

func GetSites(c *gin.Context) {
	sites, err := mappers.SitesGet()

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	if sites == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Site Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sites})
	}
}

func GetSiteById(c *gin.Context) {
	id := c.Query("id")

	site, err := mappers.SiteGetByID(id)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if site.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Site Records Found For Given ID"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": site})
	}
}

func GetSiteByName(c *gin.Context) {
	name := c.Query("name")

	site, err := mappers.SiteGetByName(name)

	// TODO: move below to utils/ or errors/ as a function
	if err != nil {
		log.Fatal(err)
	}

	// if the name is blank we can assume nothing is found
	if site.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Site Records Found For Given Name", "exists": false})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": site, "exists": true})
	}
}

func AddSite(c *gin.Context) {
	var json models.Site

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

	success, err := mappers.SiteAdd(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Site added", "id": uuid})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateSite(c *gin.Context) {
	var json models.Site

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: add error handling
	siteID := c.Query("id")

	// Set created and updated times for new and updated sites
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	json.Updated = currentTime

	success, err := mappers.SiteUpdate(json, siteID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Site updated"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteSite(c *gin.Context) {
	// TODO: add error handling
	siteID := c.Query("id")

	success, err := mappers.SiteDelete(siteID)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Site deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
