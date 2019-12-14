/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : messages.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : Our different struct message
 -----------------------------------------------------------------------------------
*/

package messages

type Message struct {
	Id uint16
}

type MessageResult struct {
	Id uint16
	Map map[uint16]bool
}

type MessageNotif struct {
	Map map[uint16]uint16
}
