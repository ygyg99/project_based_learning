package server

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// 客户端, 用来存储每个连接的信息
type Client struct {
	Id     string          // 用户id
	Socket *websocket.Conn // 用户的socket
	Send   chan []byte     // 用户发送消息的通道
}

// 定义客户端的读写操作
func (c *Client) Read(man *ClientManager) {
	defer func() {
		man.Unregister <- c
		c.Socket.Close()
	}()
	// 反复读取消息，
	for {
		_, mes, err := c.Socket.ReadMessage()
		// 如果读取消息失败则意味着客户端已经断开连接，需要注销
		if err != nil {
			man.Unregister <- c
			c.Socket.Close()
			break
		}
		// 将消息整合一下
		Mes := &Message{Sender: c.Id, Content: string(mes)}
		// 将消息转化为json，并将其发到广播端口
		jsonMes, _ := json.Marshal(Mes)
		man.Broadcast <- jsonMes
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()
	//在网络编程中需要持续处理请求/发送响应,常用for循环
	for {
		select {
		case mes, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, nil)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, mes)
		}
	}
}
