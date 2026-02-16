package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"serhii/internal/http"
	"strings"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	serv := http.NewServer()

	if err := serv.ListenAndServe(":8080"); err != nil {
		fmt.Println(err)
	}
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}
