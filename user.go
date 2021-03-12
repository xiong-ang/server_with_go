// @Title: user.go
// @Description: user struct
package main

import (
	"net"
)

// User object, send and receive message
type User struct {
	name     string
	conn     net.Conn
	sendC    chan string
	receiveC *chan string
}

// NewUser create user and start chat loop
func NewUser(conn net.Conn, receiveC *chan string) *User {
	user := &User{
		name:     conn.RemoteAddr().String(),
		conn:     conn,
		sendC:    make(chan string),
		receiveC: receiveC,
	}

	user.startChat()

	return user
}

// SendMsg send message to user
func (user *User) SendMsg(msg string) {
	user.sendC <- msg
}

func (user *User) startChat() {
	go func() {
		for {
			msg := <-user.sendC
			user.conn.Write([]byte(msg))
		}
	}()

	go func() {
		for {
			tmp := make([]byte, 256)
			user.conn.Read(tmp)
			*user.receiveC <- string(tmp)
		}
	}()
}
