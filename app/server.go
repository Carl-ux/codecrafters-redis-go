package main

import (
	"fmt"
	//Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept() //store the connection value
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 1024) //slice
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("error reading from cilent: ", err.Error())
		os.Exit(1)
	}

	// Let's ignore the client's input for now and hardcode a response.
	// We'll implement a proper Redis Protocol parser in later stages.
	pingHandler(conn)
}

func pingHandler(conn net.Conn) {
	if _, err := conn.Write([]byte("+PONG\r\n")); err != nil {
		fmt.Println("PONG Failed")
		os.Exit(1)
	}
}
