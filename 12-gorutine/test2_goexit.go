//go关键字创建协程
package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() { //goroutine创建一个无名func
		defer fmt.Println("A.defer")

		//return

		func() {
			defer fmt.Println("B.defer")
			//退出当前goroutine
			runtime.Goexit() //专门为退出goruntine设计的退出方法
			fmt.Println("B")
		}() //创建完直接调用()

		fmt.Println("A")
	}() //创建完直接调用()
	*/

	//有参版本
	go func(a int,b int) bool{
		fmt.Println("a = ",a,"b = ",b)
		return true
	}(10,20)


	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
