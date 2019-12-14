package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"prr-labo3/labo3/network"
	"strconv"
)

func main() {
	id,N := argValue()
	n := network.Network{}
	n.Init(id,N)
	console(&n)
}

func argValue() (uint16, uint16) {
	var proc string
	var procN string
	flag.StringVar(&proc, "proc", "", "Usage")
	flag.StringVar(&procN, "N", "", "Usage")
	flag.Parse()
	id,err :=strconv.Atoi(proc)
	N,err :=strconv.Atoi(procN)
	if err != nil {
		log.Fatal("Client: Please put a number !")
	}

	return uint16(id),uint16(N)

}

func console(n *network.Network) {

	log.Println("Client: Choice (number)")
	log.Println("Client: ---------------------")

	scanner := bufio.NewScanner(os.Stdin)

	m := map[uint16]uint16{}
	m[0] = 18
	m[1] = 34
	m[2] = 52
	m[3] = 11

	m2 := map[uint16]bool{}
	m2[0] = true
	m2[2] = true
	m2[3] = true
	m2[4] = true

	for{
		fmt.Println("1 - Notif")
		fmt.Println("2 - Resul")
		fmt.Println("1 - Echo")
		fmt.Print("> ")

		scanner.Scan()
		choice :=  scanner.Text()

		switch choice {
		case "1":
			n.EmitNotif(m)
		case "2":
			n.EmitResult(m2)
		case "3":
			n.EmitEcho()
		default:
			fmt.Println("Choose 1 or 2")
		}
	}

}
