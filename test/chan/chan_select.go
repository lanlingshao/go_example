package main

import "fmt"

func main() {
	test_select1()
}

// 未完成，待完善
// 为了验证src/runtime/chan.go 11行特性，
func test_select1() {
	c := make(chan int, 0)
	for i := 1; i < 10; i++ {
		select {
		case c <- i:
			fmt.Printf("%+v send item: %d\n", c, i)
		case n, ok := <-c:
			fmt.Printf("recv %d from %+v %+v\n", n, c, ok)
		default:
			fmt.Printf(".\n")
		}
	}
}
