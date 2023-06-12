package main

import (
	"fmt"
	"net"
	"os"
)

const(
	SERVER_HOST="localhost"
	SERVER_PORT="9988"
	SERVER_TYPE="tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening...")
		os.Exit(1)
	}

	defer server.Close()

	fmt.Println("Listening on port ", SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepted...")
			os.Exit(1)
		}
		go prosesClient(connection)
	}
}

func prosesClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error Read...")
	} 
	fmt.Println("Received: ", string(buffer[:mLen]))
}