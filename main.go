package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type ConnOptions struct {
	LocalAddr  *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func readBytes(conn *net.TCPConn) ([]byte, error) {
	b := make([]byte, 1024)
	_, err := conn.Read(b)

	if err != nil {
		return nil, err
	}

	return bytes.TrimRight(b, "\x00"), nil
}

func proxyPackets(lconn *net.TCPConn, rconn *net.TCPConn) {
	for {
		b, err := readBytes(lconn)

		if err != nil {
			if err == io.EOF {
				fmt.Printf(
					"Client %s disconnected\n",
					lconn.RemoteAddr().String(),
				)

				return
			}

			panic(err)
		}

		_, err = rconn.Write(b)

		if err != nil {
			panic(err)
		}
	}
}

func handleConn(lconn *net.TCPConn, opts *ConnOptions) {
	// Dial the remote host
	rconn, err := net.DialTCP("tcp", nil, opts.RemoteAddr)

	if err != nil {
		panic(err)
	}

	// Proxy packets between both connections
	go proxyPackets(lconn, rconn)
	go proxyPackets(rconn, lconn)
}

func parseConnOptions() *ConnOptions {
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

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("%s <host> <lport>:<rport>", os.Args[0])
	}

	opts := parseConnOptions()

	fmt.Println(
		"Forwarding",
		opts.LocalAddr.String(),
		"->",
		opts.RemoteAddr.String(),
	)

	listener, err := net.ListenTCP("tcp", opts.LocalAddr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for connections...")

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			panic(err)
		}

		fmt.Printf("New connection: %s\n", conn.RemoteAddr().String())
		go handleConn(conn, opts)
	}
}
