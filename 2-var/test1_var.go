package main

/*
四种变量的声明方式
*/
import (
	"fmt"
)

// 声明全局变量 方法一，方法二，方法三都可以
var gA int = 100
var gB = 200

//用方法四来申明全局变量
// :=只能用在函数体内来声明
//gC := 200

func main() {
	//方法1：
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of a = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s,type of b = %T\n", bb, bb)

	//方法三
	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of a = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//方法四：（常用方法）省去var
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f =f", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA = ", gA, ",gB = ", gB)
	//fmt.Println("gC = ",gC)

	//声明多变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ",yy = ", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//多行变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ",jj = ", jj)

}
