/*
 -----------------------------------------------------------------------------------
 Lab 		 : 01
 File    	 : network.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : ...
 -----------------------------------------------------------------------------------
*/

package network

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"prr-labo3/labo3/config"
	"prr-labo3/labo3/network/messages"
	"prr-labo3/labo3/utils"
	"time"
)

type Network struct {
	id uint16
	N  uint16
}

func (n *Network) Init(id uint16, N uint16) {
	n.id = id
	n.N = N

	go func() {
		n.initServ()
	}()
}

func (n *Network) initServ() {
	addr := utils.AddressByID(n.id)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal("Network error: Initialisation failed",err)
	}
	defer conn.Close()

	n.handleConn(conn)
}

func (n *Network) handleConn(conn net.PacketConn) {
	buf := make([]byte, 1024)
	for {
		l, cliAddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal("Network error: Reading error ",err)
		}
		s := bufio.NewScanner(bytes.NewReader(buf[0:l]))
		for s.Scan() {
			buf := s.Bytes()
			n.emitACK(conn,cliAddr)
			n.decodeMessage(buf)
		}
	}
}

func (n *Network) EmitNotif(_map map[uint16]uint16){
	notif := messages.MessageNotif{_map}
	msg := utils.EncodeMessageNotif(notif)
	buf := utils.InitMessage([]byte(config.NotifMessage),msg)
	n.emit(buf)
}

func (n *Network) EmitResult(_map map[uint16]bool){
	fmt.Println(_map)
	result := messages.MessageResult{_map}
	msg := utils.EncodeMessageResult(result)
	buf := utils.InitMessage([]byte(config.ResultMessage),msg)
	n.emit(buf)
}

func (n *Network) emitACK(conn net.PacketConn, cliAddr net.Addr) {
	ack := messages.Message{n.id}
	msg := utils.EncodeMessage(ack)
	buf := utils.InitMessage([]byte(config.AckMessage),msg)

	if _, err := conn.WriteTo(buf, cliAddr); err != nil {
		log.Fatal("Network error: Writing error ",err)
	}
}

func (n *Network) EmitEcho() {
	echo := messages.Message{n.id}
	msg := utils.EncodeMessage(echo)
	buf := utils.InitMessage([]byte(config.EchoMessage),msg)
	n.emit(buf)
}

func (n *Network) emit(msg []byte) {

	for i:= n.id; i < n.N + n.id; i++{
		id := (i + 1) % n.N
		if id != n.id{
			addr := utils.AddressByID(id)
			conn,err := net.Dial("udp",addr)
			if err != nil {
				log.Printf("The processus %d is not alive ",id)
			}

			_, err = conn.Write(msg)
			if err != nil {
				log.Fatal("Network error: Writing error ",err)
			}

			n.readACK(conn)
		}
	}
}

func (n *Network) readACK(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	//Set the deadline
	err := conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	if err != nil{
		log.Println("Timeout",err)
		return
	}

	// Read the incoming connection into the buffer.
	l, _ := conn.Read(buf)
	/*if err != nil {
		log.Println("Network error: Error reading", err.Error())
	}*/

	s := bufio.NewScanner(bytes.NewReader(buf[0:l]))

	for s.Scan(){
		buf := s.Bytes()
		if string(buf[0:3]) == config.AckMessage{
			msg := utils.DecodeMessage(buf[3:])
			log.Println("Decode : ",string(buf[0:3]),"-",msg.Id)
		}
	}

}

func (n *Network) decodeMessage(buf []byte) {

	_type := string(buf[0:3])

	switch _type {
	case config.EchoMessage:
		msg := utils.DecodeMessage(buf[3:])
		log.Println("Decode",_type,"-",msg.Id)
	case config.ResultMessage:
		msg := utils.DecodeMessageResult(buf[3:])
		log.Println("Decode",_type,"-",msg.Map)
	case config.NotifMessage:
		msg := utils.DecodeMessageNotif(buf[3:])
		log.Println("Decode",_type,"-",msg.Map)
	default:
		log.Println("Network: Incorrect type message !")
	}
	
}
