package main

import (
	"Decorator_Handle/app"
	"Decorator_Handle/deco"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := app.NewHandler()
	h := deco.NewDecoHandler(mux, logger)
	return h
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
