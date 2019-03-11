package server

import (
	"errors"
	"log"
	"net"
	"strings"
)

type Server interface {
	Run() error
	Close() error
}

func NewServer(protocol, addr string, dh DataHandler) (Server, error) {
	switch strings.ToLower(protocol) {
	case "tcp":
		return &TCPServer{
			addr:        addr,
			dataHandler: dh,
		}, nil
	case "udp":
	}
	return nil, errors.New("invalid protocol given")
}

type TCPServer struct {
	addr        string
	server      net.Listener
	dataHandler DataHandler
}

func (t *TCPServer) Run() error {
	var err error
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := t.server.Accept()
			if err != nil || conn == nil {
				log.Printf("server accept error: %s", err)
				continue
			}

			go NewConnWrapper(conn, t.dataHandler).Handle()
		}
	}()
	return nil
}

func (t *TCPServer) Close() error {
	return t.server.Close()
}
