package phony

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type MatchedRoute map[string]interface{}

var NoConfigMatchFoundError = errors.New("NoConfigMatchFoundError")

func ServeRoutes(routes []Route, port int) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Handle Favicon requests...
		if r.URL.Path == "/favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			return
		}

		// Set response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Log request
		log.Printf("Request: %s -> %s", r.Method, r.URL)

		// Match incoming request
		matchedRoute := matchRequestToRoutes(r, routes)

		// Serve JSON response
		w.WriteHeader(matchedRoute["status"].(int))
		json.NewEncoder(w).Encode(matchedRoute["data"])

	})

	address := fmt.Sprintf(":%d", port)
	log.Printf("Phony JSON Server running on: 0.0.0.0:%d", port)
	log.Println("-------------------------------------------------")
	log.Fatal(http.ListenAndServe(address, nil))
}

func matchRequestToRoutes(r *http.Request, routes []Route) MatchedRoute {
	response := make(MatchedRoute)

	// Loop over all routes and determine if incoming URL path and http method
	// matches a path + method combination specified in a route.
	for _, route := range routes {
		pathMatches := r.URL.Path == route.Path
		methodMatches := r.Method == strings.ToUpper(route.Method)
		if pathMatches && methodMatches {
			response["status"] = route.GetStatus()
			response["data"] = route.Data
		}
	}
	return response
}
