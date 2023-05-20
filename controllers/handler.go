package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Vehicle struct {
	MAKE   string   `json:"MAKE"`
	MODELS []string `json:"MODELS"`
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

	// // fmt.Println(req.Body)
	// v, _ := ioutil.ReadAll(req.Body)
	// var jsonN map[string]string
	// if err := json.Unmarshal(v, &jsonN); err != nil {
	// 	fmt.Printf("could not marshal json: %s\n", err)
	// 	return
	// }

	// fmt.Printf("json data: %s\n", jsonN)
}
