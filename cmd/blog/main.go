package main

import (
	"fmt"
	"serhii/internal/http"
)

func main() {
	serv := http.NewServer()

	if err := serv.ListenAndServe(":8080"); err != nil {
		fmt.Println(err)
	}
}
