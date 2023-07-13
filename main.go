package main

import (
	"log"
	"os"

	"github.com/elgizabbasov/technician-dashboard-demo/initializers"
	"github.com/elgizabbasov/technician-dashboard-demo/models"
	"github.com/elgizabbasov/technician-dashboard-demo/routes"
	"github.com/elgizabbasov/technician-dashboard-demo/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// initializers.LoadEnvVariables()
	// err := initializers.CreateDB()
	err := initializers.ConnectDB()

	checkErr(err)
}

type SensorData = models.SensorData

func main() {
	// initialize a Gin router
	r := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://technician-dashboard-demo.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// **** Creating SQLITE mock DB tables below ****
	utils.CreateSiteTable()
	utils.CreateDeviceTable()
	utils.CreateSensorTable()
	utils.CreateSensorDataTable()
	utils.CreateServiceTable()
	utils.CreateServiceDataTable()

	utils.PopulateTables()

	// Set up the routes
	routes.SetUpRoutes(r)

	// Start the server
	r.Run(":" + os.Getenv("PORT"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
