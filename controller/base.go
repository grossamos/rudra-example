package controller

import (
	"net/http"

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

