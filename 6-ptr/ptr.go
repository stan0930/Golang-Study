package main

import "fmt"

func changeValue(p *int) {
	*p = 1000
}

func main() {
	a := 100
	changeValue(&a)
	fmt.Println("a=", a)

}
