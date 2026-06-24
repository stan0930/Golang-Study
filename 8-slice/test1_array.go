package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝，无法改变数组值
	for index, value := range myArray {
		fmt.Println("index ,", index, ",value = ", value)
	}
}

func main() {
	//固定长度数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4} //后面6个默认是0

	myArray3 := [4]int{11, 22, 33, 44}

	//for i:=0;i<10;i++ {
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ",value = ", value)
	}
	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

}
