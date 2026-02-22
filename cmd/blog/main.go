package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"os"
	"serhii/internal/config"
	"serhii/internal/database"
	"serhii/internal/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	conf, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := setupSql()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	serv := http.NewServer(conf, database.NewMySql(db))
	if err := serv.ListenAndServe(os.Getenv("APP_PORT")); err != nil {
		fmt.Println(err)
	}
}

func dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}

func setupSql() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		return nil, fmt.Errorf("Error %s when opening DB", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Errors %s pinging DB", err)
	}

	fmt.Println("Connected to DB successfully")

	return db, nil
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
