package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"rudra-example/util"
)

var weatherStates = []string{"sunny", "windy", "rainy", "stormy"}

type IsValid struct {
    Valid bool
}

type WeatherStatus struct {
    Status string
}

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
    payload := WeatherStatus{Status: weatherStates[rand.Intn(len(weatherStates))]}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(payload)
}

func ValidateWeather(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        util.UnexplainedError(w)
        return
    }

    var input WeatherStatus
    err = json.Unmarshal(body, &input)

    if err != nil || len(input.Status) == 0 {
        util.UnexplainedError(w)
        return
    }

    fmt.Println(input.Status)

    isValid := IsValid{Valid: false}

    for _, val := range weatherStates {
        if val == input.Status {
            isValid.Valid = true
        }
    }

    json.NewEncoder(w).Encode(isValid)
}

