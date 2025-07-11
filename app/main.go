package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/codecrafters-io/redis-starter-go/app/parser"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handClient(conn)
	}
}

func handClient(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		if n > 0 {
			fmt.Println(string(buf))
			cmd, err := parser.Parse(buf)
			if err != nil {
				fmt.Print(err)
				return
			}
			result, err := command.CommandHandler(cmd)
			fmt.Println("result:")
			fmt.Println(string(result))
			if err != nil {
				fmt.Print(err)
			} else {
				conn.Write(result)
			}
		}
	}
}
