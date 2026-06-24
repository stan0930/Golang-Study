package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len=3 cap=3

	//[0,2)
	s1 := s[0:2]
	//s和s1指向的数组是一致的

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

}
