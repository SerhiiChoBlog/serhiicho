package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"serhii/internal/config"
	"serhii/internal/http"
	"strings"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	conf, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	serv := http.NewServer(conf)
	if err := serv.ListenAndServe(os.Getenv("APP_PORT")); err != nil {
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

func loadConfig() (*config.Config, error) {
	f, err := os.Open("./config/config.json")
	if err != nil {
		return nil, fmt.Errorf("Cannot open config.json file: %v", err)
	}

	conf := new(config.Config)
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(conf); err != nil {
		return nil, fmt.Errorf("Json Encoder error: %v", err)
	}

	return conf, nil
}
