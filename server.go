// @Title: server.go
// @Description: tcp server object
package main

import (
	"fmt"
	"net"
)

// Server struct
type Server struct {
	ip   string
	port uint
}

// StartServer is used to create server object and start server listening
func StartServer(ip string, port uint) *Server {
	server := &Server{
		ip:   ip,
		port: port,
	}
	server.start()

	return server
}

func (server *Server) start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.ip, server.port))
	if err != nil {
		fmt.Println("Listen error: " + err.Error() + "\n")
		return
	}
	defer listener.Close()

	userGroup := BuildUserGroup()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v\n", err)
			return
		}

		userGroup.HandlerConnection(conn)
	}
}
