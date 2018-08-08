package core

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

		// Log request
		log.Printf("Request: %s -> %s", r.Method, r.URL)

		// Handle special cases:
		// ---------------------

		// Handle Favicon requests...
		if r.URL.Path == "/favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			return
		}

		// Handle OPTIONS requests...
		if r.Method == "OPTIONS" {
			requestHeaders := r.Header["Access-Control-Request-Headers"]
			w.Header().Set("Allow", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(requestHeaders, ","))
			return
		}

		// ---------------------

		// Set response headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Match incoming request
		matchedRoute, err := matchRequestToRoutes(r, routes)

		if err != nil {
			switch err {
			case NoConfigMatchFoundError:
				w.WriteHeader(404)
			}
		} else {
			// Serve JSON response
			w.WriteHeader(matchedRoute["status"].(int))
			json.NewEncoder(w).Encode(matchedRoute["data"])
		}

	})

	address := fmt.Sprintf(":%d", port)
	log.Printf("Phony JSON Server running on: 0.0.0.0:%d", port)
	log.Println("-------------------------------------------------")
	log.Fatal(http.ListenAndServe(address, nil))
}

// Match the path string of the incoming URL against that of a stored Route.
func MatchPath(requestPath string, routePath string) bool {
	return requestPath == routePath
}

// Match the (http) method of a incoming request against that of a stored Route.
func MatchMethod(requestMethod string, routeMethod string) bool {
	return req.Method == strings.ToUpper(route.Method)
}

func matchRequestToRoutes(req *http.Request, routes []Route) (MatchedRoute, error) {
	var err error
	matchedRoute := make(MatchedRoute)

	// Loop over all routes and determine if incoming URL path and http method
	// matches a path + method combination specified in a route.
	for _, route := range routes {
		pathMatches := MatchPath(req.URL.Path, route.Path)
		methodMatches := MatchMethod(req, route)
		if pathMatches && methodMatches {
			matchedRoute["status"] = route.GetStatus()
			matchedRoute["data"] = route.Data
		}
	}

	// len(<map>) - give us the number of keys specified on the map.
	// If it's zero, then we know the map is empty...
	if len(matchedRoute) == 0 {
		err = NoConfigMatchFoundError
	}

	return matchedRoute, err
}
