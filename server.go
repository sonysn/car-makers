package main

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

func main() {
	fmt.Println("Hello world")
	server()
}

func server() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 5000")
	http.ListenAndServe("localhost:5000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {

	data, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(res, string(data))
	var jsonbody []Vehicle
	if err := json.Unmarshal(data, &jsonbody); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(jsonbody)
}
