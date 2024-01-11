package server

import (
	"encoding/json"
	"fmt"
)

// 先定义聊天应用的三个结构体

// 客户端管理器, 用来管理所有的客户端连接, 并且实现广播消息的功能
type ClientManager struct {
	Clients    map[*Client]bool // 用来存储所有的连接
	Broadcast  chan []byte      // 用来广播消息的通道
	Register   chan *Client     // 用来注册连接的通道
	Unregister chan *Client     // 用来注销连接的通道
}

// 消息, 用来存储每个消息的信息
type Message struct {
	Sender    string `json:"sender,omitempty"`    // 消息发送者
	Recipient string `json:"recipient,omitempty"` // 消息接收者
	Content   string `json:"content,omitempty"`   // 消息内容
}




// Start 方法, 用来启动客户端管理器, 并且监听各个通道上的事件
func (manager *ClientManager) Start(){
	for{
		select{
		case client := <- manager.Register:
			// 在clients表里添加相应的连接
			manager.Clients[client] = true
			// 创建信息，并转为json消息
			mes := &Message{Content: fmt.Sprintf("A new socket %s has been registered",client.Id)}
			jsonmes,_ := json.Marshal(mes)
			// 调用manger的send方法
			manager.send(client,jsonmes)
		
		case client := <- manager.Unregister:
			// 判断client是否在manager记录中
			if manager.Clients[client]{
				// 关闭client中的chan并在clients中将记录删去
				close(client.Send)
				delete(manager.Clients,client)
			}
		case mes := <- manager.Broadcast:
			// 遍历每个存在的client
			for client := range manager.Clients {				
				select{
				// 向client的信息通道中发送信息
				case client.Send <- mes:
				// 如果通道阻塞,则认定该client已断开连接，将这个client注销
				default:
					close(client.Send)
					delete(manager.Clients,client)
				}
			}
		}
	}
}

// send功能是将新上线的客户端通知给其他客户端
func (manager *ClientManager) send(ignore *Client, mes []byte){
	for client := range manager.Clients{
		if client != ignore{
			client.Send <- mes
		}
	}
}


