package monkey

import (
	"fmt"
	"reflect"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/v2/test/fake"
)

// 给函数打桩
// 对齐插桩的时候也要用一样的签名(输入输出需要一样)

func Time() {
	// 为了实现利用time.Now()返回固定时间，而不是实际时间
	now := time.Now()
	// ApplyFunc第一个参数是函数名，第二个参数是桩函数。
	// 测试完成后，patches 对象通过 Reset 成员方法删除所有测试桩
	var p = gomonkey.ApplyFunc(time.Now, func() time.Time {
		return now
	})
	defer p.Reset()
}

func Add() {
	// ApplyMethod函数的多个参数
	// 第一个参数是目标类的指针变量的反射类型，可以用 reflect.TypeOf 来获取。
	// 第二个参数是字符串形式的方法名，
	// 第三个参数是桩函数。测试完成后，patches 对象通过 Reset 成员方法删除所有测试桩

	var s *fake.Slice
	patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "Add", func(_ *fake.Slice, _ int) error {
		return nil

	})
	defer patches.Reset()
}

func ReadLeaf(url string) (string, error) {
	output := fmt.Sprintf("%s, %s!", "Hello", "World")
	return output, nil
}
