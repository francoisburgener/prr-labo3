/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : config.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : Config file for the network layer
 -----------------------------------------------------------------------------------
*/
package config

import "time"

const (
	ADDR = "127.0.0.1"
	PORT = 6000
	NotifMessage = "NOT"
	ResultMessage = "RES"
	AckMessage = "ACK"
	EchoMessage = "ECH"
	TIME_OUT = time.Second * 2

)
