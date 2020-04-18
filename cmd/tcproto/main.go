package main

import (
	"github.com/macieklad/tcproto/pkg/proto"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Printf("%v", err)
	}

	hub := proto.NewHub()
	go hub.Run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
		}

		c := hub.MakeClient(conn)
		go c.Read()
	}
}
