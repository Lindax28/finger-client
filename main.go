package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Verify number of arguments is correct
	if len(os.Args) != 2 {
		fmt.Println("Proper Usage is: ./finger user@hostname")
		os.Exit(1)
	}
	// Verify format is username + '@' + hostname
	args := strings.Split(os.Args[1], "@")
	if len(args) != 2 {
		fmt.Println("Proper Usage is: ./finger user@hostname")
		os.Exit(1)
	}
	user := args[0]
	hostname := args[1]
	// Verify username is in ASCII format
	for _, char := range user {
		if char > unicode.MaxASCII {
			fmt.Println("Username must be in ASCII format")
			os.Exit(1)
		}
	}
	// Connect to host on port 79
	conn, err := net.Dial("tcp", hostname+":79")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	// Write user data to the connection, ended with CRLF
	conn.Write([]byte(user + "\r\n"))
	// Read data from the connection
	response := make([]byte, 0)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			break
		}
		response = append(response, tmp[:n]...)
	}
	fmt.Println(string(response))
}
