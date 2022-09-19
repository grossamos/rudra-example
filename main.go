package main

import (
    "fmt"
    "rudra-example/controller"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "rudra-example/docs"
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

    autorized := r.Group("/")
    autorized.Use(controller.Authorize())
    autorized.POST("/validate", controller.ValidateWeather)


    r.GET("/", controller.BaseHandler)
    r.GET("/weather", controller.GetCurrentWeather)
    r.GET("/ping", controller.PingWeatherService)

    r.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    fmt.Println("Running on port 8080")
    r.Run(":8080")
}
