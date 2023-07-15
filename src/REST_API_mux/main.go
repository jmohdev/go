package main

import (
	"REST_API_mux/app_mux"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", app_mux.NewHandler())
}
