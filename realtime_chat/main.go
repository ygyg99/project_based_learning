package main

import (
	"chat/controller"
	"chat/global"
	"chat/server"
	"fmt"
	"net/http"
)




func main() {
	// 定义一个全局的客户端管理器
	global.Manager = server.ClientManager{
		Clients:    make(map[*server.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *server.Client),
		Unregister: make(chan *server.Client),
	}
	fmt.Println("Start application...")
	go global.Manager.Start()
	http.HandleFunc("/ws",controller.WsPage)
	http.ListenAndServe(":12345",nil)
}