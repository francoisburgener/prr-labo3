/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : utils.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : Utility methods for conversions of the network layer
 -----------------------------------------------------------------------------------
*/
package utils

import (
	"bytes"
	"encoding/gob"
	"prr-labo3/labo3/config"
	"prr-labo3/labo3/network/messages"
	"strconv"
)

/**
 * Method to get the adress:port of the processus by id
 * @param id of the processus
 * @return address:port in string
 */
func AddressByID(id uint16) string{
	port := config.PORT + id
	return config.ADDR + ":" + strconv.Itoa(int(port))
}


/**
 * Method to initialize a message
 * @param _type is the type of the message (ACK,ECH,NOT,RES)
 * @param msg is our message
 * @return the concatanation of _type + msg
 */
func InitMessage(_type []byte, msg []byte) []byte{
	return append(_type,msg...)
}

/**
 * Method to encode a Message of type ACK or ECH
 * @param msg our message
 * @return our message in byte array
 */
func EncodeMessage(msg messages.Message) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

/**
 * Method to encode a Message of type RES
 * @param msg our MessageResult
 * @return our MessageResult in byte array
 */
func EncodeMessageResult(msg messages.MessageResult) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

/**
 * Method to encode a Message of type NOT
 * @param msg our MessageNotif
 * @return our MessageNotif in byte array
 */
func EncodeMessageNotif(msg messages.MessageNotif) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

/**
 * Method to decode a byte array to Message
 * @param buf who containe our message
 * @return Message
 */
func DecodeMessage(buf []byte)  messages.Message{
	var msg messages.Message
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg

}

/**
 * Method to decode a byte array to MessageResult
 * @param buf who containe our message
 * @return MessageResult
 */
func DecodeMessageResult(buf []byte)  messages.MessageResult{
	var msg messages.MessageResult
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg
}

/**
 * Method to decode a byte array to MessageNotif
 * @param buf who containe our message
 * @return MessageNotif
 */
func DecodeMessageNotif(buf []byte) messages.MessageNotif{
	var msg messages.MessageNotif
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg
}

