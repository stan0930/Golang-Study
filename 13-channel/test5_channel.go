// 单流程下一个go只能监控一个channel的状态，
// select可以完成监控多个channel的状态
// select具备多路channel的监控状态功能
package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: //用case监视
			//如果c可写，则该case就会进来
			x = y
			y = x + y
		case <-quit: //用case监视
			fmt.Println("quit")
			return
		}
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)
	//subgo
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	//main go
	fibonacii(c, quit)
}
