package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Helper to handle CORS and send JSON
func respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow Next.js to connect
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Helper to receive JSON
func decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func main() {
	// Scenario 1: Simple Text
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" { return }
		var req struct{ Name string `json:"name"` }
		decode(r, &req)
		respond(w, map[string]string{"message": "Hello, " + req.Name})
	})

	// Scenario 2: Math
	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" { return }
		var req struct{ A, B int }
		decode(r, &req)
		respond(w, map[string]int{"result": req.A + req.B})
	})

	// Scenario 3: Array processing
	http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" { return }
		var req struct{ Numbers []int `json:"numbers"` }
		decode(r, &req)
		sum := 0
		for _, n := range req.Numbers { sum += n }
		respond(w, map[string]interface{}{"sum": sum, "count": len(req.Numbers)})
	})

	// Scenario 4: GET Request (Fetching data)
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" { return }
		respond(w, map[string]string{"status": "online", "server": "Go 1.21"})
	})

	// Scenario 5: Error Handling
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" { return }
		var req struct{ Password string `json:"password"` }
		decode(r, &req)
		if req.Password != "secret123" {
			w.WriteHeader(http.StatusUnauthorized)
			respond(w, map[string]string{"error": "Invalid password"})
			return
		}
		respond(w, map[string]string{"message": "Access granted"})
	})

	log.Println("Go server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}