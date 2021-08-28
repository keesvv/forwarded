package proxy

import (
	"fmt"
	"io"
	"net"

	"github.com/keesvv/forwarded/pkg/util"
)

type Proxy struct {
	LocalConn  *net.TCPConn
	RemoteConn *net.TCPConn
}

func (p *Proxy) proxyPackets(conn1 *net.TCPConn, conn2 *net.TCPConn) {
	for {
		b, err := util.ReadBytesBuf(conn1, 1024)

		if err != nil {
			if err == io.EOF {
				fmt.Printf(
					"Client %s disconnected\n",
					conn1.RemoteAddr().String(),
				)

				return
			}

			panic(err)
		}

		_, err = conn2.Write(b)

		if err != nil {
			panic(err)
		}
	}
}

func (p *Proxy) Start() {
	go p.proxyPackets(p.LocalConn, p.RemoteConn)
	go p.proxyPackets(p.RemoteConn, p.LocalConn)
}

func NewProxy(lconn *net.TCPConn, rconn *net.TCPConn) *Proxy {
	return &Proxy{
		LocalConn:  lconn,
		RemoteConn: rconn,
	}
}
