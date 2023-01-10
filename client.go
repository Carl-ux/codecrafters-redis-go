package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	//send message
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedInput))
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading from server", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", buf)
	}
}
