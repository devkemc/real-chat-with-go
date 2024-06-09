package entity

import "golang.org/x/net/websocket"

type ActiveUser struct {
	Name string
	Conn *websocket.Conn
}

func NewActiveUser(name string, conn *websocket.Conn) *ActiveUser {
	return &ActiveUser{
		Name: name,
		Conn: conn,
	}
}
