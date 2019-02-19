package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func connect(user, pw, host, db string, port int) {
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pw, host, port, db))
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %s\n", err.Error())
	}

	if err = database.Ping(); os.IsExist(err) {
		log.Fatalf("Couldn't establish the connection to the database: %s\n", err.Error())
	}

	log.Println("Successfully connected to the database.")
}

func disconnect() {
	err := database.Close()
	if err != nil {
		log.Printf("Couldn't disconnect from the database: %s\nContinuing...\n", err.Error())
	}
}
