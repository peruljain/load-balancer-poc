package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handler for the /healthCheck endpoint
	http.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		// Sending a response to indicate the status is OK
		log.Printf("health check successfull")
		fmt.Fprintf(w, "OK")
	})

	// Starting the server on port 8080
	fmt.Println("Server is running on port 8083")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
