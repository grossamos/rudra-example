package main

import (
	"fmt"
	"rudra-example/controller"

	"github.com/gin-gonic/gin"
)

// @title           Rudra Example Project
// @version         1.0
// @description     This is a sample project for the rudra test tool

// @license.name  BSD-2-Clause
// @license.url   https://raw.githubusercontent.com/grossamos/rudra/main/LICENSE

// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
// @BasePath  /
func main() {
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.GET("/", controller.BaseHandler)
    r.GET("/weather", controller.GetCurrentWeather)
    r.POST("/validate", controller.ValidateWeather)

    fmt.Println("Running on port 8080")
    r.Run(":8080")
}
