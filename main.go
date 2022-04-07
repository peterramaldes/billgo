package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handlers(mux)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
