package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Proper Usage is: ./finger user@hostname")
		os.Exit(1)
	}

	args := strings.Split(os.Args[1], "@")
	if len(args) != 2 {
		fmt.Println("Proper Usage is: ./finger user@hostname")
		os.Exit(1)
	}
	user := args[0]
	hostname := args[1]

	conn, err := net.Dial("tcp", hostname+":79")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write([]byte(user + "\r\n"))
	response := make([]byte, 4096)
	conn.Read(response)
	fmt.Println(string(response))
}
