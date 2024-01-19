package autotest

import "fmt"

func Hello()string{
	return "Hello, world"
}
func Add(a,b int)int{
	return a+b
}
func CheckUrl(url string)bool{
	var urlList = [2]string{"1.com", "2.com"}
	for v := range urlList{
		if urlList[v] == url{
			return true
		}
	}
	return false
}

func main(){
	fmt.Println(Hello())
}