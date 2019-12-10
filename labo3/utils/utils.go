package utils

import (
	"encoding/binary"
	"prr-labo3/labo3/config"
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
 * Method to convert an unint32 in byte array
 * @param value we want to update
 * @return the value converted in byte array
 */
func uint32ToByteArray(i uint32) []byte{
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}

/**
 * Method to convert an unint16 in byte array
 * @param value we want to update
 * @return the value converted in byte array
 */
func uint16ToByteArray(i uint16) []byte{
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	return buf
}

/**
 * Method to convert a byte array to unint32
 * @param buf we want to convert
 * @return the value in uint32
 */
func ConverByteArrayToUint32(buf []byte) uint32{
	return binary.LittleEndian.Uint32(buf)
}

/**
 * Method to convert a byte array to unint16
 * @param buf we want to convert
 * @return the value in uint16
 */
func ConverByteArrayToUint16(buf []byte) uint16{
	return binary.LittleEndian.Uint16(buf)
}