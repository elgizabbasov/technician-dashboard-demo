package routes

import (
	"github.com/CleanO2tech/clean-technician-web/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	// API device routes look like /devices?id=<guid>
	dr := r.Group("/devices")
	{
		dr.GET("/", controllers.GetDevices)
		dr.GET("", controllers.GetDeviceByID) // here /devices?id=<guid>
		dr.GET("/names", controllers.GetDeviceByName)
		dr.POST("/", controllers.AddDevice)
		dr.PUT("", controllers.UpdateDevice)
		dr.DELETE("", controllers.DeleteDevice)
	}

	// API site routes will look like /sites?id=<guid>
	sr := r.Group("/sites")
	{
		sr.GET("/", controllers.GetSites)
		sr.GET("", controllers.GetSiteById)         // here /sites?id=<guid>
		sr.GET("/names", controllers.GetSiteByName) // here /sites/names?name=<name>
		sr.POST("/", controllers.AddSite)
		sr.PUT("", controllers.UpdateSite)
		sr.DELETE("", controllers.DeleteSite)
	}

	// API sensor routes
	snr := r.Group("/sensors")
	{
		snr.GET("/", controllers.GetSensors)
		snr.GET("", controllers.GetSensorById)
		snr.GET("/deviceid", controllers.GetSensorByDeviceId)
		snr.GET("/names", controllers.GetSensorByName)
		snr.POST("/", controllers.AddSensor)
		snr.PUT("", controllers.UpdateSensor)
		snr.DELETE("", controllers.DeleteSensor)
	}

	// API Sensor Data routes
	sd := r.Group("/sensordata")
	{
		sd.GET("/", controllers.GetSensorData)
		//get based on created not id
		sd.GET("/created", controllers.GetSensorDataByDate)
		sd.GET("/latest", controllers.GetSensorDataLatest)
		sd.POST("/", controllers.AddSensorData)
		// no PUT method needed for sensordata
		sd.DELETE("", controllers.DeleteSensorData)
	}

	// API service data, where form data is data attribute, routes
	svr := r.Group("/services")
	{
		svr.GET("/", controllers.GetServices)
		svr.GET("", controllers.GetServiceById)
		svr.GET("/deviceid", controllers.GetServiceByDeviceId)
		svr.POST("/", controllers.AddService)
		svr.PUT("", controllers.UpdateService)
		svr.DELETE("", controllers.DeleteService)
	}

	// API service data, where form data is data attribute, routes
	sdr := r.Group("/servicedata")
	{
		sdr.GET("/", controllers.GetServiceData)
		sdr.GET("/created", controllers.GetServiceDataByDate)
		sdr.GET("/serviceid", controllers.GetServiceDataByServiceId)
		sdr.POST("/", controllers.AddServiceData)
		sdr.PUT("", controllers.UpdateServiceData)
		sdr.DELETE("", controllers.DeleteServiceData)
	}

	// API service data images, stored as urls to path
	sdir := r.Group("/images")
	{
		// sdir.GET("/", controllers.GetServiceDataImages)
		sdir.POST("/", controllers.AddServiceDataImage)
		// sdir.PUT("", controllers.UpdateServiceDataImage)
		// sdir.DELETE("", controllers.DeleteServiceDataImage)
	}

}
