package main

import (
	"fmt"
	"serhii-go/internal/http"
)

func main() {
	serv := http.NewServer(":8080")

	if err := serv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
