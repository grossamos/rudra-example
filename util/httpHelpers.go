package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
    Message string
}


func UnexplainedError(c *gin.Context) {
    payload := ErrorMessage{Message: "invalid request"}
    c.JSON(http.StatusBadRequest, payload)
}

