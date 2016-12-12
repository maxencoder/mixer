package proxy

import (
	"net"
	"runtime"
	"strings"

	"github.com/maxencoder/log"
	"github.com/maxencoder/mixer/config"
)

type Server struct {
	cfg *config.Config

	addr     string
	user     string
	password string

	running bool

	listener net.Listener

	//nodes map[string]*node.Node

	schemas map[string]*Schema
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password

	var err error
	netProto := "tcp"
	if strings.Contains(netProto, "/") {
		netProto = "unix"
	}
	s.listener, err = net.Listen(netProto, s.addr)

	if err != nil {
		return nil, err
	}

	log.Info("Server run MySql Protocol Listen(%s) at [%s]", netProto, s.addr)
	return s, nil
}

func (s *Server) Run() error {
	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Error("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	var conn *Conn
	var err error

	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Error("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	conn, err = s.newConn(c)

	if err != nil {
		log.Error("onConn error: %v", err)
		c.Close()
		return
	}

	conn.Run()
}
