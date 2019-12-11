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
	"io"
	"log"
	"net"
	"prr-labo3/labo3/config"
	"prr-labo3/labo3/utils"
	"prr-labo3/labo3/visit"
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
		log.Fatal(err)
	}
	defer conn.Close()

	go n.handleConn(conn)
}

func (n *Network) handleConn(conn net.PacketConn) {
	buf := make([]byte, 1024)
	for {
		l, cliAddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(bytes.NewReader(buf[0:l]))
		for s.Scan() {
			buf := s.Bytes()
			n.emitACK(conn,cliAddr)
			n.decodeMessage(buf)
		}
	}
}

func (n *Network) EmitNotif(array []visit.Visit){

}

func (n *Network) EmitResult(id uint16){
	msg := utils.InitMessage(id,[]byte(config.ResultMessage))
	n.emit(msg)
}

func (n *Network) emitACK(conn net.PacketConn, cliAddr net.Addr) {
	msg := utils.InitMessage(n.id,[]byte(config.AckMessage))
	if _, err := conn.WriteTo(msg, cliAddr); err != nil {
		log.Fatal(err)
	}
}

func (n *Network) EmitEcho() {
	msg := utils.InitMessage(n.id,[]byte(config.EchoMessage))
	n.emit(msg)
}

func (n *Network) emit(msg []byte) {

	for i:= n.id; i < n.N + n.id; i++{
		id := (n.id + 1) % n.N

		if id != n.id{
			addr := utils.AddressByID(id)
			conn,err := net.Dial("udp",addr)
			if err != nil {
				log.Printf("The processus %d is not alive",id)
			}

			mustCopy(conn,bytes.NewReader(msg))

			//Set the deadline
			err2 := conn.SetReadDeadline(time.Now().Add(time.Second * 2))
			if err2 != nil{
				log.Println("Timeout")
				continue
			}
			n.readACK(conn)
		}
	}
}

func (n *Network) readACK(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	l, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Network error: Error reading:", err.Error())
	}

	s := bufio.NewScanner(bytes.NewReader(buf[0:l]))

	for s.Scan(){
		n.decodeMessage(s.Bytes())
	}

}

func (n *Network) decodeMessage(buf []byte) {
	
}


func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}



func main() {

}
