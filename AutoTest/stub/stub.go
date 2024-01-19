package main

import (
	"fmt"
	"github.com/prashantv/gostub"
)
var str = "xdcute.com"

func main() {
	// 1.对全局变量进行打桩
	// 不论是调用 Stub 函数还是 StubFunc 函数，
	// 都会生成一个 Stubs 对象，Stubs 对象仍然有 Stub 方法和 StubFunc 方法，
	// 所以在一个测试用例中可以同时对多个全局变量、函数或过程打桩。
	// 全局变量、函数或过程会将初始值存在一个 map 中，并在延迟语句中通过 Reset 方法统一做回滚处理

	//输出
	//learnku
	//xdcute
	stubs := gostub.Stub(&str, "learnku")
	defer stubs.Reset()
	fmt.Println(str)
	// 可以多次打桩
	stubs.Stub(&str, "xdcute")
	fmt.Println(str)

	// 2. 函数打桩
	// 2.1 有输入的函数
	var printStr = func(val string) string {
        return val
    }
    stubs = gostub.Stub(&printStr, func(val string) string {
        return "hello," + val
    })
	defer stubs.Reset()
    fmt.Println("After stub:", printStr("xdcute"))

	//输出
	//After stub: hello,xdcute

	// 无输入值的打桩
	var printStr2 = func(val string) string {
        return val
    }
    // StubFunc 第一个参数必须是一个函数变量的指针，该指针指向的必须是一个函数变量，第二个参数为函数 mock 的返回值
    stubs = gostub.StubFunc(&printStr2, "xdcute,万生世代")
    defer stubs.Reset()
    fmt.Println("After stub:", printStr2("lalala"))

    //输出
    //After stub: xdcute,万生世代

	// 3.过程打桩(没有返回值的函数)

	var printStr3 = func(val string) {
    fmt.Println(val)
	}
	stubs = gostub.StubFunc(&printStr3)
	printStr3("xdcute")
	defer stubs.Reset()
}


