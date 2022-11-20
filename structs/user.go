package structs

import "github.com/gorilla/websocket"

type ConnectUser struct {
	Websocket *websocket.Conn
	ClientIP  string
	UserInfo  UserInfo
}

type UserInfo struct {
	UserName string
	Message  string
}

func NewConnectedUser(ws *websocket.Conn, clientIP string, userInfo UserInfo) *ConnectUser {
	return &ConnectUser{
		Websocket: ws,
		ClientIP:  clientIP,
		UserInfo:  userInfo,
	}
}
