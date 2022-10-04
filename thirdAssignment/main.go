package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Status Data `json:"status"`
}

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(autoReload)

	mux.Handle("/auto-reload", middleware(endpoint))

	fmt.Println("Listening to port 3030")
	err := http.ListenAndServe(":3030", mux)

	log.Fatal(err)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Using middleware")
		next.ServeHTTP(w, r)
	})
}

func autoReload(w http.ResponseWriter, r *http.Request) {
	for {
		min := 1
		max := 100
		rand.Seed(time.Now().Unix())
		numberWater := rand.Intn(max - min)
		numberWind := rand.Intn(max - min)

		updateData(numberWater, numberWind)
		logWater(numberWater, numberWind)
		logWind(numberWater, numberWind)

		time.Sleep(time.Second * 15)
	}
}

func updateData(numberWater int, numberWind int) {
	data := Data{}
	data.Water = numberWater
	data.Wind = numberWind
	dataWaterWind := Status{
		Status: data,
	}

	jsonprint, err := json.MarshalIndent(dataWaterWind, "", "    ")
	if err != nil {
		fmt.Println("json error")
		return
	}

	err = ioutil.WriteFile("file.json", jsonprint, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonprint))
}

func logWater(numberWater int, numberWind int) {
	if numberWater < 6 {
		result := "Aman"
		fmt.Println("Status Water: ", result)
	} else if numberWater > 6 && numberWater <= 15 {
		result := "Siaga"
		fmt.Println("Status Water: ", result)
	} else if numberWater > 15 {
		result := "Bahaya"
		fmt.Println("Status Water: ", result)
	} else {
		fmt.Println("Measurement not found")
	}
}

func logWind(numberWater int, numberWind int) {
	if numberWind < 6 {
		result := "Aman"
		fmt.Println("Status Wind: ", result)
	} else if numberWind > 6 && numberWind <= 15 {
		result := "Siaga"
		fmt.Println("Status Wind: ", result)
	} else if numberWind > 15 {
		result := "Bahaya"
		fmt.Println("Status Wind: ", result)
	} else {
		fmt.Println("Measurement not found")
	}
}
