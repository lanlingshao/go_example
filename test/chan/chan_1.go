package main

import "fmt"

func main() {
	c := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		c <- i
	}
	close(c)
	for j := 1; j <= 4; j++ {
		n := <-c
		fmt.Printf("%+v ", n)
	}
	fmt.Printf("\n")
	c <- 1
}

// 1、从关闭的channel获取值会先获取channel中缓存的值，然后缓存值取完后，会获取默认值
// 2、发送值给已关闭的channel会panic
// result:
// 1 2 3 0
// panic: send on closed channel
