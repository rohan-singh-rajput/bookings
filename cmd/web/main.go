package main

import (
	"fmt"
	"net/http"

	"github.com/rohan-singh-rajput/bookings/pkg/handlers"
)

const PortNumber = ":8080"

func main() {

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("listening at port", PortNumber)

	err := http.ListenAndServe(PortNumber, nil)

	if err != nil {
		fmt.Println("Error connecting to the server")
	}

}
