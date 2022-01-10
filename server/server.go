package main

import (
	"fmt"
	"log"

	"github.com/danstanke/simple-get-deviation-app/server/app/randorg"
)

func main() {
	inty, err := randorg.GetIntegers(5)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(inty)
	}
	log.Println("Server start...")
	//http.ListenAndServe(":8080", nil)
}
