package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type payload struct {
	Token string `json:"token"`
	Data  string `json:"data"`
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from Forms!")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	var data payload
	if err := json.Unmarshal(d, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	log.Printf("Received data: %+v", data)
	fmt.Fprintf(w, "Form submitted successfully!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/submit", submitHandler)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
