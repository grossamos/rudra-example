package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
    Message string
    Status string
}

// Base URL Information godoc
// @Summary      provide basic info on this API
// @Description  get information on rudra example API
// @Produce      json
 //@Success      200 {object}  controller.BaseResponse
// @Router       / [get]
func BaseHandler(ctx *gin.Context) {
    payload := BaseResponse{Message: "Welcome to the rudra example service", Status: "healthy"}
    ctx.JSON(http.StatusOK, payload)
}


func Authorize() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization");
        if auth == "" {
            c.AbortWithStatus(401)
            return
        }

        if !strings.HasPrefix(auth, "Basic") {
            c.AbortWithStatus(403)
            return
        }
        auth = auth[6:]

        if auth == "Admin" {
            c.Set("role", "admin")
        } else if auth == "User" {
            c.Set("role", "user")
        } else {
            c.AbortWithStatus(403)
            return
        }

        c.Next()
    }
}
