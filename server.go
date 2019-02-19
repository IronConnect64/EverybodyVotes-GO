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
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("EverybodyVotes-GO - 1.2.0")

	// We need some additional data for the server.
	log.Println("Loading configuration...")
	config := load()

	// And here's the second part.
	connect(config.Database.Username, config.Database.Password, config.Database.Hostname, config.Database.Database, config.Database.Port)

	// Let's make an gin engine for our server.
	log.Println("Initializing server...")
	server := gin.Default()

	server.GET("/polls", pollsHandler)
	server.POST("/register", registerHander)
	server.POST("/unregister", unregisterHandler)
	server.POST("/vote", voteHandler)

	// If possible, we can replace this with RunTLS() in the future.
	log.Println("Starting server...")
	server.Run(":" + config.Port)

	// This gets executed after the server has been shutdown
	defer disconnect()
	log.Fatalln("Caught signal, shutting down...")
}

func checkToken(creds auth) bool {
	result, err := database.Query(fmt.Sprintf("select Token from Users where MAC='%s';", creds.ID))
	if os.IsExist(err) {
		log.Printf("[1/2] An error occurred when checking the Token. ID: %s - Token: %s.\nError: %s", creds.ID, creds.Token, err.Error())
		return false
	}

	column, err := result.Columns()
	if os.IsExist(err) {
		log.Printf("[2/2] An error occurred when checking the Token for the ID %s. Token: %s.\nError: %s", creds.ID, creds.Token, err.Error())
		return false
	}

	for _, v := range column {
		if v == creds.Token {
			log.Printf("ID %s used Token %s to log in and succeeded.", creds.ID, creds.Token)
			return true
		}
	}

	log.Printf("ID %s used Token %s to log in, but failed.", creds.ID, creds.Token)
	return false
}

func pollsHandler(ctx *gin.Context) {
	if len(ctx.GetHeader("HTTP_X_PSP_BROWSER")) == 0 {
		ctx.AbortWithStatus(http.StatusUnavailableForLegalReasons)
	}

	var credentials auth

	if err := json.NewDecoder(ctx.Request.Body).Decode(&credentials); os.IsExist(err) {
		log.Printf("Couldn't identify user.")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	} else if !checkToken(credentials) {
		log.Printf("User with ID %s tried to log in with invalid credentials.", credentials.ID)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	result, err := database.Query("select * from polls;")
	if os.IsExist(err) {
		log.Printf("[1/2] User with MAC %s tried to fetch data and failed.", ctx.GetHeader("MAC"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	stuff, err := result.Columns()
	if os.IsExist(err) {
		log.Printf("[2/2] User with MAC %s tried to fetch data and failed.", ctx.GetHeader("MAC"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	latest, _ := strconv.ParseBool(stuff[4])

	if err := json.NewEncoder(ctx.Writer).Encode(&poll{
		ID:       stuff[0],
		Question: stuff[1],
		A:        stuff[2],
		B:        stuff[3],
		Latest:   latest,
	}); os.IsExist(err) {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Printf("User with MAC %s succeeded to fetch the latest poll.", ctx.GetHeader("MAC"))
}

func registerHander(ctx *gin.Context) {

}

func unregisterHandler(ctx *gin.Context) {

}

func voteHandler(ctx *gin.Context) {

}
