package main

import (
	"fmt"
	"net/http"
)

func main() {
	route := http.NewServeMux()
	server := &http.Server{Handler: route, Addr: ":8081"}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Print("Error: ", err)
	}
}
