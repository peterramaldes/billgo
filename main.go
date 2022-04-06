package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handlers(mux)
	http.ListenAndServe(":3000", mux)
}
