package monkey

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/fake"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestApplyMethodSeq(t *testing.T) {
	e := &fake.Etcd{}
	convey.Convey("default times is 1", t, func() {
	  info1 := "hello cpp"
	  info2 := "hello golang"
	  info3 := "hello gomonkey"
	// 设置三个不同的实例，设置要由Etcd客户端，模拟Retrieve方法返回的参数
	  outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{info1, nil}},
		{Values: gomonkey.Params{info2, nil}},
		{Values: gomonkey.Params{info3, nil}},
	  }
	  patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(e), "Retrieve", outputs)
	  defer patches.Reset()
	  output, err := e.Retrieve("")
	  assert.Equal(t, nil, err)
	  assert.Equal(t, info1, output)
	  output, err = e.Retrieve("")
	  assert.Equal(t, nil, err)
	  assert.Equal(t, info2, output)
	  output, err = e.Retrieve("")
	  assert.Equal(t, nil, err)
	  assert.Equal(t, info3, output)
	})
  }
  


  func TestApplyFuncReturn(t *testing.T) {
    convey.Convey("TestApplyFuncReturn", t, func() {
        convey.Convey("declares the values to be returned", func() {
            info := "hello cpp"
            patches := gomonkey.ApplyFuncReturn(ReadLeaf, info, nil)
            defer patches.Reset()
            for i := 0; i < 10; i++ {
                output, err := ReadLeaf("")
                convey.So(err, convey.ShouldEqual, nil)
                convey.So(output, convey.ShouldEqual, info)
            }
        })
    })
}


// 对全局变量打标
var num = 10

func TestApplyGlobalVar(t *testing.T) {
	convey.Convey("TestApplyGlobalVar", t, func() {

	convey.Convey("change", func() {
      patches := gomonkey.ApplyGlobalVar(&num, 150)
      defer patches.Reset()
      assert.Equal(t, num, 150)
    })

    convey.Convey("recover", func() {
      assert.Equal(t, num, 10)
    })
  })
}
