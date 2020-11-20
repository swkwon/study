package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "server:9000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn.Write([]byte("hello server."))
		time.Sleep(1 * time.Second)
	}
}
