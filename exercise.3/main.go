package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type cityWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  int     `json:"temp_min"`
		TempMax  int     `json:"temp_max"`
	} `json:"main"`

	Visibility int `json:"visibility"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Dt int `json:"dt"`

	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`

	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func singleCity(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	vars := mux.Vars(r)
	city := getCity(vars["name"])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v\n%vc %v", city.Name, int(math.Round(city.Main.Temp)), city.Weather[0].Description)

	end := time.Now()
	fmt.Fprintf(w, "\n%v", end.Sub(start))
}

func allCity(w http.ResponseWriter, r *http.Request) {

	cityName := []string{
		"bangkok",
		"hobart",
		"nairobi",
		"newyork",
		"kupang",
	}

	start := time.Now()

	var allCity []*cityWeather

	for i := range cityName {
		allCity = append(allCity, getCity(cityName[i]))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	for i := range allCity {
		showMsg(w, allCity[i])
	}

	end := time.Now()
	fmt.Fprintf(w, "\n%v", end.Sub(start))
}

func getCity(citiName string) *cityWeather {
	url := "http://localhost:8882/api/v1/weather/" + citiName
	rep, _ := http.Get(url)

	city := new(cityWeather)
	json.NewDecoder(rep.Body).Decode(city)
	return city
}

func showMsg(w http.ResponseWriter, city *cityWeather) {
	fmt.Fprintf(w, "%v\n%vc %v\n\n", city.Name, int(math.Round(city.Main.Temp)), city.Weather[0].Description)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/all", allCity).Methods("GET")
	r.HandleFunc("/weather/{name}", singleCity).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
