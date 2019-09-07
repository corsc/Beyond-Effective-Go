package _3_naive

import (
	"fmt"
	"net"
)

func Send(conn net.Conn, data interface{}) error {
	// convert data to bytes
	payload := []byte(fmt.Sprintf("%s", data))

	_, err := conn.Write(payload)
	return err
}
