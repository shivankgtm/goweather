package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func get_weather(city_name string) (float64, error) {

	APIKEY := os.Getenv("WEATHER_API_KEY")
	base_url := "http://api.weatherapi.com/v1/current.json?"
	final_url := base_url + "key=" + APIKEY + "&q=" + city_name

	resp, err := http.Get(final_url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return 0, err
	}

	current, ok := data["current"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("Invalid JSON response")
	}

	temp_c, ok := current["temp_c"].(float64)
	if !ok {
		return 0, fmt.Errorf("Temperature not found in response")
	}

	return temp_c, nil
}

func main() {
	args := os.Args[1:]
	temp, _ := get_weather(args[0])
	fmt.Println("Temp in", args[0], "is", temp)
}
