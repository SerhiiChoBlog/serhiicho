package main

import (
	"fmt"
	"log"
	"serhii/internal/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	serv := http.NewServer()

	if err := serv.ListenAndServe(":8080"); err != nil {
		fmt.Println(err)
	}
}
