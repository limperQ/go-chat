package handlers

import (
	"log"
	"net/http"
	"shev-chat/structs"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var users = make(map[structs.ConnectUser]int)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)

	defer func() {
		if err := ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	log.Println("Client connected:", ws.RemoteAddr().String())

	var socketClient *structs.ConnectUser = structs.NewConnectedUser(ws, ws.RemoteAddr().String())
	users[*socketClient] = 0
	log.Println("Number client connected ...", len(users))

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Ws disconnected waiting", err.Error())
			delete(users, *socketClient)
			log.Println("Number of client still connected ...", len(users))
			return
		}

		for client := range users {
			if err = client.Websocket.WriteMessage(messageType, message); err != nil {
				log.Println("Cloud not send Message to ", client.ClientIP, err.Error())
			}
		}
	}
}
