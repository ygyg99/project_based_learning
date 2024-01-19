package autotest

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func Test_sample(t *testing.T) {
	in := Hello()
	out := "Hello, world"
	if in != out {
		t.Errorf("got %q want %q", in, out)
	}
}

func Test_s2(t *testing.T) {
	sum := Add(5, 5)
	main()
	if sum == 10 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}

// 利用Goconvey库
func TestCheckUrl(t *testing.T) {
	convey.Convey("TestCheckUrl", t, func() {
		convey.Convey("TestCheckUrl true", func() {
			ok := CheckUrl("1.com")
			convey.So(ok, convey.ShouldBeTrue)
		})

		convey.Convey("TestCheckUrl false", func() {
			ok := CheckUrl("3.com")
			convey.So(ok, convey.ShouldBeFalse)
		})

		convey.Convey("TestCheckUrl null", func() {
			ok := CheckUrl("")
			convey.So(ok, convey.ShouldBeFalse)
		})
	})
}

func TestCheckUrl2(t *testing.T) {
	ok := CheckUrl("1.com")
	assert.True(t, ok)
	// 结合表格进行多个测试
	assert := assert.New(t)
	var tests = []struct {
		input  string
		expect bool
	}{
		{"1.com", true},
		{"3.com", false},
		{"", false},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		assert.Equal(CheckUrl(test.input), test.expect)
	}
}
