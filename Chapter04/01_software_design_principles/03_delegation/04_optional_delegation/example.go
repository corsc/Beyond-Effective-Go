package _4_optional_delegation

import (
	"fmt"
	"net"
)

func Send(conn net.Conn, data interface{}) error {
	// convert data to bytes using optional delegation
	var payload []byte
	if encoder, ok := data.(ByteEncoder); ok {
		payload = encoder.Encode()
	} else {
		payload = []byte(fmt.Sprintf("%s", data))
	}

	_, err := conn.Write(payload)
	return err
}

type ByteEncoder interface {
	Encode() []byte
}
