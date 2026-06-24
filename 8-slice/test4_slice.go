package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//向切片numbers追加一个元素1,number len = 4 [0,0,0,1]
	numbers = append(numbers, 1)

	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)

	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)
	//向一个容量cap已经满的slice追加元素，slice会扩容一个cap
	numbers = append(numbers, 3)

	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("------------")
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
