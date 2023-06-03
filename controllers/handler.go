package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Vehicle struct {
	MAKE   string   `json:"MAKE"`
	MODELS []string `json:"MODELS"`
}

type Keys struct {
	KEY string
}

var timesRequested int

func increment() {
	timesRequested++
}

func Handler(res http.ResponseWriter, req *http.Request) {

	data, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(res, string(data))
	var jsonbody []Vehicle
	if err := json.Unmarshal(data, &jsonbody); err != nil {
		log.Fatal(err)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(jsonbody)
	increment()
	fmt.Printf("This service has been used %d times\n", timesRequested)

	// // fmt.Println(req.Body)
	// v, _ := ioutil.ReadAll(req.Body)
	// var jsonN map[string]string
	// if err := json.Unmarshal(v, &jsonN); err != nil {
	// 	fmt.Printf("could not marshal json: %s\n", err)
	// 	return
	// }

	// fmt.Printf("json data: %s\n", jsonN)
}

func APIKEY(res http.ResponseWriter, req *http.Request) {
	apiKey := os.Getenv("APIKEY")
	//fmt.Println(APIKEY)

	data := Keys{apiKey}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}
