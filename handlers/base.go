package handlers

import (
	"encoding/json"
	"net/http"
)


type BaseResponse struct {
    Message string
    Status string
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
    payload := BaseResponse{Message: "Welcome to the rudra example service", Status: "healthy"}
    json.NewEncoder(w).Encode(payload)
}
