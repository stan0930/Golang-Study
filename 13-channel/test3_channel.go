package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	for {
		//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {//从 c 里面取一个值，放到 data 里,同时用 ok 判断这次是否成功取到
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
