package main

import (
	"log"
	"net/http"
	"os"
	handler "university-search/handlers"
)

//The app's main function
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("$PORT has not been set, defaulting to: %s", handler.DEFAULT_PORT)
		port = handler.DEFAULT_PORT
	}

	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.UNI_INFO_PATH, handler.UniversitiesHandler)
	http.HandleFunc(handler.COUNTRIES_PATH, handler.CountriesHandler)
	http.HandleFunc(handler.DIAGNOSTICS_PATH, handler.DiagHandler)

	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
	//http.HandleFunc()
}
