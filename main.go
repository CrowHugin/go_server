package main

import (
	"fmt"
	"net/http"
)

func main() {
	route := http.NewServeMux()
	route.Handle("/", http.FileServer(http.Dir(".")))
	server := &http.Server{Handler: route, Addr: ":8081"}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Print("Error: ", err)
	}
}
