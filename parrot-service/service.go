package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jainsaur/parrot"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error in starting chat service - %v\n", err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error in accepting connection - %v\n", err)
			continue
		}
		go parrot.HandleConnection(conn)
	}
}
