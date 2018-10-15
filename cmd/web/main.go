package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new serve mux (router)
	mux := http.NewServeMux()

	// Use the mux.HandleFunc() method to register our Home function in handler.go
	// as the handler for the "/" URL pattern.
	mux.HandleFunc("/", Home)

	// register ShowSymptoms function to handle "/symptoms" URL pattern
	mux.HandleFunc("/symptoms", ShowSymptoms)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":3030")
	// and the serve mux we just created. If ListenAndServe() returns an error we
	// use the log.Fatal() function to record the error message and exit the program.
	log.Println("Starting server on :3030")
	err := http.ListenAndServe(":3030", mux)
	log.Fatal(err)
}
