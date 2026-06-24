package main

import "fmt"

func main() {
	//申明slice1是一个切片，并初始化
	//slice1 := []int{1, 2, 3}

	//声明slice是一个切片，但是并没有给slice分配空间，len=0
	//var slice1 []int
	//通过make开辟空间,默认值是0
	//slice1 = make([]int, 3)

	//声明slice是一个切片，同时给slice分配空间，初始化0
	//var slice1 []int = make(int[], 3)

	//声明slice是一个切片，同时给slice分配空间，初始化0,通过:=推到是一个切片
	slice1 := make([]int, 3)

	fmt.Println("len = %d,slice = %v\n", len(slice1), slice1)

	//判断一个slice是否为0（没空间）
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}

}
