package constiota

import "fmt"

// const来定义枚举类型
const (
	//可以在const()添加一个关键字 iota ，每行的iota都会累加1，第一行iota的默认值是0
	BEIJING  = 10 * iota //iota = 0
	SHANGHAI             //iota = 1
	SHENZHEN             //iota = 2
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	//常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)
	//常量不允许修改
	//length = 100

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b= ", b)
	fmt.Println("a = ", c, "b= ", d)
	fmt.Println("a = ", e, "b= ", f)
	fmt.Println("a = ", g, "b= ", h)
	fmt.Println("a = ", i, "b= ", k)

	//iota只能够配合const()一起使用，iota只有在cosnt进行累加
	//var a int = iota

}
