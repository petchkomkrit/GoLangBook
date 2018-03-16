package main

import (
	"encoding/json"
	"fmt"
	"math"
	"mux"
	"net/http"
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

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if vars["name"] == "" {
		name = "All"
	}

	url := "http://localhost:8882/api/v1/weather/" + name
	rep, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "ERROR, %s", err)
	}

	city := new(cityWeather)
	json.NewDecoder(rep.Body).Decode(city)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v\n%vc %v", city.Name, int(math.Round(city.Main.Temp)), city.Weather[0].Description)

	/*
		if err != nil {
			fmt.Fprintf(w, "ERROR, %s!", err)
		}

		fmt.Fprintf(w, "%s", rep.Body)

		fmt.Fprintf(w, "Hello, %s!", name) */
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{name}", HomePageHandle).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
