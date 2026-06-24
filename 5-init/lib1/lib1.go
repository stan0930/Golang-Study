package lib1

import "fmt"

// 当前lib1包提供的API
func Lib1Test() {//Lib1Test开头大写说明是对外开放接口
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1.init()...")
}
