package main

import "fmt"

func printArray(myArray []int) {
	//切片是引用传递
	//_表示匿名变量，不关心的东西可以用匿名
	for _, value := range myArray {
		fmt.Println("value", value)
	}
	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray(myArray)

	fmt.Print("=======\n")

	for _, value := range myArray {
		fmt.Println("value=", value)
	}
}
