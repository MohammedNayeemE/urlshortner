package main

import (
	"log" 
	"net/http"
	"github.com/MohammedNayeemE/url-shortner/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	
	if err:= godotenv.Load() ; err != nil {
		log.Println("No .env file found");
	}

	baseUrl := os.Getenv("BASE_URL");

	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
		log.Fatal("BASE_URL not found");
	}

	handler.Init(baseUrl);

	router := mux.NewRouter();

	router.HandleFunc("/shorten" , handler.ShortenUrl).Methods("POST");
	router.HandleFunc("/{shorturl}" , handler.RedirectUrl).Methods("GET");

	log.Println("server running at 8080");

	log.Fatal(http.ListenAndServe(":8080" , router));

}
