package controller

import (
	"chat/server"
	"chat/global"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func WsPage(res http.ResponseWriter, req *http.Request) {
	// 将http请求升级为webSocket请求
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	// 将获得的webSocket构建出一个client并进行注册
	client := &server.Client{Id: uuid.New().String(), Socket: conn, Send: make(chan []byte)}
	global.Manager.Register <- client

	// 启用协程调用client的Read和Write
	go client.Read(&global.Manager)
	go client.Write()
}
