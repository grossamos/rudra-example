package util

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
    Message string
}


func UnexplainedError(w http.ResponseWriter) {
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(ErrorMessage{Message: "invalid request"})
}

func SetContentTypeJsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        next.ServeHTTP(w, r)
    })
}
