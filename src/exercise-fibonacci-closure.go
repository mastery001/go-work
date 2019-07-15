package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	back1 , back2 := 0 , 1
	return func() int {
		temp := back1
		back1,back2 = back2,(back1+back2)
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
