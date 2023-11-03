package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = ":8080"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	listener, err := net.Listen(SERVER_TYPE, SERVER_PORT)

	if(err != nil) {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, connError := listener.Accept()

		if(connError != nil) {
			fmt.Println("Error connecting:", connError.Error())
			os.Exit(1)
		}

		// get message, output
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("Message Received:", string(message))

		// user logs in (auth)
		// Database has users, conversations and messages
		// when user logs in load their conversations
		// conversation has a list of user ids and messages
		// 
	}
}
