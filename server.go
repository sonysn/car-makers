package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"carmakes/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	server()
}

func server() {
	PORT := os.Getenv("PORT")
	http.HandleFunc("/", controllers.Handler)
	http.HandleFunc("/getAPIKEY", controllers.APIKEY)
	fmt.Println("Server listening on port:", PORT)
	port := ":" + PORT
	http.ListenAndServe(port, nil)
}
