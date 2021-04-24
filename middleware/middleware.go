package middleware

import (
	"flag"
	"net/http"
	"streamtec/controllers"
	"streamtec/util"

	cors "github.com/itsjamie/gin-cors"

	"github.com/gin-gonic/gin"
)

var (
	cameracontroll = new(controllers.CameraController)
	configcontroll = new(controllers.ConfigController)
)

func Middleware() {

	flag.Parse()
	dir := util.GetDir()

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Token",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.Static("/public", dir+"/assets/")

	linkcamera := router.Group("/camera")
	{
		linkcamera.GET("", cameracontroll.ShowAllHandler)
		linkcamera.POST("", cameracontroll.CreateUpdateHandler)
		linkcamera.GET("/stream/restart", cameracontroll.StartStreamHandler)
	}

	linkconfig := router.Group("/config")
	{
		linkconfig.POST("", configcontroll.CreateUpdateHandler)
		linkconfig.GET("", configcontroll.ShowAllHandler)
	}

	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API STREAMTEC ON"})
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Page Not Found")
	})
	router.Run(":8881")
}

func StartEngine() {
	cameracontroll.StartStream()
}
