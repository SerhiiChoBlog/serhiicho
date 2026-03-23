package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"serhii/internal/config"
	"serhii/internal/database"
	"serhii/internal/http"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := godotenv.Load(); err != nil {
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

	serv := http.NewServer(conf, database.NewDatabase(db))
	if err := serv.ListenAndServe(os.Getenv("APP_PORT")); err != nil {
		fmt.Println(err)
	}
}

func dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}

func setupSql() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn())
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
