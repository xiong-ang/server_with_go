// @Title: userGroup.go
// @Description: used to group user
package main

import (
	"net"
)

// UserGroup is used to group user and manager user message
type UserGroup struct {
	users    map[string]*User
	receiveC chan string
}

// BuildUserGroup is used to create UserGroup object and startReceiveMsg
func BuildUserGroup() *UserGroup {
	ret := &UserGroup{
		users:    make(map[string]*User),
		receiveC: make(chan string),
	}

	ret.startReceiveMsg()
	return ret
}

// HandlerConnection is used to handle tcp connection
func (userGroup *UserGroup) HandlerConnection(conn net.Conn) {
	user := NewUser(conn, &userGroup.receiveC)
	userGroup.addUser(conn.RemoteAddr().String(), user)
}

func (userGroup *UserGroup) startReceiveMsg() {
	go func() {
		for input := range userGroup.receiveC {
			userGroup.handleMsg("", input)
		}
	}()
}

func (userGroup *UserGroup) addUser(name string, user *User) {
	userGroup.users[name] = user
}

func (userGroup *UserGroup) handleMsg(name string, msg string) {
	for key, user := range userGroup.users {
		if key != name {
			user.SendMsg(msg)
		}
	}
}
