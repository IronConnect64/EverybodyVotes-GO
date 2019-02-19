//		Copyright (c) 2019 IronConnect64
//
//		EverybodyVotes-GO, a custom server for the Everybody Votes channel on the PSP.
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE.

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
