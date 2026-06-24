package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string , value :"abcd">
	a = "aceld"

	//pair<type:string , value :"abcd">
	var allType interface{}
	allType = a

	str, _ := allType.(string) //.(type)断言传回两个值，第一个是变量值，第二个是变量类型
	fmt.Println(str)
}
