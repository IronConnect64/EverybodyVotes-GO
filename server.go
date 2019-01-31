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
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ErrorCheck makes error checking easier.
func ErrorCheck(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

func main() {
	log.Println("Everybody-Votes-GO - 1.0.0")

	log.Println("Loading configuration...")
	cfg := Config{}
	file, err := ioutil.ReadFile("config.json")
	ErrorCheck(err)
	ErrorCheck(json.Unmarshal([]byte(file), &cfg))

	log.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(cfg.Address, nil))

	// From there on, everything server-related will be redirected to our handler function.
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() == "PSP" {
		// Do PSP stuff
	} else if r.UserAgent() == "PlayStation Vita" {
		// Do Vita stuff
	} else {
		w.Write([]byte("Whatever you're trying, you're very wrong here."))
	}
}
