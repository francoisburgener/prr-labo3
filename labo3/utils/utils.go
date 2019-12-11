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

func InitMessage(_type []byte, msg []byte) []byte{
	return append(_type,msg...)
}

func EncodeMessage(msg messages.Message) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

func EncodeMessageResult(msg messages.MessageResult) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

func EncodeMessageNotif(msg messages.MessageNotif) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)

	return buf.Bytes()
}

func DecodeMessage(buf []byte)  messages.Message{
	var msg messages.Message
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg

}

func DecodeMessageResult(buf []byte)  messages.MessageResult{
	var msg messages.MessageResult
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg
}

func DecodeMessageNotif(buf []byte) messages.MessageNotif{
	var msg messages.MessageNotif
	decoder := gob.NewDecoder(bytes.NewReader(buf))
	decoder.Decode(&msg)
	return msg
}

