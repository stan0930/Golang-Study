// channel有堵塞作用
package main

import (
	"fmt"
	"time"
)

func main() { //主协程maingo
	//定义一个channel
	//c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行...")

		//c <- 666 //将666发送给c
	}() //执行
	time.Sleep(2 * time.Second)
	//num := <-c //从c中接受数据，并赋值给num

	//fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
	
}
