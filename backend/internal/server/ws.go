package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
    ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} 

func (* Server) WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err:= upgrader.Upgrade(w,r,nil)
	if err!=nil{
		fmt.Println("Websocket Update error: ",err)
		return
	}
	defer conn.Close()
	for {
		messageType,msg,err:=conn.ReadMessage()
		if err!=nil{
			fmt.Println("Read error: ",err)
			break
		}
		fmt.Printf("Received %s\n",msg)
		if err:=conn.WriteMessage(messageType, msg); err!=nil{
			fmt.Println("Write error:",err)
			break
		}
	}
}

