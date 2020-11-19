package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go connHandler(c)
	}
}

func connHandler(conn net.Conn) {
	recvBuf := make([]byte, 1024)
	for {
		n, err := conn.Read(recvBuf)
		if err != nil {
			if err == io.EOF {
				log.Println("connection is closed from client")
			} else {
				log.Println("fail to receive data. err:", err)
			}
			return
		}
		log.Println(string(recvBuf[:n]))
	}
}
