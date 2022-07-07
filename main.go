package main

import (
	"fmt"
	"log"
	"net/http"

	"rudra-example/handlers"
	"rudra-example/util"

	"github.com/gorilla/mux"
)


func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.BaseHandler).Methods("GET")
    router.HandleFunc("/weather", handlers.GetCurrentWeather).Methods("GET")
    router.HandleFunc("/validate", handlers.ValidateWeather).Methods("POST")

    router.Use(util.SetContentTypeJsonMiddleware)

    fmt.Println("Listening on port :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
