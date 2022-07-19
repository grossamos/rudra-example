package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"rudra-example/util"

	"github.com/gin-gonic/gin"
)

var weatherStates = []string{"sunny", "windy", "rainy", "stormy"}

type IsValid struct {
    Valid bool
}

type WeatherStatus struct {
    Status string
}

// Fetch Current Weather godoc
// @Summary      provides current weather
// @Description  gets current weather from a set of states
// @Produce      json
// @Success      200  {object}  controller.WeatherStatus
// @Router       /weather [get]
func GetCurrentWeather(c *gin.Context) {
    payload := WeatherStatus{Status: weatherStates[rand.Intn(len(weatherStates))]}
    c.JSON(http.StatusOK, payload)
}

// Validate Weather State godoc
// @Summary      validates weather state
// @Description  checks if a given state is a valid weather state
// @Security     BasicAuth 
// @Produce      json
// @Success      200  {object}  controller.IsValid
// @Failure      400  {object}  util.ErrorMessage
// @Router       /validate [post]
func ValidateWeather(c *gin.Context) {
    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        util.UnexplainedError(c)
        return
    }

    var input WeatherStatus
    err = json.Unmarshal(body, &input)

    if err != nil || len(input.Status) == 0 {
        util.UnexplainedError(c)
        return
    }

    fmt.Println(input.Status)

    isValid := IsValid{Valid: false}

    for _, val := range weatherStates {
        if val == input.Status {
            isValid.Valid = true
        }
    }
    c.JSON(http.StatusOK, isValid)
}

