package main

import (
	"log"
	"net/http"
)

func randomMeanHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	numberOfRequests := query.Get("requests")
	numberOfIntegers := query.Get("length")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(numberOfRequests + " " + numberOfIntegers))
}

func main() {
	log.Println("Server start...")

	http.HandleFunc("/random/mean", randomMeanHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
