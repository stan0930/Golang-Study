package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 返回多个返回值但是都是匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777

}

// 返回多个返回值，有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--- foo3 ---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// r1,r2 属于foo3的形参，初始化默认都是0
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000

	return
	//return 100,200

}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 300)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

}
