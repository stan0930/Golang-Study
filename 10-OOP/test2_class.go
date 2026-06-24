//类名、属性名、方法名首字母大写表示对外(其他包)可以访问，否则只能在本包访问

package main

import "fmt"

// 如果类名首字母大写,其他包也能够访问当前类
type Hero struct {
	//如果类的属性首字母大写,表示该属性是对外能够访问的,否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int //私有
}

// func (this Hero) Show() {
// 	fmt.Println("Name = ",this.Name)
// 	fmt.Println("Ad = ",this.Ad)
// 	fmt.Println("Level = ",this.Level)
// }

// func (this Hero) GetName() string{
// 	fmt.Println("Name = ", this.Name)
// 	return this.Name
// }

//	func (this Hero) SetName(newName string){
//		//this 只调用该方法的对象的一个副本（拷贝）
//	}
func (this Hero) Show() { //方法首字母大写,其他模块,其他包中都可以访问
						  //这个this只是副本，并不是当前对象
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	fmt.Println("Name = ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName

}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("li4")

	hero.Show()

	//fmt.Println()  go语言中方法，变量首字母大小写是有区别的

}
