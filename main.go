/*
 * A weather web server. Uses go lanuage.
 * Based upon article: https://howistart.org/posts/go/1
 *
 * Nick Brandaleone - December 2015
 *
 * API Key: e17c8d8d35f1cf3504ad080013ab3bd6
*/
package main

import (
 	"encoding/json"
 	"net/http"
 	"strings"
 )
 
func main() {
	http.HandleFunc("/", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request){
		city := strings.SplitN(r.URL.Path,"/", 3)[2]
	
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query(city string)(weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=e17c8d8d35f1cf3504ad080013ab3bd6")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	//fmt.Printf("No Errors!")
	return d, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json"main"`
}
