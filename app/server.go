package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//stage1 bind to a port
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept() //store the connection value
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
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
		//stage2 respond to ping cmd
		if _, err := conn.Write([]byte("+PONG\r\n")); err != nil {
			fmt.Println("PONG Failed")
			os.Exit(1)
		}
	}
}
