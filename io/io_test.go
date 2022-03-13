package io

import (
	"bufio"
	"bytes"
	"testing"
)

func TestRead(t *testing.T) {
	buf := []byte{0x00, 0x01, 0x03, 0x48, 0x69, 0x21, 0x20, 0xE8, 0x07, 0x1f, 0x91} //Byte(0), Byte(1), String(3), UVarInt(20), UVarInt(1000), UShort(8081))
	r := bufio.NewReader(bytes.NewBuffer(buf))

	if ReadByte(r) != 0 {
		t.Fatalf("1st byte != 0")
	}
	if ReadByte(r) != 1 {
		t.Fatalf("2nd byte != 1")
	}
	if ReadString(r) != "Hi!" {
		t.Fatalf("String != \"Hi!\"")
	}
	if ReadUVarInt(r) != 32 {
		t.Fatalf("1st UVarInt != 20")
	}
	if ReadUVarInt(r) != 1000 {
		t.Fatalf("2nd UVarInt != 1000")
	}
	if ReadUShort(r) != 8081 {
		t.Fatalf("UShort != 8081")
	}
}
