package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
// var _ = net.Listen
// var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 4096) // 4kb buffer
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Conn read error:", err)
		return
	}

	rawRequest := string(buf[:n])

	lines := strings.Split(rawRequest, "\r\n")

	if len(lines) > 0 {
		requestLine := lines[0]
		parts := strings.Split(requestLine, " ")
		if len(parts) >= 2 {
			method := parts[0]
			path := parts[1]

			if strings.ToLower(method) == "get" && path == "/" {
				okResponse := "HTTP/1.1 200 OK\r\n\r\n"
				_, err = conn.Write([]byte(okResponse))
				if err != nil {
					fmt.Println("Error writing OK response:", err.Error())
				}
				return
			}
		}
	}

	notFoundResponse := "HTTP/1.1 404 Not Found\r\n\r\n"
	_, err = conn.Write([]byte(notFoundResponse))

	if err != nil {
		fmt.Println("Error writing Not Found response:", err.Error())
	}
}
