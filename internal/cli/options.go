package cli

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type ConnOptions struct {
	LocalAddr  *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func ParseConnOptions() *ConnOptions {
	lhost := "localhost"
	rhost := os.Args[1]

	ports := strings.Split(os.Args[2], ":")
	lport, _ := strconv.ParseUint(ports[0], 10, 16)
	rport, _ := strconv.ParseUint(ports[1], 10, 16)

	ltcpAddr, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s:%d", lhost, lport),
	)

	if err != nil {
		panic(err)
	}

	rtcpAddr, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s:%d", rhost, rport),
	)

	if err != nil {
		panic(err)
	}

	return &ConnOptions{
		LocalAddr:  ltcpAddr,
		RemoteAddr: rtcpAddr,
	}
}
