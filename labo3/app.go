package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
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
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	scanner := bufio.NewScanner(bytes.NewReader(buf[:n]))
	scanner.Scan()
	buffer := scanner.Bytes()

	if string(buffer[0:3]) == "ACK" {
		var msg Message
		decoder := gob.NewDecoder(bytes.NewReader(buffer[3:]))
		decoder.Decode(&msg)
		fmt.Println(msg)
	} else {
		fmt.Println("nope")
	}
}
