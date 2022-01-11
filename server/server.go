package main

import (
	"log"
	"net/http"

	"github.com/danstanke/simple-get-deviation-app/server/app/handlers"
)

func main() {
	log.Println("Server start...")

	http.HandleFunc("/random/mean", handlers.RandomMeanHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
