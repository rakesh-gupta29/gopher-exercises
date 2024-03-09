package main

import (
	"fmt"
	"log"
	"net/http"
)

type CustomHandler struct {
	URLMap   map[string]string
	fallback http.Handler
}

type FallbackHandler struct{}

func (fallback *FallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fallback for the router called "))
}
func (handler *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if destURL, ok := handler.URLMap[r.URL.Path]; ok {
		http.Redirect(w, r, destURL, http.StatusPermanentRedirect)
	} else {
		handler.fallback.ServeHTTP(w, r)
	}
}

var URLMap = map[string]string{
	"/google": "https://www.google.com",
	"/github": "https://www.github.com",
}

func main() {

	fallbackHandler := &FallbackHandler{}

	customHandler := &CustomHandler{
		URLMap:   URLMap,
		fallback: fallbackHandler,
	}
	mux := http.NewServeMux()

	mux.Handle("/", customHandler)
	if err := http.ListenAndServe("localhost:4000", mux); err != nil {
		log.Fatal("failed to start the server:", err)
	} else {
		fmt.Println("Server listening on port 4000...")

	}
}
