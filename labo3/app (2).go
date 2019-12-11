package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
)

type Message struct {
	Name   string
	Level  int
	Potato bool
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9001")
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var buffer bytes.Buffer
	msg := Message{"jack", 17, true}
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(msg)

	conn.Write(append([]byte("ACP"), buffer.Bytes()...))
}
