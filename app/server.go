package main

import (
	"fmt"
	"io"
	//Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Redis go build 3rd stage")

	// Uncomment this block to pass the first stage
	//stage1 bind to a port
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
	//stage3 forloop
	for {
		buf := make([]byte, 1024) //slice
		_, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading from cilent: ", err.Error())
			os.Exit(1)
		}

		// Let's ignore the client's input for now and hardcode a response.
		// We'll implement a proper Redis Protocol parser in later stages.
		//stage2 respond to ping cmd
		pingHandler(conn)
	}
}

func pingHandler(conn net.Conn) {
	if _, err := conn.Write([]byte("+PONG\r\n")); err != nil {
		fmt.Println("PONG Failed")
		os.Exit(1)
	}
}
