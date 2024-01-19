package main

import (
    "fmt"
    "log"
)
//定义了一个订单接口，有一个获取名称的方法
type OrderDBI interface {
    GetName(orderid int) (string)
}
//定义结构体
type OrderInfo struct {
    orderid int
}
//实现接口的方法
func (order OrderInfo) GetName(orderid int) string {
    log.Println("原本应该连接数据库去取名称")
    return "xdcute"
}

func main() {
    //创建接口实例
    var orderDBI OrderDBI = new(OrderInfo)
    //调用方法，返回名称
	// 假设GetName是一个数据库操作，那调试的时候就不太方便
	// mock是指在测试过程中，对于一些不容易构造或不易获取的对象，用一个虚拟对象
	// 可以模拟一个假的数据库对象，提前定义好返回的内容

    // 踩坑，利用mockgen命令行的时候，会找不到go文件，需要输入绝对路径 

    ret := orderDBI.GetName(1)
    fmt.Println("取到的用户名:",ret)
}