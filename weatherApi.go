package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const apiKey = "04f16be0efe531bd5bad74fcfaa01a74"

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/weather/", weatherCheck)
	http.ListenAndServe(":8000", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func weatherCheck(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	data, err := query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content=Type", "application/json, charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?appid=" + apiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	var d weatherData

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}
