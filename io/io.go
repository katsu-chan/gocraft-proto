package io

import (
	"bufio"
	"encoding/binary"
	"net"
)

// func readUShort(buf []byte) uint16 {
// 	short_big := buf[0]
// 	short_little := buf[1]

// 	short := uint16(short_big)*0x100 + uint16(short_little)

// 	buf = buf[2:]

// 	return short
// }

func Init() (rw bufio.ReadWriter) {
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	rw = *bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	return
}

func ReadString(r *bufio.Reader) string {
	len := ReadByte(r)
	buf := make([]byte, len)
	r.Read(buf)
	return string(buf)
}

func ReadUVarInt(r *bufio.Reader) uint64 {
	a, _ := binary.ReadUvarint(r)
	return a
}

func ReadUShort(r *bufio.Reader) uint16 {
	buf := make([]byte, 2)
	r.Read(buf)
	return binary.BigEndian.Uint16(buf)
}

func ReadByte(r *bufio.Reader) byte {
	a, _ := r.ReadByte()
	return a
}

func WriteString(str string, w bufio.Writer) {
	w.WriteByte(byte(len(str)))
	w.WriteString(str)
}

func WriteUVarInt(a uint64, w bufio.Writer) {
	buf := make([]byte, 64)
	w.Write(buf[:binary.PutUvarint(buf, a)])
	//binary.PutUvarint(buf, a)
}

func WriteUShort(a uint16, w bufio.Writer) {
	binary.Write(&w, binary.BigEndian, a)
}

func WriteByte(a byte, w bufio.Writer) {
	w.WriteByte(a)
}
