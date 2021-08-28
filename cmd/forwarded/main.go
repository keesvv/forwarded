package main

import (
	"fmt"
	"net"
	"os"

	"github.com/keesvv/forwarded/internal/cli"
	"github.com/keesvv/forwarded/pkg/proxy"
)

func handleConn(lconn *net.TCPConn, opts *cli.ConnOptions) {
	// Dial the remote host
	rconn, err := net.DialTCP("tcp", nil, opts.RemoteAddr)

	if err != nil {
		panic(err)
	}

	// Proxy packets between both connections
	pxy := proxy.NewProxy(lconn, rconn)
	pxy.Start()
}

func main() {
	if len(os.Args) < 3 {
		cli.PrintUsage()
	}

	opts := cli.ParseConnOptions()

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

	// TODO: move to own package
	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			panic(err)
		}

		fmt.Printf("New connection: %s\n", conn.RemoteAddr().String())
		go handleConn(conn, opts)
	}
}
