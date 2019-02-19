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
	"log"

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

	server.POST("/login", loginHandler)
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

func loginHandler(ctx *gin.Context) {

}

func pollsHandler(ctx *gin.Context) {

}

func registerHander(ctx *gin.Context) {

}

func unregisterHandler(ctx *gin.Context) {

}

func voteHandler(ctx *gin.Context) {

}
