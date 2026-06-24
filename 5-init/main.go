package main

import (
	
	mylib2 "GolangStudy/5-init/lib2" //mylib2别名
	_ "GolangStudy/5-init/lib1"      //_ 导入不使用
	//. "GolangStudy/5-init/lib2" // . 尽量不要用，会有歧义，尽量用别名
)

func main() {
	//lib1.Lib1Test()
	mylib2.Lib2Test()
	//Lib2Test()
}
