package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

type DataHandler interface {
	Handle(string) error
}

type ConnWrapper struct {
	io.ReadWriteCloser
	log     *log.Logger
	handler DataHandler
}

func NewConnWrapper(conn net.Conn, handler DataHandler) *ConnWrapper {
	return &ConnWrapper{
		conn,
		log.New(os.Stdout, conn.RemoteAddr().String(), log.LstdFlags),
		handler,
	}
}

func (t ConnWrapper) Handle() {
	defer t.Close()

	var (
		conn = bufio.NewReadWriter(bufio.NewReader(t), bufio.NewWriter(t))
		data string
		err  error
	)

	for {
		if data, err = conn.ReadString('\n'); err != nil {
			t.log.Printf("read bytes error: %s", err)
		}
		err = t.handler.Handle(data)

	}
}
