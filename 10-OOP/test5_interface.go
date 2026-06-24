//interface{}万能空接口，类型断言

package main

import "fmt"

// interface{}是万能数据类型
func myFunc(a interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(a)

	//interface{} 该如何区分 此时如何引用的的底层数据类型到底是什么？

	//给interface{}提供“类型断言”的机制
	v, ok := a.(string)//
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string")

		fmt.Printf("value type is %T\n", v)
	}
}

type book struct {
	auth string
}

func main() {
	book := book{auth: "Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)

}
