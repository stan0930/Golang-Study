package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aceld", 18}

	// 注意：方法接收者是 *User，所以 Call() 属于 *User 的方法集，不属于 User。
	// 如果这里传入的是 user，reflect.TypeOf(input) 得到的是 main.User，NumMethod() 拿不到 Call。
	// 如果想通过反射获取 Call 方法，需要传入 &user，让 inputType 变成 *main.User。

	DoFieldAndMethod(&user)
}

func DoFieldAndMethod(input interface{}) {
	// 原始 type 和 value
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	fmt.Println("origin inputType is:", inputType)
	fmt.Println("origin inputValue is:", inputValue)

	// 如果是指针，取 Elem，才能遍历字段
	fieldType := inputType
	fieldValue := inputValue
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
		fieldValue = fieldValue.Elem()
	}

	fmt.Println("fieldType is:", fieldType.Name())
	fmt.Println("fieldValue is:", fieldValue)

	// 通过 type 获取字段
	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		value := fieldValue.Field(i).Interface()

		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}

	// 通过 type 获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
