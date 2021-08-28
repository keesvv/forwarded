package util

import (
	"bytes"
	"net"
)

func ReadBytesBuf(conn *net.TCPConn, bufSize int) ([]byte, error) {
	b := make([]byte, bufSize)
	_, err := conn.Read(b)

	if err != nil {
		return nil, err
	}

	return bytes.TrimRight(b, "\x00"), nil
}
