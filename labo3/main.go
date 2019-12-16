/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : main.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        :  Entry point for the lab3: Implementation of bully algorithm
 -----------------------------------------------------------------------------------
*/

package main

import (
	"flag"
	"log"
	"prr-labo3/labo3/processus"
	"strconv"
)

func main() {
	id,N, aptitude := argValue()
	process := processus.Processus{}
	process.Init(id,N, aptitude)
}

func argValue() (uint16, uint16, uint16) {
	var proc string
	var procN string
	var apt string
	flag.StringVar(&proc, "proc", "", "Usage")
	flag.StringVar(&procN, "N", "", "Usage")
	flag.StringVar(&apt, "apt", "", "Usage")

	flag.Parse()
	id,err :=strconv.Atoi(proc)
	N,err :=strconv.Atoi(procN)
	aptitude,err :=strconv.Atoi(apt)

	if err != nil {
		log.Fatal("Client: Please put a number !")
	}

	return uint16(id),uint16(N), uint16(aptitude)
}
